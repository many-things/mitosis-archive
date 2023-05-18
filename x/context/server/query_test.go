package server

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gotest.tools/assert"
	"testing"

	"github.com/many-things/mitosis/x/context/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/stretchr/testify/require"
)

func setupQueryServer(t testing.TB) (keeper.Keeper, QueryServer, context.Context) {
	k, ctx, _, _, _ := testkeeper.ContextKeeper(t)
	return k, NewQueryServer(k), sdk.WrapSDKContext(ctx)
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

func TestQueryVault(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-1", "osmo1deaddeaddeaddeaddead"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-2", "osmo1c0ffeec0ffeec0ffee"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-3", "osmo1beefbeefbeefbeefbeef"))

	vaultResp, err := s.Vault(wctx, &QueryVault{Chain: "osmo-test-1"})
	require.Nil(t, err)

	assert.Equal(t, vaultResp.Chain, "osmo-test-1")
	assert.Equal(t, vaultResp.Vault, "osmo1deaddeaddeaddeaddead")
}

func TestQueryVaults(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-1", "osmo1deaddeaddeaddeaddead"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-2", "osmo1c0ffeec0ffeec0ffee"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-3", "osmo1beefbeefbeefbeefbeef"))

	vaultsResp, err := s.Vaults(wctx, &QueryVaults{Pagination: nil})
	require.Nil(t, err)

	assert.DeepEqual(t, vaultsResp, &QueryVaultsResponse{
		Vaults: []*QueryVaultResponse{
			{
				Chain: "osmo-test-1",
				Vault: "osmo1deaddeaddeaddeaddead",
			}, {
				Chain: "osmo-test-2",
				Vault: "osmo1c0ffeec0ffeec0ffee",
			}, {
				Chain: "osmo-test-3",
				Vault: "osmo1beefbeefbeefbeefbeef",
			},
		},
		Page: &query.PageResponse{Total: 3},
	})
}
