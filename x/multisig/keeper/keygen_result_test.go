package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/testutils"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func Test_RegisterKeygenResult(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	pubKey := types.KeygenResult{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.KeygenResult_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	}
	err := k.RegisterKeygenResult(ctx, chainID, &pubKey)
	assert.NilError(t, err)

	// test generated successfully
	savedPubKey, err := repo.Load(pubKey.KeyID)
	assert.NilError(t, err)
	require.Equal(t, pubKey, *savedPubKey)
}

func Test_AddParticipantKeygenResult(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)
	publicKey := testutils.GenPublicKey(t)

	// try to create not exist pubKey
	err := k.AddParticipantKeygenResult(ctx, chainID, 0, valAddr, publicKey)
	assert.Error(t, err, "pubkey: not found")

	_ = repo.Create(&types.KeygenResult{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.KeygenResult_Item{},
	})

	// try to create exist pubKey
	err = k.AddParticipantKeygenResult(ctx, chainID, 0, valAddr, publicKey)
	assert.NilError(t, err)

	pubKey, err := repo.Load(0)
	assert.NilError(t, err)
	assert.DeepEqual(t, pubKey.Items[0], &types.KeygenResult_Item{
		Participant: valAddr,
		PubKey:      publicKey,
	})
	assert.DeepEqual(t, len(pubKey.Items), 1)

	// try to change exist pubKey
	newPublicKey := testutils.GenPublicKey(t)
	err = k.AddParticipantKeygenResult(ctx, chainID, 0, valAddr, newPublicKey)
	assert.NilError(t, err)

	pubKey, err = repo.Load(0)
	assert.NilError(t, err)
	assert.DeepEqual(t, pubKey.Items[0], &types.KeygenResult_Item{
		Participant: valAddr,
		PubKey:      newPublicKey,
	})
	assert.DeepEqual(t, len(pubKey.Items), 1)

	// try to append new publicKEy
	newValAddr := testutils.GenValAddress(t)
	newValPublicKey := testutils.GenPublicKey(t)
	err = k.AddParticipantKeygenResult(ctx, chainID, 0, newValAddr, newValPublicKey)
	assert.NilError(t, err)

	pubKey, err = repo.Load(0)
	assert.NilError(t, err)
	assert.DeepEqual(t, pubKey.Items[1], &types.KeygenResult_Item{
		Participant: newValAddr,
		PubKey:      newValPublicKey,
	})
	assert.DeepEqual(t, len(pubKey.Items), 2)
}

func TestKeeper_RemoveParticipantKeygenResult(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to remove participant pubkey not exists
	err := k.RemoveParticipantKeygenResult(ctx, chainID, 0, valAddr)
	assert.Error(t, err, "pubkey: not found")

	_ = repo.Create(&types.KeygenResult{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.KeygenResult_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	})

	// try to remove participant pubkey not exists
	err = k.RemoveParticipantKeygenResult(ctx, chainID, 0, testutils.GenValAddress(t))
	assert.Error(t, err, "pubkey item: not found")

	// try to remove participant pubkey exists
	err = k.RemoveParticipantKeygenResult(ctx, chainID, 0, valAddr)
	assert.NilError(t, err)
}

func Test_DeleteKeygenResult(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to delete not exist pubKey
	err := k.DeleteKeygenResult(ctx, chainID, 0)
	assert.Error(t, err, "pubkey: not found")

	// try to delete exist pubKey
	pubKey := types.KeygenResult{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.KeygenResult_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	}
	err = repo.Create(&pubKey)
	assert.NilError(t, err)

	err = k.DeleteKeygenResult(ctx, chainID, pubKey.KeyID)
	assert.NilError(t, err)

	// validate
	_, err = repo.Load(pubKey.KeyID)
	assert.Error(t, err, "keygen: not found")
}

func Test_QueryKeygenResult(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to query not exist pubKey
	_, err := k.QueryKeygenResult(ctx, chainID, 0)
	assert.Error(t, err, "keygen: not found")

	// try to query exist pubKey
	pubKey := types.KeygenResult{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.KeygenResult_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	}
	err = repo.Create(&pubKey)
	assert.NilError(t, err)

	res, err := k.QueryKeygenResult(ctx, chainID, pubKey.KeyID)
	assert.NilError(t, err)
	require.Equal(t, pubKey, *res)
}

func Test_QueryKeygenResultList(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainKeygenResultRepo(cdc, ctx.KVStore(storeKey), chainID)

	var i uint64
	for i = 0; i < 10; i++ {
		pubKey := types.KeygenResult{
			Chain: chainID,
			KeyID: i,
			Items: []*types.KeygenResult_Item{{
				Participant: testutils.GenValAddress(t),
				PubKey:      testutils.GenPublicKey(t),
			}},
		}
		_ = repo.Create(&pubKey)
	}

	res, _, err := k.QueryKeygenResultList(ctx, chainID, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	require.Equal(t, len(res), 10)
}
