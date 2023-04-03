package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/types"
)

var _ types.GenesisKeeper = &keeper{}

func (k keeper) ExportGenesis(ctx sdk.Context) (genesis *types.GenesisState, err error) {
	return
}

func (k keeper) ImportGenesis(ctx sdk.Context, genesis *types.GenesisState) (err error) {
	return
}
