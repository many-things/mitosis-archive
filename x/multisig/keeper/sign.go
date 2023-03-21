package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

// RegisterSignEvent is register new SignEvent
func (k keeper) RegisterSignEvent(ctx sdk.Context, chainId string, sign *types.Sign) (uint64, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	sigId, err := signRepo.Create(sign)
	if err != nil {
		return 0, err
	}

	return sigId, nil
}

// RemoveSignEvent is remove specific Sign Event
func (k keeper) RemoveSignEvent(ctx sdk.Context, chainId string, id uint64) error {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	err := signRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSignStatus is update status of specific Event
func (k keeper) UpdateSignStatus(ctx sdk.Context, chainId string, id uint64, newStatus types.Sign_Status) (*types.Sign, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

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

// QuerySign is get specific Sign instance
func (k keeper) QuerySign(ctx sdk.Context, chainId string, id uint64) (*types.Sign, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	sign, err := signRepo.Load(id)
	if err != nil {
		return nil, err
	}

	return sign, nil
}

// QuerySignList returns sign list of specific chain
func (k keeper) QuerySignList(ctx sdk.Context, chainId string, page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Sign], *query.PageResponse, error) {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	results, pageResp, err := signRepo.Paginate(page)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
