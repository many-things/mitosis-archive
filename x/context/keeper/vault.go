package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/keeper/state"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/pkg/errors"
)

var _ types.VaultKeeper = keeper{}

func (k keeper) RegisterVault(ctx sdk.Context, chain string, vault string) error {
	vaultRepo := state.NewKVVaultRepo(k.cdc, ctx.KVStore(k.storeKey))

	if err := vaultRepo.Save(chain, vault); err != nil {
		return errors.Wrap(err, "save vault")
	}

	return nil
}

func (k keeper) ClearVault(ctx sdk.Context, chain string) error {
	vaultRepo := state.NewKVVaultRepo(k.cdc, ctx.KVStore(k.storeKey))

	if err := vaultRepo.Clear(chain); err != nil {
		return errors.Wrap(err, "clear vault")
	}

	return nil
}

func (k keeper) QueryVault(ctx sdk.Context, chain string) (string, error) {
	vaultRepo := state.NewKVVaultRepo(k.cdc, ctx.KVStore(k.storeKey))

	vault, err := vaultRepo.Load(chain)
	if err != nil {
		return "", errors.Wrap(err, "load vault")
	}

	return vault, nil
}

func (k keeper) QueryVaults(ctx sdk.Context, page *query.PageRequest) ([]mitotypes.KV[string, string], *query.PageResponse, error) {
	vaultRepo := state.NewKVVaultRepo(k.cdc, ctx.KVStore(k.storeKey))

	vaults, pageResp, err := vaultRepo.List(page)
	if err != nil {
		return nil, nil, errors.Wrap(err, "list vaults")
	}

	return vaults, pageResp, nil
}
