package keeper

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper/store"
)

type ValidatorKeeper interface {
	RegisterProxy(ctx sdk.Context, validator sdk.ValAddress, proxyAddr sdk.AccAddress) error

	GetRegisteredProxy(ctx sdk.Context, validator sdk.ValAddress) (sdk.AccAddress, error)
}

type validatorKeeper struct {
	storeKey storetypes.StoreKey
}

func newValidatorKeeper(storeKey storetypes.StoreKey) ValidatorKeeper {
	return validatorKeeper{storeKey}
}

func (k validatorKeeper) RegisterProxy(ctx sdk.Context, validator sdk.ValAddress, proxyAddr sdk.AccAddress) error {
	valStore := store.NewValidatorProxyRepo(ctx, k.storeKey)

	if err := valStore.Store(validator, proxyAddr); err != nil {
		return err
	}

	return nil
}

func (k validatorKeeper) GetRegisteredProxy(ctx sdk.Context, validator sdk.ValAddress) (sdk.AccAddress, error) {
	valStore := store.NewValidatorProxyRepo(ctx, k.storeKey)

	proxy, err := valStore.Get(validator)
	if err != nil {
		return nil, err
	}

	return proxy, nil
}
