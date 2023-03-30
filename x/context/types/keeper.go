package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params

	SetParams(ctx sdk.Context, params Params)
}

type OperationKeeper interface {
}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context) (genesis *GenesisState, err error)
	ImportGenesis(ctx sdk.Context, genesis *GenesisState) error
}
