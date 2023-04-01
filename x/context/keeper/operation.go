package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/keeper/state"
	"github.com/many-things/mitosis/x/context/types"
)

var _ types.OperationKeeper = &keeper{}

// InitOperation generate operation. I guess ids means Poll id.
func (k keeper) InitOperation(ctx sdk.Context, chain string, ids []uint64) (uint64, error) {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	op := types.Operation{
		Chain:  chain,
		Id:     0, // go filled by Load
		EvtIds: ids,
		Status: types.Operation_StatusPending,
	}

	opID, err := opRepo.Create(&op)
	if err != nil {
		return 0, err // TODO: require wrap errors
	}

	return opID, nil
}

// StartKeygenOperation (ctx sdk.Context, id uint64)
func (k keeper) StartKeygenOperation(_ sdk.Context, _ uint64) error {
	// TODO: StartKeygenOperation is required for this ?

	return nil
}

// FinishKeygenOperation (ctx sdk.Context, id uint64)
func (k keeper) FinishKeygenOperation(_ sdk.Context, _ uint64) error {
	return nil
}

func (k keeper) StartSignOperation(ctx sdk.Context, id uint64) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	op, err := opRepo.Load(id)
	if err != nil {
		return err // TODO: require wrap whole errors
	}

	op.Status = types.Operation_StatusInitSign
	err = opRepo.Save(op)

	if err != nil {
		return err // TODO: require wrap error
	}

	return nil
}

func (k keeper) FinishSignOperation(ctx sdk.Context, id uint64) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	op, err := opRepo.Load(id)
	if err != nil {
		return nil
	}

	op.Status = types.Operation_StatusFinishSign
	err = opRepo.Save(op)

	if err != nil {
		return err // TODO: require wrap error
	}

	return nil
}

// FinishOperation (ctx sdk.Context, id uint64, receipt []uint64)
func (k keeper) FinishOperation(ctx sdk.Context, id uint64, _ []uint64) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	op, err := opRepo.Load(id)
	if err != nil {
		return err
	}

	op.Status = types.Operation_StatusFinishSign
	err = opRepo.Save(op)

	if err != nil {
		return err
	}

	// TODO: save receipt

	return nil
}
