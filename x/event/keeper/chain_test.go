package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChain(t *testing.T) {
	k, ctx := testkeeper.EventKeeper(t)

	// test register
	{
		testcases := []struct {
			chain     string
			expected  byte
			expectErr bool
		}{
			{
				chain:     "mitosis-1",
				expected:  0x00,
				expectErr: false,
			},
			{
				chain:     "osmosis-1",
				expected:  0x01,
				expectErr: false,
			},
			{
				chain:     "cosmoshub-3",
				expected:  0x02,
				expectErr: false,
			},
			{
				chain:     "mitosis-1",
				expected:  0x00,
				expectErr: true,
			},
		}
		for _, testcase := range testcases {
			actual, err := k.RegisterChain(ctx, testcase.chain)
			if !testcase.expectErr {
				require.NoError(t, err)
				require.Equal(t, testcase.expected, actual)
			} else {
				require.Error(t, err)
			}
		}
	}

	// test unregister
	{
		testcases := []struct {
			chain       string
			expectedErr bool
		}{
			{
				chain:       "mitosis-1",
				expectedErr: false,
			},
			{
				chain:       "mitosis-1",
				expectedErr: true,
			},
			{
				chain:       "axelar-1",
				expectedErr: true,
			},
		}
		for _, testcase := range testcases {
			err := k.UnregisterChain(ctx, testcase.chain)
			if !testcase.expectedErr {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		}
	}

	// test queries

	// query single
	_, err := k.QueryChain(ctx, "mitosis-1")
	require.Error(t, err)
	prefix, err := k.QueryChain(ctx, "osmosis-1")
	require.NoError(t, err)
	require.Equal(t, byte(0x01), prefix)

	// query multiple
	prefixes, resp, err := k.QueryChains(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, &query.PageResponse{
		NextKey: nil,
		Total:   2,
	}, resp)
	require.Equal(t, []mitotypes.KV[string, byte]{
		{Key: "cosmoshub-3", Value: 0x02},
		{Key: "osmosis-1", Value: 0x01},
	}, prefixes)
}
