package keeper

import (
	sdkerrutils "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/txconv"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/keeper/state"
	"github.com/many-things/mitosis/x/context/types"
	evttypes "github.com/many-things/mitosis/x/event/types"
)

var _ types.OperationKeeper = &keeper{}

func (k keeper) InitOperation(ctx sdk.Context, chain string, poll *evttypes.Poll) (uint64, error) {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))
	signerRepo := state.NewKVSignerRepo(k.cdc, ctx.KVStore(k.storeKey))

	req := poll.GetPayload().GetReq()
	if req == nil {
		return 0, sdkerrutils.Wrap(sdkerrors.ErrPanic, "invalid event payload type")
	}

	signer, err := signerRepo.Load(chain)
	if err != nil {
		return 0, sdkerrutils.Wrapf(sdkerrors.ErrNotFound, "signer not found for chain %s", chain)
	}

	txPayload, txBytesToSign, err := txconv.Converter.Convert(
		signer.TxConvSigner(),
		req.DestChain, req.OpId, req.OpArgs...,
	)
	if err != nil {
		return 0, sdkerrutils.Wrap(sdkerrors.ErrPanic, "convert event to sign target")
	}

	op := types.Operation{
		Chain:         chain,
		ID:            0, // go filled by Load
		PollID:        poll.GetId(),
		Status:        types.Operation_StatusPending,
		TxPayload:     txPayload,
		TxBytesToSign: txBytesToSign,
		Result:        nil,
	}

	opID, err := opRepo.Create(&op)
	if err != nil {
		return 0, sdkerrutils.Wrap(sdkerrors.ErrPanic, "create operation")
	}

	return opID, nil
}

func (k keeper) StartSignOperation(ctx sdk.Context, id, sigID uint64) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	op, err := opRepo.Load(id)
	if err != nil {
		return err // TODO: require wrap whole errors
	}

	op.Status = types.Operation_StatusInitSign
	op.SigID = sigID

	if err := opRepo.Save(op); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
	}
	if err = opRepo.Shift(op.ID, types.Operation_StatusInitSign); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
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

	if err := opRepo.Save(op); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
	}
	if err := opRepo.Shift(op.ID, types.Operation_StatusFinishSign); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
	}

	return nil
}

func (k keeper) FinishOperation(ctx sdk.Context, id uint64, poll *evttypes.Poll) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	res := poll.GetPayload().GetRes()
	if res == nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "invalid event payload type")
	}

	op, err := opRepo.Load(id)
	if err != nil {
		return err
	}

	op.Status = types.Operation_StatusFinalized
	op.Result = &types.OperationResult{
		Ok:     res.Ok,
		Result: res.Result,
	}
	if err := opRepo.Save(op); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
	}
	if err := opRepo.Shift(op.ID, types.Operation_StatusFinalized); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
	}

	return nil
}

func (k keeper) QueryOperation(ctx sdk.Context, id uint64) (*types.Operation, error) {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	return opRepo.Load(id)
}

func (k keeper) QueryOperations(ctx sdk.Context, pageReq *query.PageRequest) ([]*types.Operation, *query.PageResponse, error) {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	ops, pageResp, err := opRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	rtn := mitotypes.MapKV(
		ops,
		func(_ uint64, v *types.Operation, _ int) *types.Operation { return v },
	)

	return rtn, pageResp, nil
}

func (k keeper) QueryOperationsByStatus(ctx sdk.Context, status types.Operation_Status, pageReq *query.PageRequest) ([]*types.Operation, *query.PageResponse, error) {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	ops, pageResp, err := opRepo.PaginateStatus(status, pageReq)
	if err != nil {
		return nil, nil, err
	}

	rtn := mitotypes.MapKV(
		ops,
		func(_ uint64, v *types.Operation, _ int) *types.Operation { return v },
	)

	return rtn, pageResp, nil
}
