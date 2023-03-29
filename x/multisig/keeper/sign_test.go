package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func Test_RegisterSignEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	sign := types.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        types.Sign_StatusAssign,
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
	valAddr := genValAddr(t)

	// try to remove not exist sign event
	err := k.RemoveSignEvent(ctx, chainID, 0)
	assert.Error(t, err, "sign: not found")

	// try to remove exist sign event
	sign := types.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        types.Sign_StatusAssign,
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
	valAddr := genValAddr(t)

	// try to update not exist sign event
	_, err := k.UpdateSignStatus(ctx, chainID, 0, types.Sign_StatusComplete)
	assert.Error(t, err, "sign: not found")

	// try to update exist sign event
	sign := types.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        types.Sign_StatusAssign,
	}
	_ = repo.Save(&sign)

	updated, err := k.UpdateSignStatus(ctx, chainID, 0, types.Sign_StatusComplete)
	assert.NilError(t, err)
	assert.Equal(t, updated.Status, types.Sign_StatusComplete)
}

func Test_QuerySign(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := genValAddr(t)

	// try to query not exist sign event
	_, err := k.QuerySign(ctx, chainID, 0)
	assert.Error(t, err, "sign: not found")

	// try to query exist sign event
	sign := types.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         "1",
		Participants:  []sdk.ValAddress{valAddr},
		MessageToSign: []byte("test"),
		Status:        types.Sign_StatusAssign,
	}
	_ = repo.Save(&sign)

	res, err := k.QuerySign(ctx, chainID, 0)
	assert.NilError(t, err)
	require.Equal(t, res, &sign)
}

func Test_QuerySignList(_ *testing.T) {
	// TODO: implements
}
