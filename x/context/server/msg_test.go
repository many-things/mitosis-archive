package server

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/many-things/mitosis/testutil/keeper"
)

func setupMsgServer(t testing.TB) (MsgServer, context.Context) {
	k, ctx := keepertest.ContextKeeper(t)
	return NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
