package context

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/keeper"
	"github.com/many-things/mitosis/x/context/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.ImportGenesis(ctx, &genState); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper, chains []string) *types.GenesisState {
	genesis, err := k.ExportGenesis(ctx, chains)
	if err != nil {
		panic(err)
	}

	return genesis
}
