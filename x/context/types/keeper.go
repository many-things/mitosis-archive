package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params

	SetParams(ctx sdk.Context, params Params)
}

type OperationKeeper interface {
	InitOperation(ctx sdk.Context, chain string, ids []uint64) (uint64, error)

	StartKeygenOperation(ctx sdk.Context, id uint64) error

	FinishKeygenOperation(ctx sdk.Context, id uint64) error

	FinishOperation(ctx sdk.Context, id uint64, receipt []uint64) error
}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context) (genesis *GenesisState, err error)

	ImportGenesis(ctx sdk.Context, genesis *GenesisState) error
}
