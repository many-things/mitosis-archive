package context_test

import (
	"testing"

	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/testutil/nullify"
	"github.com/many-things/mitosis/x/context"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()

	k, ctx, _, _, _ := keepertest.ContextKeeper(t)
	context.InitGenesis(ctx, k, *genesisState)
	got := context.ExportGenesis(ctx, k, []string{})
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
