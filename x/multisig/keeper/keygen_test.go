package keeper_test

import (
	crand "crypto/rand"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

const (
	chainID = "chain"
)

func genValAddr(t *testing.T) sdk.ValAddress {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)
	return bz
}

func genNotfoundErrMsg(id uint64) string {
	return fmt.Sprintf("cannot find keygen: id %d", id)
}

func Test_RegisterKeygenEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	valAddr := genValAddr(t)

	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []sdk.ValAddress{valAddr},
		Status:       1,
	}
	kgID, err := k.RegisterKeygenEvent(ctx, chainID, &keygen)
	assert.NilError(t, err)
	require.Equal(t, kgID, uint64(0))

	repo := state.NewKVChainKeygenRepo(cdc, ctx.KVStore(storeKey), chainID)

	savedKeygen, err := repo.Load(kgID)
	assert.NilError(t, err)
	require.Equal(t, *savedKeygen, keygen)
}

func Test_QueryKeygenEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []sdk.ValAddress{valAddr},
		Status:       1,
	}

	// try to load not exist keygen
	_, err := k.QueryKeygen(ctx, chainID, 0)
	assert.Error(t, err, genNotfoundErrMsg(0))

	// try to load exist keygen
	err = repo.Save(&keygen)
	assert.NilError(t, err)

	res, err := k.QueryKeygen(ctx, chainID, 0)
	assert.NilError(t, err)
	require.Equal(t, res, &keygen)
}

func Test_SaveKeygenEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	// try to update not exist value
	_, err := k.UpdateKeygenStatus(ctx, chainID, 3, types.Keygen_StatusExecute)
	assert.Error(t, err, genNotfoundErrMsg(3))

	// try to update exist variable
	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []sdk.ValAddress{valAddr},
		Status:       1,
	}
	_, err = repo.Create(&keygen)
	assert.NilError(t, err)

	updatedKeygen, err := k.UpdateKeygenStatus(ctx, chainID, keygen.KeyID, types.Keygen_StatusExecute)
	assert.NilError(t, err)
	assert.Equal(t, updatedKeygen.Status, types.Keygen_StatusExecute)
}

func Test_RemoveKeygenEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	valAddr := genValAddr(t)

	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []sdk.ValAddress{valAddr},
		Status:       1,
	}

	// try to delete not exist keygen
	err := k.RemoveKeygenEvent(ctx, chainID, uint64(27))
	assert.Error(t, err, genNotfoundErrMsg(27))

	// try to delete exist keygen
	repo := state.NewKVChainKeygenRepo(cdc, ctx.KVStore(storeKey), chainID)
	err = repo.Save(&keygen)
	assert.NilError(t, err)

	err = k.RemoveKeygenEvent(ctx, chainID, 0)
	assert.NilError(t, err)

	// check keygen deleted
	_, err = repo.Load(0)
	assert.Error(t, err, genNotfoundErrMsg(0))
}
