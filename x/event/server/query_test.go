package server

import (
	"context"
	"github.com/many-things/mitosis/x/event/keeper"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/event/types"
	"github.com/stretchr/testify/require"
)

func setupQueryServer(t testing.TB) (keeper.Keeper, types.QueryServer, context.Context) {
	k, ctx := keepertest.EventKeeper(t)
	return *k, NewQueryServer(*k), sdk.WrapSDKContext(ctx)
}

func TestParamsQuery(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	params := types.DefaultParams()
	k.SetParams(wctx, params)

	response, err := s.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}