package multisig

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper, chains []byte) *types.GenesisState {
	genesis, err := k.ExportGenesis(ctx, chains)

	if err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
