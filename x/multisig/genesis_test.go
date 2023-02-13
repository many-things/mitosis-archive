package multisig_test

import (
	"testing"

	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/testutil/nullify"
	"github.com/many-things/mitosis/x/multisig"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MultisigKeeper(t)
	multisig.InitGenesis(ctx, *k, genesisState)
	got := multisig.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
