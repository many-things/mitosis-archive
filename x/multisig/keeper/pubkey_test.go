package keeper_test

import (
	crand "crypto/rand"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func genPublicKey(t *testing.T) types.PublicKey {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)
	return bz
}

var (
	kvPubKeyRepoKey = []byte{0x02}
)

func genNotFoundPubKeyMsg(id uint64, participant sdk.ValAddress) string {
	return fmt.Sprintf("cannot find pubkey: %s for id %d", participant, id)
}

func Test_RegisterPubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	pubKey := types.PubKey{
		Chain:       chainID,
		KeyID:       0,
		Participant: valAddr,
		PubKey:      genPublicKey(t),
	}
	err := k.RegisterPubKey(ctx, chainID, &pubKey)
	assert.NilError(t, err)

	// test generated successfully
	savedPubKey, err := repo.Load(pubKey.KeyID, pubKey.Participant)
	assert.NilError(t, err)
	require.Equal(t, pubKey, *savedPubKey)
}

func Test_RemovePubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	// try to delete not exist pubKey
	err := k.RemovePubKey(ctx, chainID, 0, valAddr)
	assert.Error(t, err, genNotFoundPubKeyMsg(0, valAddr))

	// try to delete exist pubKey
	pubKey := types.PubKey{
		Chain:       chainID,
		KeyID:       0,
		Participant: valAddr,
		PubKey:      genPublicKey(t),
	}
	err = repo.Create(&pubKey)
	assert.NilError(t, err)

	err = k.RemovePubKey(ctx, chainID, pubKey.KeyID, pubKey.Participant)
	assert.NilError(t, err)

	// validate
	_, err = repo.Load(pubKey.KeyID, pubKey.Participant)
	assert.Error(t, err, genNotFoundPubKeyMsg(pubKey.KeyID, pubKey.Participant))
}

func Test_QueryPubKey(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	// try to query not exist pubKey
	_, err := k.QueryPubKey(ctx, chainID, 0, valAddr)
	assert.Error(t, err, genNotFoundPubKeyMsg(0, valAddr))

	// try to query exist pubKey
	pubKey := types.PubKey{
		Chain:       chainID,
		KeyID:       0,
		Participant: valAddr,
		PubKey:      genPublicKey(t),
	}
	err = repo.Create(&pubKey)
	assert.NilError(t, err)

	res, err := k.QueryPubKey(ctx, chainID, pubKey.KeyID, pubKey.Participant)
	assert.NilError(t, err)
	require.Equal(t, pubKey, *res)
}

func Test_QueryPubKeyList(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)

	for i := 0; i < 10; i++ {
		pubKey := types.PubKey{
			Chain:       chainID,
			KeyID:       1,
			Participant: genValAddr(t),
			PubKey:      genPublicKey(t),
		}
		_ = repo.Create(&pubKey)
	}

	res, _, err := k.QueryPubKeyList(ctx, chainID, 1, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	require.Equal(t, len(res), 10)
}
