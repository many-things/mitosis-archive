package multisig

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/pkg/errors"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.ImportGenesis(ctx, &genState); err != nil {
		panic(errors.Wrap(err, "failed to import genesis"))
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper, chains []byte) *types.GenesisState {
	genesis, err := k.ExportGenesis(ctx, chains)
	if err != nil {
		panic(errors.Wrap(err, "failed to export genesis"))
	}

	return genesis
}
