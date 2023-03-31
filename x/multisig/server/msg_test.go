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

func Test_StartKeygen_Success(t *testing.T) {
	// TODO: implements

}

func Test_StartKeygen_Failure(t *testing.T) {
	// TODO: impelment
}

func Test_SubmitPubKey(t *testing.T) {
	// TODO: implement
}

func Test_SubmitSignature(t *testing.T) {
	// TODO: implement
}
