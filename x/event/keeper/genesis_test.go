package keeper_test

import (
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/event/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesis(t *testing.T) {
	k, ctx := testkeeper.EventKeeper(t)

	// import
	genesis := types.DefaultGenesis()
	require.NoError(t, k.ImportGenesis(ctx, genesis))

	// export
	exported, err := k.ExportGenesis(ctx)
	require.NoError(t, err)

	// validation
	require.Equal(t, genesis, exported)
}
