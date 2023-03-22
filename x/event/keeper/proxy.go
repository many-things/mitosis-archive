package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

// determinism
var _ types.ProxyKeeper = keeper{}

func (k keeper) RegisterProxy(ctx sdk.Context, val sdk.ValAddress, prx sdk.AccAddress) error {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	if err := proxyRepo.Save(val, prx); err != nil {
		return err
	}

	return nil
}

func (k keeper) ClearProxy(ctx sdk.Context, val sdk.ValAddress) error {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	if err := proxyRepo.Delete(val); err != nil {
		return err
	}

	return nil
}

func (k keeper) QueryProxy(ctx sdk.Context, val sdk.ValAddress) (sdk.AccAddress, bool) {
	return state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey)).Load(val)
}

func (k keeper) QueryProxyReverse(ctx sdk.Context, prx sdk.AccAddress) (sdk.ValAddress, bool) {
	return state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey)).LoadByProxy(prx)
}

func (k keeper) QueryProxies(ctx sdk.Context, pageReq *query.PageRequest) ([]mitotypes.KV[sdk.ValAddress, sdk.AccAddress], *query.PageResponse, error) {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	accs, pageResp, err := proxyRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return accs, pageResp, nil
}
