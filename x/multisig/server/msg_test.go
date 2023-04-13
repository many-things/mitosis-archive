package server

import (
	"context"
	"fmt"
	"github.com/many-things/mitosis/pkg/testutils"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/types"
	"gotest.tools/assert"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, MsgServer, context.Context) {
	k, ctx, _, _ := keepertest.MultisigKeeper(t)
	return k, NewMsgServer(k, nil), sdk.WrapSDKContext(ctx)
}

func Test_StartKeygen_Failure(t *testing.T) {
	k, s, ctx := setupMsgServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := testutils.GenValAddress(t)
	otherAddr := testutils.GenValAddress(t)
	keyID := exported.KeyID(fmt.Sprintf("%s-%d", chainID, 0))

	// Request not exist Keygen event
	_, err := s.StartKeygen(wctx, &MsgStartKeygen{
		Module:       "module",
		KeyID:        keyID,
		Participants: []sdk.ValAddress{valAddr},
	})
	assert.Error(t, err, "keygen: not found")

	// Request Already finished Keygen event
	kg := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []*types.Keygen_Participant{{Address: valAddr}},
		Status:       types.Keygen_StatusComplete,
	}
	_, _ = k.RegisterKeygenEvent(wctx, chainID, &kg)

	_, err = s.StartKeygen(wctx, &MsgStartKeygen{
		Module:       "module",
		KeyID:        keyID,
		Participants: []sdk.ValAddress{valAddr},
	})
	assert.Error(t, err, "keygen: cannot start finished keygen: invalid request")

	// Request wrong participant
	_, _ = k.UpdateKeygenStatus(wctx, chainID, 0, types.Keygen_StatusAssign)
	_, err = s.StartKeygen(wctx, &MsgStartKeygen{
		Module:       "module",
		KeyID:        keyID,
		Participants: []sdk.ValAddress{otherAddr},
	})
	assert.Error(t, err, "keygen: invalid participants: invalid request")
}

func Test_StartKeygen_Success(t *testing.T) {
	k, s, ctx := setupMsgServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := testutils.GenValAddress(t)

	// StartKeygen requires registered keygen
	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []*types.Keygen_Participant{{Address: valAddr}},
		Status:       types.Keygen_StatusAssign,
	}

	_, _ = k.RegisterKeygenEvent(wctx, chainID, &keygen)

	// Send StartKeygen
	_, err := s.StartKeygen(wctx, &MsgStartKeygen{
		Module:       "module",
		KeyID:        exported.KeyID(fmt.Sprintf("%s-%d", chainID, keygen.KeyID)),
		Participants: []sdk.ValAddress{valAddr},
	})
	assert.NilError(t, err)

	stat, err := k.QueryKeygen(wctx, chainID, 0)
	assert.NilError(t, err)
	assert.Equal(t, stat.Status, types.Keygen_StatusExecute)

	// Re-send. Not changed.
	_, err = s.StartKeygen(wctx, &MsgStartKeygen{
		Module:       "module",
		KeyID:        exported.KeyID(fmt.Sprintf("%s-%d", chainID, keygen.KeyID)),
		Participants: []sdk.ValAddress{valAddr},
	})
	assert.NilError(t, err)

	stat, err = k.QueryKeygen(wctx, chainID, 0)
	assert.NilError(t, err)
	assert.Equal(t, stat.Status, types.Keygen_StatusExecute)
}

func Test_SubmitPubKey(t *testing.T) {
	k, s, ctx := setupMsgServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := testutils.GenValAddress(t)
	pubKey := testutils.GenPublicKey(t)

	// ensure pubkey not exist yet
	_, err := k.QueryKeygenResult(wctx, chainID, 0)
	assert.Error(t, err, "keygen: not found")

	_, err = s.SubmitPubkey(wctx, &MsgSubmitPubkey{
		Module:      "module",
		KeyID:       exported.KeyID(fmt.Sprintf("%s-%d", chainID, 0)),
		Participant: valAddr,
		PubKey:      pubKey,
	})
	assert.NilError(t, err)

	// check pubKey
	res, err := k.QueryKeygenResult(wctx, chainID, 0)
	assert.NilError(t, err)
	assert.DeepEqual(t, res.Items[0].PubKey, pubKey)
}

func Test_SubmitSignature(t *testing.T) {
	k, s, ctx := setupMsgServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := testutils.GenValAddress(t)
	signature := exported.Signature("signature")

	// ensure signature not exist yet
	_, err := k.QuerySignResult(wctx, chainID, 0)
	assert.Error(t, err, "sign_signature: not found")

	// request signature
	_, err = s.SubmitSignature(wctx, &MsgSubmitSignature{
		Module:      "module",
		SigID:       exported.SigID(fmt.Sprintf("%s-%d", chainID, 0)),
		Participant: valAddr,
		Signature:   signature,
	})
	assert.NilError(t, err)

	// ensure signature exists
	res, err := k.QuerySignResult(wctx, chainID, 0)
	assert.NilError(t, err)
	assert.DeepEqual(t, res.Items[0].Signature, signature)
}
