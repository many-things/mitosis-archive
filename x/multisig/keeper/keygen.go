package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

type KeygenKeeper interface {
	RegisterKeygenEvent(ctx sdk.Context, chainId string, keygen *types.Keygen) (uint64, error)

	QueryKeygen(ctx sdk.Context, chainId string, id uint64) (*types.Keygen, error)
	QueryKeygenList(ctx sdk.Context, chainId string, pageReq *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error)
}

// RegisterKeygenEvent is register keygen into chainId's KV store
func (k keeper) RegisterKeygenEvent(ctx sdk.Context, chainId string, keygen *types.Keygen) (uint64, error) {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	keygenId, err := keygenRepo.Save(keygen)
	if err != nil {
		return 0, err
	}

	return keygenId, nil
}

// QueryKeygen is fetch one keygen event from chain
func (k keeper) QueryKeygen(ctx sdk.Context, chainId string, id uint64) (*types.Keygen, error) {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	keygen, err := keygenRepo.Load(id)
	if err != nil {
		return nil, err
	}

	return keygen, nil
}

// QueryKeygenList is fetch multiple keygens from chain
func (k keeper) QueryKeygenList(ctx sdk.Context, chainId string, pageReq *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error) {
	keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	results, pageResp, err := keygenRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
