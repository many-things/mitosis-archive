package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

// RegisterSignEvent is register new SignEvent
func (k keeper) RegisterSignEvent(ctx sdk.Context, chainID string, sign *exported.Sign) (uint64, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	sigID, err := signRepo.Create(sign)
	if err != nil {
		return 0, err
	}

	event := types.EventSigningStart{
		Chain:         sign.Chain,
		SigId:         sign.SigID,
		KeyId:         sign.KeyID,
		OpId:          sign.OpId,
		Participants:  sign.Participants,
		MessageToSign: sign.MessageToSign,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		return 0, err
	}

	return sigID, nil
}

// RemoveSignEvent is remove specific Sign Event
func (k keeper) RemoveSignEvent(ctx sdk.Context, chainID string, id uint64) error {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	if err := signRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

// UpdateSignStatus is update status of specific Event
func (k keeper) UpdateSignStatus(ctx sdk.Context, chainID string, id uint64, newStatus exported.Sign_Status) (*exported.Sign, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	sign, err := signRepo.Load(id)
	if err != nil {
		return nil, err
	}

	sign.Status = newStatus
	err = signRepo.Save(sign)

	if err != nil {
		return nil, err
	}

	return sign, nil
}

func (k keeper) SetResultSignature(ctx sdk.Context, chainID string, sigID uint64, signature exported.Signature) error {
	signRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	signResult, err := signRepo.Load(sigID)
	if err != nil {
		return err
	}
	signResult.ResultSignature = signature
	if err := signRepo.Save(signResult); err != nil {
		return err
	}

	return nil
}

// QuerySign is get specific Sign instance
func (k keeper) QuerySign(ctx sdk.Context, chainID string, id uint64) (*exported.Sign, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	sign, err := signRepo.Load(id)
	if err != nil {
		return nil, err
	}

	return sign, nil
}

// QuerySignList returns sign list of specific chain
func (k keeper) QuerySignList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.Sign], *query.PageResponse, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	results, pageResp, err := signRepo.Paginate(page)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
