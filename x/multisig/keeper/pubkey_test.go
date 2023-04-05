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

var (
	kvPubKeyRepoKey = []byte{0x02}
)

func Test_RegisterPubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	pubKey := types.PubKey{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.PubKey_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	}
	err := k.RegisterPubKey(ctx, chainID, &pubKey)
	assert.NilError(t, err)

	// test generated successfully
	savedPubKey, err := repo.Load(pubKey.KeyID)
	assert.NilError(t, err)
	require.Equal(t, pubKey, *savedPubKey)
}

func Test_AddParticipantPubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)
	publicKey := testutils.GenPublicKey(t)

	// try to create not exist pubKey
	err := k.AddParticipantPubKey(ctx, chainID, 0, valAddr, publicKey)
	assert.Error(t, err, "pubkey: not found")

	_ = repo.Create(&types.PubKey{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.PubKey_Item{},
	})

	// try to create exist pubKey
	err = k.AddParticipantPubKey(ctx, chainID, 0, valAddr, publicKey)
	assert.NilError(t, err)

	pubKey, err := repo.Load(0)
	assert.NilError(t, err)
	assert.DeepEqual(t, pubKey.Items[0], &types.PubKey_Item{
		Participant: valAddr,
		PubKey:      publicKey,
	})
	assert.DeepEqual(t, len(pubKey.Items), 1)

	// try to change exist pubKey
	newPublicKey := testutils.GenPublicKey(t)
	err = k.AddParticipantPubKey(ctx, chainID, 0, valAddr, newPublicKey)
	assert.NilError(t, err)

	pubKey, err = repo.Load(0)
	assert.NilError(t, err)
	assert.DeepEqual(t, pubKey.Items[0], &types.PubKey_Item{
		Participant: valAddr,
		PubKey:      newPublicKey,
	})
	assert.DeepEqual(t, len(pubKey.Items), 1)

	// try to append new publicKEy
	newValAddr := testutils.GenValAddress(t)
	newValPublicKey := testutils.GenPublicKey(t)
	err = k.AddParticipantPubKey(ctx, chainID, 0, newValAddr, newValPublicKey)
	assert.NilError(t, err)

	pubKey, err = repo.Load(0)
	assert.NilError(t, err)
	assert.DeepEqual(t, pubKey.Items[1], &types.PubKey_Item{
		Participant: newValAddr,
		PubKey:      newValPublicKey,
	})
	assert.DeepEqual(t, len(pubKey.Items), 2)
}

func TestKeeper_RemoveParticipantPubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to remove participant pubkey not exists
	err := k.RemoveParticipantPubKey(ctx, chainID, 0, valAddr)
	assert.Error(t, err, "pubkey: not found")

	_ = repo.Create(&types.PubKey{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.PubKey_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	})

	// try to remove participant pubkey not exists
	err = k.RemoveParticipantPubKey(ctx, chainID, 0, testutils.GenValAddress(t))
	assert.Error(t, err, "pubkey item: not found")

	// try to remove participant pubkey exists
	err = k.RemoveParticipantPubKey(ctx, chainID, 0, valAddr)
	assert.NilError(t, err)
}

func Test_DeletePubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to delete not exist pubKey
	err := k.DeletePubKey(ctx, chainID, 0)
	assert.Error(t, err, "pubkey: not found")

	// try to delete exist pubKey
	pubKey := types.PubKey{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.PubKey_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	}
	err = repo.Create(&pubKey)
	assert.NilError(t, err)

	err = k.DeletePubKey(ctx, chainID, pubKey.KeyID)
	assert.NilError(t, err)

	// validate
	_, err = repo.Load(pubKey.KeyID)
	assert.Error(t, err, "pubkey: not found")
}

func Test_QueryPubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to query not exist pubKey
	_, err := k.QueryPubKey(ctx, chainID, 0)
	assert.Error(t, err, "pubkey: not found")

	// try to query exist pubKey
	pubKey := types.PubKey{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.PubKey_Item{{
			Participant: valAddr,
			PubKey:      testutils.GenPublicKey(t),
		}},
	}
	err = repo.Create(&pubKey)
	assert.NilError(t, err)

	res, err := k.QueryPubKey(ctx, chainID, pubKey.KeyID)
	assert.NilError(t, err)
	require.Equal(t, pubKey, *res)
}

func Test_QueryPubKeyList(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)

	var i uint64
	for i = 0; i < 10; i++ {
		pubKey := types.PubKey{
			Chain: chainID,
			KeyID: i,
			Items: []*types.PubKey_Item{{
				Participant: testutils.GenValAddress(t),
				PubKey:      testutils.GenPublicKey(t),
			}},
		}
		_ = repo.Create(&pubKey)
	}

	res, _, err := k.QueryPubKeyList(ctx, chainID, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	require.Equal(t, len(res), 10)
}
