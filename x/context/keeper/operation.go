package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/types"
)

var _ types.OperationKeeper = &keeper{}

func (k keeper) InitOperation(ctx sdk.Context, chain string, ids []uint64) (uint64, error) {
	return 0, nil
}

func (k keeper) StartKeygenOperation(ctx sdk.Context, id uint64) error {
	return nil
}

func (k keeper) FinishKeygenOperation(ctx sdk.Context, id uint64) error {
	return nil
}

func (k keeper) FinishOperation(ctx sdk.Context, id uint64, receipt []uint64) error {
	return nil
}
