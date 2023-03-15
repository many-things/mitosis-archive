package sim

import (
<<<<<<<< HEAD:x/event/sim/vote_event.go
	"github.com/many-things/mitosis/x/event/server"
========
	"github.com/many-things/mitosis/x/multisig/server"
>>>>>>>> feat/queue:x/multisig/sim/start_keygen.go
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
)

func SimulateMsgStartKeygen(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
<<<<<<<< HEAD:x/event/sim/vote_event.go
		msg := &server.MsgVoteEvent{
			Voter: simAccount.Address.String(),
========
		msg := &server.MsgStartKeygen{
			Creator: simAccount.Address.String(),
>>>>>>>> feat/queue:x/multisig/sim/start_keygen.go
		}

		// TODO: Handling the StartKeygen simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "StartKeygen simulation not implemented"), nil, nil
	}
}
