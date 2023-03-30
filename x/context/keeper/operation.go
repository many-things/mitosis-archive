package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/types"
)

var _ types.OperationKeeper = &keeper{}

func (k keeper) SubmitOperation(ctx sdk.Context, chain string, ids []uint64) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (k keeper) UpdateOperationStatus(ctx sdk.Context, id uint64, status types.Operation_Status) error {
	//TODO implement me
	panic("implement me")
}
