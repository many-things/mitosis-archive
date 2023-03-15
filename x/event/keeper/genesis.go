package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

func (k Keeper) ExportGenesis(ctx sdk.Context) (genesis *types.GenesisState, err error) {
	pollRepo := state.NewKVPollRepo(k.cdc, ctx.KVStore(k.storeKey))
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

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

	return
}

func (k Keeper) ImportGenesis(ctx sdk.Context, genesis *types.GenesisState) error {
	pollRepo := state.NewKVPollRepo(k.cdc, ctx.KVStore(k.storeKey))
	proxyRepo := state.NewKVProxyRepo(k.cdc, ctx.KVStore(k.storeKey))

	k.SetParams(ctx, genesis.Params)

	if err := pollRepo.ImportGenesis(genesis.Poll); err != nil {
		return err
	}

	if err := proxyRepo.ImportGenesis(genesis.Proxy); err != nil {
		return err
	}

	return nil
}
