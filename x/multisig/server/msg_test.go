package server

import (
	"context"
	"fmt"
	"github.com/many-things/mitosis/pkg/testutils"
	"github.com/many-things/mitosis/x/multisig/types"
	"gotest.tools/assert"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, MsgServer, context.Context) {
	k, ctx, _, _ := keepertest.MultisigKeeper(t)
	return k, NewMsgServer(k), sdk.WrapSDKContext(ctx)
}

func Test_StartKeygen_Failure(t *testing.T) {
	// TODO: impelment
}

func Test_StartKeygen_Success(t *testing.T) {
	k, s, ctx := setupMsgServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := testutils.GenValAddress(t)

	// StartKeygen requires registered keygen
	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []sdk.ValAddress{valAddr},
		Status:       types.Keygen_StatusAssign,
	}

	_, _ = k.RegisterKeygenEvent(wctx, chainID, &keygen)

	// Send StartKeygen
	_, err := s.StartKeygen(wctx, &MsgStartKeygen{
		Module:       "module",
		KeyID:        types.KeyID(fmt.Sprintf("%s-%d", chainID, keygen.KeyID)),
		Participants: []sdk.ValAddress{valAddr},
	})
	assert.NilError(t, err)

	stat, err := k.QueryKeygen(wctx, chainID, 0)
	assert.NilError(t, err)
	assert.Equal(t, stat.Status, types.Keygen_StatusExecute)

	// Re-send. Not changed.
	_, err = s.StartKeygen(wctx, &MsgStartKeygen{
		Module:       "module",
		KeyID:        types.KeyID(fmt.Sprintf("%s-%d", chainID, keygen.KeyID)),
		Participants: []sdk.ValAddress{valAddr},
	})
	assert.NilError(t, err)

	stat, err = k.QueryKeygen(wctx, chainID, 0)
	assert.NilError(t, err)
	assert.Equal(t, stat.Status, types.Keygen_StatusExecute)
}

func Test_SubmitPubKey(t *testing.T) {
	// TODO: implement
}

func Test_SubmitSignature(t *testing.T) {
	// TODO: implement
}
