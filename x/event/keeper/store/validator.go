package store

import (
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/event/types"
)

type ValidatorProxyRepo interface {
	Store(validator sdk.ValAddress, proxy sdk.AccAddress) error
	Get(validator sdk.ValAddress) (sdk.AccAddress, error)
	Iter(start, end []byte, handle func(validator sdk.ValAddress, proxy sdk.AccAddress) (stop bool))
}

type validatorProxyRepo struct{ store.KVStore }

func NewValidatorProxyRepo(ctx sdk.Context, key storetypes.StoreKey) ValidatorProxyRepo {
	return validatorProxyRepo{prefix.NewStore(
		ctx.KVStore(key),
		types.GetValidatorProxyPrefix(),
	)}
}

func (s validatorProxyRepo) Store(validator sdk.ValAddress, proxy sdk.AccAddress) error {
	s.KVStore.Set(validator, proxy)
	return nil
}

func (s validatorProxyRepo) Get(validator sdk.ValAddress) (sdk.AccAddress, error) {
	if !s.KVStore.Has(validator.Bytes()) {
		return sdk.AccAddress{}, errors.ErrNotFound
	}
	return s.KVStore.Get(validator), nil
}

func (s validatorProxyRepo) Iter(start, end []byte, handle func(validator sdk.ValAddress, proxy sdk.AccAddress) (stop bool)) {
	iter := s.Iterator(start, end)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		if stop := handle(iter.Key(), iter.Value()); stop {
			return
		}
	}
}
