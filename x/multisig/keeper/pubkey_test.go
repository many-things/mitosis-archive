package keeper_test

import (
	crand "crypto/rand"
	"testing"

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

func Test_RemovePubKey(_ *testing.T) {
	// TODO: implements
}

func Test_QueryPubKey(_ *testing.T) {
	// TODO: implements
}

func Test_QueryPubKeyList(_ *testing.T) {
	// TODO: implements
}
