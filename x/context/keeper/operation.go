package keeper

import (
	"encoding/base64"

	sdkerrutils "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/msgconv"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/keeper/state"
	"github.com/many-things/mitosis/x/context/types"
	evttypes "github.com/many-things/mitosis/x/event/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/pkg/errors"
)

var _ types.OperationKeeper = &keeper{}

// ctx, chainID, poll
func (k keeper) InitOperation(ctx sdk.Context, _ string, poll *evttypes.Poll) (uint64, error) {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))
	opHashIndexRepo := state.NewKVOperationHashIndexRepo(k.cdc, ctx.KVStore(k.storeKey), poll.Chain)
	vaultRepo := state.NewKVVaultRepo(k.cdc, ctx.KVStore(k.storeKey))

	req := poll.GetPayload().GetReq()
	if req == nil {
		return 0, sdkerrutils.Wrap(sdkerrors.ErrPanic, "invalid event payload type")
	}

	vault, err := vaultRepo.Load(req.DestChain)
	if err != nil {
		return 0, errors.Wrap(err, "load vault")
	}

	txPayload, txBytesToSign, err := msgconv.ToMsgs(req.DestChain, vault, req.OpId, req.OpArgs, req.Funds)
	if err != nil {
		return 0, errors.Wrap(err, "convert to msgs")
	}

	op := types.Operation{
		Chain:         req.DestChain,
		ID:            0, // go filled by Load
		PollID:        poll.GetId(),
		Status:        types.Operation_StatusPending,
		SignerPubkey:  nil,
		TxPayload:     txPayload,
		TxBytesToSign: txBytesToSign,
		Result:        nil,
	}

	opID, err := opRepo.Create(&op)
	if err != nil {
		return 0, sdkerrutils.Wrapf(sdkerrors.ErrPanic, "create operation. err=%v", err)
	}

	err = opHashIndexRepo.Create(poll.Payload.TxHash, opID)
	if err != nil {
		return 0, sdkerrutils.Wrapf(err, "create operation index. err=%v", err)
	}

	txHash := base64.StdEncoding.EncodeToString(poll.GetPayload().TxHash)
	err = ctx.EventManager().EmitTypedEvent(
		&types.EventOperationInitialized{
			PollID:      poll.Id,
			OperationID: opID,
			ChainID:     poll.GetChain(),
			TxHash:      txHash,
		},
	)
	if err != nil {
		return 0, errors.Wrap(err, "emit event")
	}

	return opID, nil
}

func (k keeper) StartSignOperation(ctx sdk.Context, id, sigID uint64, pubkey exported.PublicKey) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	op, err := opRepo.Load(id)
	if err != nil {
		return err // TODO: require wrap whole errors
	}

	if err = opRepo.Shift(op.ID, types.Operation_StatusInitSign); err != nil {
		return sdkerrutils.Wrapf(sdkerrors.ErrPanic, "shift operation %v", err)
	}

	op.Status = types.Operation_StatusInitSign
	op.SignerPubkey = pubkey
	op.SigID = sigID

	if err := opRepo.Save(op); err != nil {
		return sdkerrutils.Wrapf(sdkerrors.ErrPanic, "save operation %v", err)
	}

	err = ctx.EventManager().EmitTypedEvent(
		&types.EventOperationSigningStarted{
			OperationID: op.ID,
			SignID:      op.SigID,
			Signer:      op.SignerPubkey,
		},
	)
	if err != nil {
		return errors.Wrap(err, "emit event")
	}

	return nil
}

func (k keeper) FinishSignOperation(ctx sdk.Context, id uint64, signature []byte) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))

	op, err := opRepo.Load(id)
	if err != nil {
		return err
	}

	op.Status = types.Operation_StatusFinishSign

	if err := opRepo.Save(op); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
	}
	if err := opRepo.Shift(op.ID, types.Operation_StatusFinishSign); err != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrPanic, "save operation")
	}

	emitEvent := &types.EventOperationSigningFinished{
		OperationID: op.ID,
		SignID:      op.SigID,
		Signer:      op.SignerPubkey,
		Signature:   signature,
		ChainID:     op.Chain,
	}
	if err := ctx.EventManager().EmitTypedEvent(emitEvent); err != nil {
		return errors.Wrap(err, "emit event")
	}

	return nil
}

func (k keeper) FinishOperation(ctx sdk.Context, id uint64, poll *evttypes.Poll) error {
	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))
	opHashIndexRepo := state.NewKVOperationHashIndexRepo(k.cdc, ctx.KVStore(k.storeKey), poll.Chain)

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

	err = opHashIndexRepo.Create(poll.Payload.TxHash, op.ID)
	if err != nil {
		return sdkerrutils.Wrap(err, "save tx payload")
	}

	err = ctx.EventManager().EmitTypedEvent(
		&types.EventOperationFinalized{
			OperationID: op.ID,
			ReqPollID:   op.PollID,
			RespPollID:  poll.Id,
		},
	)
	if err != nil {
		return errors.Wrap(err, "emit event")
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

func (k keeper) QueryOperationByHash(ctx sdk.Context, chain string, hash []byte) (*types.Operation, error) {
	idxRepo := state.NewKVOperationHashIndexRepo(k.cdc, ctx.KVStore(k.storeKey), chain)

	opID, err := idxRepo.Load(hash)
	if err != nil {
		return nil, err
	}

	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))
	op, err := opRepo.Load(opID)
	if err != nil {
		return nil, err
	}

	return op, nil
}
