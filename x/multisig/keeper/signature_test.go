package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/testutils"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"

	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"gotest.tools/assert"
)

func Test_RegisterSignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	signSignature := exported.SignSignature{
		Chain: chainID,
		SigID: 0,
		Items: []*exported.SignSignature_Item{{
			Participant: valAddr,
			Signature:   exported.Signature("signature"),
		}},
	}
	err := k.RegisterSignature(ctx, chainID, &signSignature)
	assert.NilError(t, err)

	// validate registered successfully
	res, err := repo.Load(0)
	assert.NilError(t, err)
	assert.DeepEqual(t, &signSignature, res)
}

func Test_RemoveSignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to remove not exist signature
	err := k.RemoveSignature(ctx, chainID, 0)
	assert.Error(t, err, "sign_signature: not found")

	// try to remove exist signature
	signSignature := exported.SignSignature{
		Chain: chainID,
		SigID: 0,
		Items: []*exported.SignSignature_Item{{
			Participant: valAddr,
			Signature:   exported.Signature("signature"),
		}},
	}
	_ = repo.Save(&signSignature)

	err = k.RemoveSignature(ctx, chainID, 0)
	assert.NilError(t, err)

	// validate signature not exists
	_, err = repo.Load(0)
	assert.Error(t, err, "sign_signature: not found")
}

func Test_QuerySignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to query not exist signature
	_, err := k.QuerySignature(ctx, chainID, 0)
	assert.Error(t, err, "sign_signature: not found")

	// try to query exist signature
	signSignature := exported.SignSignature{
		Chain: chainID,
		SigID: 0,
		Items: []*exported.SignSignature_Item{{
			Participant: valAddr,
			Signature:   exported.Signature("signature"),
		}},
	}
	_ = repo.Save(&signSignature)

	res, err := k.QuerySignature(ctx, chainID, 0)
	assert.NilError(t, err)
	assert.DeepEqual(t, res, &signSignature)
}

func Test_QuerySignatureList(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignatureRepo(cdc, ctx.KVStore(storeKey), chainID)

	var signatures []mitosistype.KV[uint64, *exported.SignSignature]
	var i uint64
	for i = 0; i < 10; i++ {
		signSignature := exported.SignSignature{
			Chain: chainID,
			SigID: i,
			Items: []*exported.SignSignature_Item{{
				Participant: sdk.ValAddress(fmt.Sprintf("addr%d", i)),
				Signature:   exported.Signature(fmt.Sprintf("signature%d", i)),
			}},
		}

		_ = repo.Save(&signSignature)
		signatures = append(signatures, mitosistype.NewKV(i, &signSignature))
	}

	res, _, err := k.QuerySignatureList(ctx, chainID, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	assert.DeepEqual(t, res, signatures)
}
