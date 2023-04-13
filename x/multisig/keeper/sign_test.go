package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/testutils"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"

	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func Test_RegisterSignEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	sign := exported.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        exported.Sign_StatusAssign,
	}
	signID, err := k.RegisterSignEvent(ctx, chainID, &sign)
	assert.NilError(t, err)

	savedSign, err := repo.Load(signID)
	assert.NilError(t, err)
	require.Equal(t, savedSign, &sign)
}

func Test_RemoveSignEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to remove not exist sign event
	err := k.RemoveSignEvent(ctx, chainID, 0)
	assert.Error(t, err, "sign: not found")

	// try to remove exist sign event
	sign := exported.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        exported.Sign_StatusAssign,
	}
	err = repo.Save(&sign)
	assert.NilError(t, err)

	err = k.RemoveSignEvent(ctx, chainID, 0)
	assert.NilError(t, err)

	// try to check sign removed
	_, err = repo.Load(0)
	assert.Error(t, err, "sign: not found")
}

func Test_UpdateSignStatus(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to update not exist sign event
	_, err := k.UpdateSignStatus(ctx, chainID, 0, exported.Sign_StatusComplete)
	assert.Error(t, err, "sign: not found")

	// try to update exist sign event
	sign := exported.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        exported.Sign_StatusAssign,
	}
	_ = repo.Save(&sign)

	updated, err := k.UpdateSignStatus(ctx, chainID, 0, exported.Sign_StatusExecute)
	assert.NilError(t, err)
	assert.Equal(t, updated.Status, exported.Sign_StatusExecute)
}

func Test_QuerySign(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to query not exist sign event
	_, err := k.QuerySign(ctx, chainID, 0)
	assert.Error(t, err, "sign: not found")

	// try to query exist sign event
	sign := exported.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        exported.Sign_StatusAssign,
	}
	_ = repo.Save(&sign)

	res, err := k.QuerySign(ctx, chainID, 0)
	assert.NilError(t, err)
	require.Equal(t, res, &sign)
}

func Test_QuerySignList(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	var signs []mitosistype.KV[uint64, *exported.Sign]
	var i uint64
	for i = 0; i < 10; i++ {
		sign := exported.Sign{
			Chain:         chainID,
			SigID:         i,
			KeyID:         "1",
			Participants:  []sdk.ValAddress{valAddr},
			MessageToSign: []byte("test"),
			Status:        exported.Sign_StatusAssign,
		}
		_ = repo.Save(&sign)

		signs = append(signs, mitosistype.NewKV(i, &sign))
	}

	res, _, err := k.QuerySignList(ctx, chainID, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	require.Equal(t, res, signs)
}
