package hook

import (
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/keeper"
	"github.com/many-things/mitosis/x/context/types"
	multisigexport "github.com/many-things/mitosis/x/multisig/exported"
	multisigtypes "github.com/many-things/mitosis/x/multisig/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(
	ctx sdk.Context,
	_ abci.RequestEndBlock,
	baseKeeper keeper.Keeper,
	multisigKeeper types.MultisigKeeper,
) []abci.ValidatorUpdate {
	ops, _, err := baseKeeper.QueryOperationsByStatus(
		ctx,
		types.Operation_StatusPending,
		&query.PageRequest{Limit: query.MaxLimit},
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to query pending operations"))
	}

	for _, op := range ops {
		keygenResp, _, err := multisigKeeper.QueryKeygenList(ctx, op.Chain, nil)
		if err != nil {
			panic(errors.Wrap(err, "failed to query keygen list"))
		}

		sigID, err := multisigKeeper.RegisterSignEvent(ctx, op.Chain, &multisigexport.Sign{
			Chain: op.Chain,
			KeyID: fmt.Sprintf("%s-%d", keygenResp[0].Value.Chain, keygenResp[0].Value.KeyID),
			OpId:  op.ID,
			Participants: mitotypes.Map(
				keygenResp[0].Value.Participants,
				func(t *multisigtypes.Keygen_Participant, _ int) sdk.ValAddress { return t.Address },
			),
			MessageToSign: op.TxBytesToSign,
			Status:        multisigexport.Sign_StatusAssign,
		})
		if err != nil {
			panic(errors.Wrap(err, "failed to register sign event"))
		}

		if err := baseKeeper.StartSignOperation(ctx, op.ID, sigID); err != nil {
			panic(errors.Wrap(err, "failed to start sign operation"))
		}
	}

	return []abci.ValidatorUpdate{}
}
