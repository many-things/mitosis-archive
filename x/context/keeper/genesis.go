package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/keeper/state"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/pkg/errors"
)

var _ types.GenesisKeeper = &keeper{}

func (k keeper) ExportGenesis(ctx sdk.Context, chains []string) (genesis *types.GenesisState, err error) {
	genesis = &types.GenesisState{}

	genesis.Params = k.GetParams(ctx)

	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))
	opRepoGen, err := opRepo.ExportGenesis()
	if err != nil {
		return nil, errors.Wrap(err, "export operation repo genesis")
	}
	genesis.Operation = opRepoGen

	for _, chain := range chains {
		opIdxRepo := state.NewKVOperationHashIndexRepo(k.cdc, ctx.KVStore(k.storeKey), chain)
		opIdxRepoGen, err := opIdxRepo.ExportGenesis()
		if err != nil {
			return nil, errors.Wrap(err, "export operation index repo genesis")
		}
		genesis.OperationIdx.ChainSet = append(genesis.OperationIdx.ChainSet, opIdxRepoGen)
	}

	vaultRepo := state.NewKVVaultRepo(k.cdc, ctx.KVStore(k.storeKey))
	vaultRepoGen, err := vaultRepo.ExportGenesis()
	if err != nil {
		return nil, errors.Wrap(err, "export vault repo genesis")
	}
	genesis.Vault = vaultRepoGen

	return
}

func (k keeper) ImportGenesis(ctx sdk.Context, genesis *types.GenesisState) (err error) {
	k.SetParams(ctx, genesis.Params)

	opRepo := state.NewKVOperationRepo(k.cdc, ctx.KVStore(k.storeKey))
	if err = opRepo.ImportGenesis(genesis.Operation); err != nil {
		return errors.Wrap(err, "import operation repo genesis")
	}

	for _, chain := range genesis.OperationIdx.ChainSet {
		opIdxRepo := state.NewKVOperationHashIndexRepo(k.cdc, ctx.KVStore(k.storeKey), chain.ChainID)
		if err = opIdxRepo.ImportGenesis(chain); err != nil {
			return errors.Wrap(err, "import operation index repo genesis")
		}
	}

	vaultRepo := state.NewKVVaultRepo(k.cdc, ctx.KVStore(k.storeKey))
	if err = vaultRepo.ImportGenesis(genesis.Vault); err != nil {
		return errors.Wrap(err, "import vault repo genesis")
	}

	return
}
