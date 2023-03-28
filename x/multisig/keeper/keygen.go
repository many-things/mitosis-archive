package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

// RegisterKeygenEvent is register keygen into chainID's KV store
func (k keeper) RegisterKeygenEvent(ctx sdk.Context, chainID string, keygen *types.Keygen) (uint64, error) {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	keygenID, err := keygenRepo.Create(keygen)
	if err != nil {
		return 0, err
	}

	return keygenID, nil
}

// RemoveKeygenEvent is remove keygen
func (k keeper) RemoveKeygenEvent(ctx sdk.Context, chainID string, id uint64) error {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	if err := keygenRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

// UpdateKeygenStatus is change keygen status
func (k keeper) UpdateKeygenStatus(ctx sdk.Context, chainID string, id uint64, newStatus types.Keygen_Status) (*types.Keygen, error) {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	keygen, err := keygenRepo.Load(id)
	if err != nil {
		return nil, err
	}

	keygen.Status = newStatus
	err = keygenRepo.Save(keygen)

	if err != nil {
		return nil, err
	}

	return keygen, nil
}

// QueryKeygen is fetch one keygen event from chain
func (k keeper) QueryKeygen(ctx sdk.Context, chainID string, id uint64) (*types.Keygen, error) {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	keygen, err := keygenRepo.Load(id)
	if err != nil {
		return nil, err
	}

	return keygen, nil
}

// QueryKeygenList is fetch multiple keygens from chain
func (k keeper) QueryKeygenList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error) {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	results, pageResp, err := keygenRepo.Paginate(page)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
