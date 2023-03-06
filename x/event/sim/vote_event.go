package sim

import (
	"github.com/many-things/mitosis/x/event/server"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/many-things/mitosis/x/event/keeper"
	"github.com/many-things/mitosis/x/event/types"
)

func SimulateMsgVoteEvent(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &server.MsgVoteEvent{
			Voter: simAccount.Address.String(),
		}

		// TODO: Handling the VoteEvent simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "VoteEvent simulation not implemented"), nil, nil
	}
}
