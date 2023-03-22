package event

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k types.GenesisKeeper, genState types.GenesisState) {
	if err := k.ImportGenesis(ctx, &genState); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k types.GenesisKeeper) *types.GenesisState {
	genesis, err := k.ExportGenesis(ctx)
	if err != nil {
		panic(err)
	}
	return genesis
}
