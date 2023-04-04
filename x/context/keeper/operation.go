package keeper

import (
	sdkerrutils "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/pkg/txconv"
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

	signer, err := signerRepo.LoadByChain(chain)
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
	}

	opID, err := opRepo.Create(&op)
	if err != nil {
		return 0, sdkerrutils.Wrap(sdkerrors.ErrPanic, "create operation")
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

	if err = opRepo.Save(op); err != nil {
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

	res := poll.GetPayload().GetRes()
	if res == nil {
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
