package store

import (
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

type ValidatorProxyRepo interface {
	Store(validator sdk.ValAddress, proxy sdk.AccAddress) error
	Get(validator sdk.ValAddress) (sdk.AccAddress, error)
}

type validatorProxyRepo struct{ store.KVStore }

func NewValidatorProxyRepo(ctx sdk.Context, key storetypes.StoreKey) ValidatorProxyRepo {
	return validatorProxyRepo{ValidatorProxyStore(ctx, key)}
}

func (s validatorProxyRepo) Store(validator sdk.ValAddress, proxy sdk.AccAddress) error {
	s.KVStore.Set(validator.Bytes(), proxy.Bytes())
	return nil
}

func (s validatorProxyRepo) Get(validator sdk.ValAddress) (sdk.AccAddress, error) {
	if !s.KVStore.Has(validator.Bytes()) {
		return sdk.AccAddress{}, errors.ErrNotFound
	}
	return s.KVStore.Get(validator), nil
}
