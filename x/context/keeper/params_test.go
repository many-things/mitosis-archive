package keeper_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/context/types"
)

func TestGetParams(t *testing.T) {
	k, ctx, _, _ := testkeeper.ContextKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
