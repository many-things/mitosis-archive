package server

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/many-things/mitosis/testutil/keeper"
)

func setupMsgServer(t testing.TB) (MsgServer, context.Context) {
	k, ctx, _, _ := keepertest.MultisigKeeper(t)
	return NewMsgServer(k), sdk.WrapSDKContext(ctx)
}
