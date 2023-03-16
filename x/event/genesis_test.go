package event_test

import (
	"testing"

	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/testutil/nullify"
	"github.com/many-things/mitosis/x/event"
	"github.com/many-things/mitosis/x/event/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()

	k, ctx := keepertest.EventKeeper(t)
	event.InitGenesis(ctx, *k, *genesisState)
	got := event.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
