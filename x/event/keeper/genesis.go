package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

// determinism
var _ types.GenesisKeeper = keeper{}

func (k keeper) ExportGenesis(ctx sdk.Context) (genesis *types.GenesisState, err error) {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	genesis = types.DefaultGenesis()

	// export params
	genesis.Params = k.GetParams(ctx)

	// export proxies
	if genesis.Proxy, err = proxyRepo.ExportGenesis(); err != nil {
		return nil, err
	}

	// export chains
	if genesis.Chain, err = chainRepo.ExportGenesis(); err != nil {
		return nil, err
	}

	// export polls
	for _, v := range genesis.GetChain().GetItemSet() {
		pollRepo := state.NewKVPollRepo(k.cdc, v.GetPrefix()[0], ctx.KVStore(k.storeKey))
		poll, err := pollRepo.ExportGenesis()
		if err != nil {
			return nil, err
		}

		genesis.Poll.ChainSet = append(genesis.Poll.ChainSet, poll)
	}

	// export snapshot
	if genesis.Snapshot, err = snapshotRepo.ExportGenesis(); err != nil {
		return nil, err
	}

	return
}

func (k keeper) ImportGenesis(ctx sdk.Context, genesis *types.GenesisState) error {
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	k.SetParams(ctx, genesis.Params)

	for _, v := range genesis.GetPoll().GetChainSet() {
		pollRepo := state.NewKVPollRepo(k.cdc, v.GetChain()[0], ctx.KVStore(k.storeKey))
		if err := pollRepo.ImportGenesis(v); err != nil {
			return err
		}
	}

	if err := proxyRepo.ImportGenesis(genesis.GetProxy()); err != nil {
		return err
	}

	if err := chainRepo.ImportGenesis(genesis.GetChain()); err != nil {
		return err
	}

	if err := snapshotRepo.ImportGenesis(genesis.GetSnapshot()); err != nil {
		return err
	}

	return nil
}
