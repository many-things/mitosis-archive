package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

// determinism
var _ types.ChainKeeper = keeper{}

func (k keeper) RegisterChain(ctx sdk.Context, chain string) (byte, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	if _, err := chainRepo.Load(chain); err == nil {
		return 0x0, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "chain %s already registered", chain)
	}

	prefix, err := chainRepo.Save(chain)
	if err != nil {
		return 0x0, err
	}

	return prefix, nil
}

func (k keeper) UnregisterChain(ctx sdk.Context, chain string) error {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	if _, err := chainRepo.Load(chain); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "chain %s is not registered", chain)
	}

	if err := chainRepo.Delete(chain); err != nil {
		return err
	}

	return nil
}

func (k keeper) QueryChain(ctx sdk.Context, chain string) (byte, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	prefix, err := chainRepo.Load(chain)
	if err != nil {
		return 0x0, err
	}

	return prefix, nil
}

func (k keeper) QueryChains(ctx sdk.Context, pageReq *query.PageRequest) ([]mitotypes.KV[string, byte], *query.PageResponse, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	chains, pageResp, err := chainRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return chains, pageResp, nil
}
