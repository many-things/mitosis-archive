package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/testutils"

	mitosistype "github.com/many-things/mitosis/pkg/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
	"gotest.tools/assert"
)

func Test_RegisterSignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	sig := types.Signature("signature")
	err := k.RegisterSignature(ctx, chainID, 0, valAddr, sig)
	assert.NilError(t, err)

	// validate registered successfully
	res, err := repo.Load(0, valAddr)
	assert.NilError(t, err)
	assert.DeepEqual(t, sig, res)
}

func Test_RemoveSignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to remove not exist signature
	err := k.RemoveSignature(ctx, chainID, 0, valAddr)
	assert.Error(t, err, "signature: not found")

	// try to remove exist signature
	sig := types.Signature("signature")
	_ = repo.Create(0, valAddr, sig)

	err = k.RemoveSignature(ctx, chainID, 0, valAddr)
	assert.NilError(t, err)

	// validate signature not exists
	_, err = repo.Load(0, valAddr)
	assert.Error(t, err, "signature: not found")
}

func Test_QuerySignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to query not exist signature
	_, err := k.QuerySignature(ctx, chainID, 0, valAddr)
	assert.Error(t, err, "signature: not found")

	// try to query exist signature
	sig := types.Signature("signature")
	_ = repo.Create(0, valAddr, sig)

	res, err := k.QuerySignature(ctx, chainID, 0, valAddr)
	assert.NilError(t, err)
	assert.DeepEqual(t, res, sig)
}

func Test_QuerySignatureList(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)

	var signatures []mitosistype.KV[sdk.ValAddress, types.Signature]
	var i uint64
	for i = 0; i < 10; i++ {
		sig := types.Signature(fmt.Sprintf("signature%d", i))
		valAddr := sdk.ValAddress(fmt.Sprintf("addr%d", i))

		_ = repo.Create(0, valAddr, sig)
		signatures = append(signatures, mitosistype.NewKV(valAddr, sig))
	}

	res, _, err := k.QuerySignatureList(ctx, chainID, 0, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	assert.DeepEqual(t, res, signatures)
}
