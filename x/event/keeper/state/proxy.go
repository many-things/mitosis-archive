package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

type ProxyRepo interface {
	Load(validator sdk.ValAddress) (sdk.AccAddress, error)

	Save(validator sdk.ValAddress, proxy sdk.AccAddress) error

	Delete(validator sdk.ValAddress) error

	Paginate(page *query.PageRequest) ([]mitotypes.KV[sdk.ValAddress, sdk.AccAddress], *query.PageResponse, error)
}

var (
	kvProxyRepoItemsPrefix = []byte{0x01}
)

type kvProxyRepo struct {
	cdc  codec.Codec
	root store.KVStore
}

func NewKVProxyRepo(cdc codec.Codec, store store.KVStore) ProxyRepo {
	return kvProxyRepo{cdc, store}
}

func (k kvProxyRepo) Load(validator sdk.ValAddress) (sdk.AccAddress, error) {
	return prefix.NewStore(k.root, kvProxyRepoItemsPrefix).Get(validator.Bytes()), nil
}

func (k kvProxyRepo) Save(validator sdk.ValAddress, proxy sdk.AccAddress) error {
	prefix.NewStore(k.root, kvProxyRepoItemsPrefix).Set(validator.Bytes(), proxy.Bytes())
	return nil
}

func (k kvProxyRepo) Delete(validator sdk.ValAddress) error {
	prefix.NewStore(k.root, kvProxyRepoItemsPrefix).Delete(validator.Bytes())
	return nil
}

func (k kvProxyRepo) Paginate(page *query.PageRequest) ([]mitotypes.KV[sdk.ValAddress, sdk.AccAddress], *query.PageResponse, error) {
	ps := prefix.NewStore(k.root, kvPollRepoItemsPrefix)

	var rs []mitotypes.KV[sdk.ValAddress, sdk.AccAddress]
	pageResp, err := query.Paginate(ps, page, func(key []byte, value []byte) error {
		var (
			val sdk.ValAddress = key
			prx sdk.AccAddress = value
		)
		rs = append(rs, mitotypes.NewKV(val, prx))
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return rs, pageResp, nil
}
