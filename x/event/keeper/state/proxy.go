package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/types"
)

type ProxyRepo interface {
	Load(validator sdk.ValAddress) (sdk.AccAddress, error)

	LoadByProxy(proxyAccount sdk.AccAddress) (sdk.ValAddress, error)

	Save(validator sdk.ValAddress, proxy sdk.AccAddress) error

	Delete(validator sdk.ValAddress) error

	Paginate(page *query.PageRequest) ([]mitotypes.KV[sdk.ValAddress, sdk.AccAddress], *query.PageResponse, error)

	// ExportGenesis returns the entire module's state
	ExportGenesis() (genState *types.GenesisProxy, err error)

	// ImportGenesis sets the entire module's state
	ImportGenesis(genState *types.GenesisProxy) error
}

var (
	kvProxyRepoItemsPrefix        = []byte{0x01}
	kvProxyRepoItemsReversePrefix = []byte{0x02}
)

type kvProxyRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVProxyRepo(cdc codec.BinaryCodec, store store.KVStore) ProxyRepo {
	return kvProxyRepo{cdc, store}
}

func (k kvProxyRepo) Load(validator sdk.ValAddress) (sdk.AccAddress, error) {
	return prefix.NewStore(k.root, kvProxyRepoItemsPrefix).Get(validator.Bytes()), nil
}

func (k kvProxyRepo) LoadByProxy(proxyAccount sdk.AccAddress) (sdk.ValAddress, error) {
	return prefix.NewStore(k.root, kvProxyRepoItemsReversePrefix).Get(proxyAccount.Bytes()), nil
}

func (k kvProxyRepo) Save(validator sdk.ValAddress, proxy sdk.AccAddress) error {
	prefix.NewStore(k.root, kvProxyRepoItemsPrefix).Set(validator.Bytes(), proxy.Bytes())
	prefix.NewStore(k.root, kvProxyRepoItemsReversePrefix).Set(proxy.Bytes(), validator.Bytes())
	return nil
}

func (k kvProxyRepo) Delete(validator sdk.ValAddress) error {
	ps := prefix.NewStore(k.root, kvProxyRepoItemsPrefix)
	proxy := ps.Get(validator)
	if proxy == nil {
		return errors.ErrKeyNotFound
	}

	ps.Delete(validator)
	prefix.NewStore(k.root, kvProxyRepoItemsReversePrefix).Delete(proxy)
	return nil
}

func (k kvProxyRepo) Paginate(page *query.PageRequest) ([]mitotypes.KV[sdk.ValAddress, sdk.AccAddress], *query.PageResponse, error) {
	ps := prefix.NewStore(k.root, kvProxyRepoItemsPrefix)

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

func (k kvProxyRepo) ExportGenesis() (genState *types.GenesisProxy, err error) {
	genState = &types.GenesisProxy{}

	_, err = query.Paginate(
		prefix.NewStore(k.root, kvProxyRepoItemsPrefix),
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, value []byte) error {
			genState.ItemSet = append(
				genState.ItemSet,
				&types.GenesisProxy_ItemSet{
					Validator:    key,
					ProxyAccount: value,
				},
			)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return
}

func (k kvProxyRepo) ImportGenesis(genState *types.GenesisProxy) error {
	ps := prefix.NewStore(k.root, kvProxyRepoItemsPrefix)
	psr := prefix.NewStore(k.root, kvProxyRepoItemsReversePrefix)
	for _, item := range genState.GetItemSet() {
		ps.Set(item.GetValidator().Bytes(), item.GetProxyAccount().Bytes())
		psr.Set(item.GetProxyAccount().Bytes(), item.GetValidator().Bytes())
	}

	return nil
}
