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

func Test_RemoveSignEvent(_ *testing.T) {
	// TODO: implements
}

func Test_UpdateSignStatus(_ *testing.T) {
	// TODO: implements
}

func Test_QuerySign(_ *testing.T) {
	// TODO: implements
}

func Test_QuerySignList(_ *testing.T) {
	// TODO: implements
}
