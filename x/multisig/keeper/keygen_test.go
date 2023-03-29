package keeper_test

import (
	crand "crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
)

const (
	chainID = "chain"
)

func genValAddr(t *testing.T) sdk.ValAddress {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)
	return bz
}

func Test_RegisterKeygenEvent(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	valAddr := genValAddr(t)

	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        1,
		Participants: []sdk.ValAddress{valAddr},
		Status:       1,
	}
	kgID, err := k.RegisterKeygenEvent(ctx, chainID, &keygen)
	require.NoError(t, err)
	require.Equal(t, kgID, uint64(0))

	repo := state.NewKVChainKeygenRepo(cdc, ctx.KVStore(storeKey), chainID)

	savedKeygen, err := repo.Load(kgID)
	require.NoError(t, err)
	require.Equal(t, savedKeygen, &keygen)
}
