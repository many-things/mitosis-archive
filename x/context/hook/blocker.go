package hook

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/x/context/keeper"
	"github.com/many-things/mitosis/x/context/types"
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
		sigID, err := multisigKeeper.RegisterSignEvent(ctx, op.Chain, nil)
		if err != nil {
			panic(errors.Wrap(err, "failed to register sign event"))
		}

		if err := baseKeeper.StartSignOperation(ctx, op.ID, sigID); err != nil {
			panic(errors.Wrap(err, "failed to start sign operation"))
		}
	}

	return []abci.ValidatorUpdate{}
}
