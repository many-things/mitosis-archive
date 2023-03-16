package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

var _ GenesisKeeper = Keeper{}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context) (genesis *types.GenesisState, err error)
	ImportGenesis(ctx sdk.Context, genesis *types.GenesisState) error
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (genesis *types.GenesisState, err error) {
	defaultChain := byte(0x01) // TODO: make chain registry
	pollRepo := state.NewKVPollRepo(k.cdc, defaultChain, ctx.KVStore(k.storeKey))
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	genesis = types.DefaultGenesis()

	// export params
	genesis.Params = k.GetParams(ctx)

	// export polls
	if genesis.Poll, err = pollRepo.ExportGenesis(); err != nil {
		return nil, err
	}

	// export proxies
	if genesis.Proxy, err = proxyRepo.ExportGenesis(); err != nil {
		return nil, err
	}

	// export chains
	if genesis.Chain, err = chainRepo.ExportGenesis(); err != nil {
		return nil, err
	}

	return
}

func (k Keeper) ImportGenesis(ctx sdk.Context, genesis *types.GenesisState) error {
	defaultChain := byte(0x01) // TODO: make chain registry
	pollRepo := state.NewKVPollRepo(k.cdc, defaultChain, ctx.KVStore(k.storeKey))
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))

	k.SetParams(ctx, genesis.Params)

	if err := pollRepo.ImportGenesis(genesis.Poll); err != nil {
		return err
	}

	if err := proxyRepo.ImportGenesis(genesis.Proxy); err != nil {
		return err
	}

	if err := chainRepo.ImportGenesis(genesis.Chain); err != nil {
		return err
	}

	return nil
}
