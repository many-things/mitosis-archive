package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	evttypes "github.com/many-things/mitosis/x/event/types"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params

	SetParams(ctx sdk.Context, params Params)
}

type OperationKeeper interface {
	InitOperation(ctx sdk.Context, chain string, poll *evttypes.Poll) (uint64, error)

	StartSignOperation(ctx sdk.Context, id uint64) error

	FinishSignOperation(ctx sdk.Context, id uint64) error

	FinishOperation(ctx sdk.Context, id uint64, poll *evttypes.Poll) error
}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context) (genesis *GenesisState, err error)

	ImportGenesis(ctx sdk.Context, genesis *GenesisState) error
}
