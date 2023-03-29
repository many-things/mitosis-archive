package keeper_test

import (
	crand "crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
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

func TestGetParams(t *testing.T) {
	k, ctx, _, _ := testkeeper.MultisigKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
