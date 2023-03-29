package keeper_test

import (
	"testing"

	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
	"gotest.tools/assert"
)

func Test_RegisterSignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	sig := types.Signature("signature")
	err := k.RegisterSignature(ctx, chainID, 0, valAddr, sig)
	assert.NilError(t, err)

	// validate registered successfully
	res, err := repo.Load(0, valAddr)
	assert.NilError(t, err)
	assert.DeepEqual(t, sig, res)
}

func Test_RemoveSignature(_ *testing.T) {
	// TODO: implements
}

func Test_QuerySignature(_ *testing.T) {
	// TODO: implements
}

func Test_QuerySignatureList(_ *testing.T) {
	// TODO: implements
}
