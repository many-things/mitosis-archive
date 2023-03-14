package sim

import (
	"github.com/many-things/mitosis/x/multisig/server"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
)

func SimulateMsgSubmitPubkey(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &server.MsgSubmitPubkey{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SubmitPubkey simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SubmitPubkey simulation not implemented"), nil, nil
	}
}