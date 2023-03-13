package keeper_test

import (
	"github.com/many-things/mitosis/x/multisig/server"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.MultisigKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &server.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &server.QueryParamsResponse{Params: params}, response)
}
