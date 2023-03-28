package keeper_test

import (
	crand "crypto/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProxy(t *testing.T) {
	k, ctx := testkeeper.EventKeeper(t)

	var val sdk.ValAddress
	var acc sdk.AccAddress
	{
		tmp := make([]byte, 32)
		_, err := crand.Read(tmp)
		require.NoError(t, err)
		val = tmp
	}
	{
		tmp := make([]byte, 32)
		_, err := crand.Read(tmp)
		require.NoError(t, err)
		acc = tmp
	}

	require.NoError(t, k.RegisterProxy(ctx, val, acc))

	accRes, found := k.QueryProxy(ctx, val)
	require.True(t, found)
	require.Equal(t, acc, accRes)

	valRes, found := k.QueryProxyReverse(ctx, acc)
	require.True(t, found)
	require.Equal(t, val, valRes)

	proxySet, _, err := k.QueryProxies(ctx, &query.PageRequest{Limit: query.MaxLimit})
	require.NoError(t, err)
	require.Equal(t,
		[]mitotypes.KV[sdk.ValAddress, sdk.AccAddress]{
			mitotypes.NewKV(val, acc),
		},
		proxySet,
	)
}
