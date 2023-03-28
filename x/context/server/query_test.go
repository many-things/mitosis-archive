package server

import (
	"context"
	"testing"

	"github.com/many-things/mitosis/x/context/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/stretchr/testify/require"
)

func setupQueryServer(t testing.TB) (keeper.Keeper, QueryServer, context.Context) {
	k, ctx := testkeeper.ContextKeeper(t)
	return *k, NewQueryServer(*k), sdk.WrapSDKContext(ctx)
}

func TestParamsQuery(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	params := types.DefaultParams()
	k.SetParams(wctx, params)

	response, err := s.Params(wctx, &QueryParams{})
	require.NoError(t, err)
	require.Equal(t, &QueryParamsResponse{Params: params}, response)
}
