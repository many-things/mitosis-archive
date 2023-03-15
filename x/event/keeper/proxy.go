package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
)

var _ ProxyKeeper = Keeper{}

type ProxyKeeper interface {
	RegisterProxy(ctx sdk.Context, val sdk.ValAddress, prx sdk.AccAddress) error
	ClearProxy(ctx sdk.Context, val sdk.ValAddress) error
}

func (k Keeper) RegisterProxy(ctx sdk.Context, val sdk.ValAddress, prx sdk.AccAddress) error {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	if err := proxyRepo.Save(val, prx); err != nil {
		return err
	}

	return nil
}

func (k Keeper) ClearProxy(ctx sdk.Context, val sdk.ValAddress) error {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	if err := proxyRepo.Delete(val); err != nil {
		return err
	}

	return nil
}

func (k Keeper) QueryProxy(ctx sdk.Context, val sdk.ValAddress) (sdk.AccAddress, error) {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	acc, err := proxyRepo.Load(val)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (k Keeper) QueryProxies(ctx sdk.Context, pageReq *query.PageRequest) ([]mitotypes.KV[sdk.ValAddress, sdk.AccAddress], *query.PageResponse, error) {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	accs, pageResp, err := proxyRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return accs, pageResp, nil
}
