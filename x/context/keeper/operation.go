package keeper

import (
	sdkerrutils "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/pkg/txconv"
	txconvtypes "github.com/many-things/mitosis/pkg/txconv/types"
	"github.com/many-things/mitosis/x/context/keeper/state"
	"github.com/many-things/mitosis/x/context/types"
	evttypes "github.com/many-things/mitosis/x/event/types"
)

var _ types.OperationKeeper = &keeper{}

func (k keeper) InitOperation(ctx sdk.Context, chain string, poll *evttypes.Poll) (uint64, error) {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	wreq, ok := poll.GetPayload().GetEvent().(*evttypes.Event_Req)
	if !ok {
		return 0, sdkerrutils.Wrap(sdkerrors.ErrPanic, "invalid event payload type")
	}
	req := wreq.Req

	signer := txconvtypes.NewCosmosSigner(nil, "osmo", 0, 0)

	unsignedTx, bytesToSign, err := txconv.Converter.Convert(signer, req.DestChain, req.OpId, req.OpArgs...)
	if err != nil {
		return 0, err // TODO: require wrap errors
	}

	_ = unsignedTx
	_ = bytesToSign

	op := types.Operation{
		Chain:  chain,
		Id:     0, // go filled by Load
		EvtId:  poll.GetId(),
		Status: types.Operation_StatusPending,
	}

	opID, err := opRepo.Create(&op)
	if err != nil {
		return 0, err // TODO: require wrap errors
	}

	return opID, nil
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

func (k keeper) FinishOperation(ctx sdk.Context, id uint64, poll *evttypes.Poll) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	res, ok := poll.GetPayload().GetEvent().(*evttypes.Event_Res)
	if !ok {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "invalid event payload type")
	}
	_ = res

	// TODO: res -> archive and finalize

	op, err := opRepo.Load(id)
	if err != nil {
		return err
	}

	op.Status = types.Operation_StatusFinalized
	err = opRepo.Save(op)

	if err != nil {
		return err
	}

	// TODO: save receipt

	return nil
}
