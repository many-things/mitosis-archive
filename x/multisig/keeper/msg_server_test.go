package keeper_test

import (
	"context"
	"github.com/many-things/mitosis/x/multisig/server"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper"
)

func setupMsgServer(t testing.TB) (server.MsgServer, context.Context) {
	k, ctx := keepertest.MultisigKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
