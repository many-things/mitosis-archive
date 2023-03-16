package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
)

var _ ChainKeeper = Keeper{}

type ChainKeeper interface {
	RegisterChain(ctx sdk.Context, chain string) (byte, error)

	UnregisterChain(ctx sdk.Context, chain string) error

	QueryChain(ctx sdk.Context, chain string) (byte, error)

	QueryChains(ctx sdk.Context, pageReq *query.PageRequest) ([]mitotypes.KV[string, byte], *query.PageResponse, error)
}

func (k Keeper) RegisterChain(ctx sdk.Context, chain string) (byte, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	prefix, err := chainRepo.Save(chain)
	if err != nil {
		return 0x0, err
	}

	return prefix, nil
}

func (k Keeper) UnregisterChain(ctx sdk.Context, chain string) error {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	if err := chainRepo.Delete(chain); err != nil {
		return err
	}

	return nil
}

func (k Keeper) QueryChain(ctx sdk.Context, chain string) (byte, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	prefix, err := chainRepo.Load(chain)
	if err != nil {
		return 0x0, err
	}

	return prefix, nil
}

func (k Keeper) QueryChains(ctx sdk.Context, pageReq *query.PageRequest) ([]mitotypes.KV[string, byte], *query.PageResponse, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	chains, pageResp, err := chainRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return chains, pageResp, nil
}
