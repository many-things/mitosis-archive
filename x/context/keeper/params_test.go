package keeper_test

import (
	"testing"

	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/context/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ContextKeeper(t)
	params := types.DefaultParams()

	//k.SetParams(ctx, params)
	//
	//require.EqualValues(t, params, k.GetParams(ctx))
}
