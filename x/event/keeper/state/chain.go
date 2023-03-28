package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/types"
)

type ChainRepo interface {
	Load(chain string) (byte, error)

	Save(chain string) (byte, error)

	Delete(chain string) error

	Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[string, byte], *query.PageResponse, error)

	// ExportGenesis returns the entire module's state
	ExportGenesis() (genState *types.GenesisChain, err error)

	// ImportGenesis sets the entire module's state
	ImportGenesis(genState *types.GenesisChain) error
}

var (
	kvChainRepoKeyLatestID = []byte{0x01}
	kvChainRepoPrefixItems = []byte{0x02}
)

type kvChainRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVChainRepo(cdc codec.BinaryCodec, root store.KVStore) ChainRepo {
	return &kvChainRepo{
		cdc:  cdc,
		root: prefix.NewStore(root, kvChainRepoKey),
	}
}

func (k kvChainRepo) Load(chain string) (byte, error) {
	ps := prefix.NewStore(k.root, kvChainRepoPrefixItems)

	v := ps.Get([]byte(chain))
	if v == nil {
		return 0x0, errors.ErrKeyNotFound
	}

	return v[0], nil
}

func (k kvChainRepo) Save(chain string) (byte, error) {
	v := k.root.Get(kvChainRepoKeyLatestID)
	if v == nil {
		v = []byte{0x00}
	}

	prefix.NewStore(k.root, kvChainRepoPrefixItems).Set([]byte(chain), v)
	k.root.Set(kvChainRepoKeyLatestID, []byte{v[0] + 0x01})

	return v[0], nil
}

func (k kvChainRepo) Delete(chain string) error {
	ps := prefix.NewStore(k.root, kvChainRepoPrefixItems)
	if !ps.Has([]byte(chain)) {
		return errors.ErrKeyNotFound
	}

	ps.Delete([]byte(chain))
	return nil
}

func (k kvChainRepo) Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[string, byte], *query.PageResponse, error) {
	ps := prefix.NewStore(k.root, kvChainRepoPrefixItems)

	var rs []mitotypes.KV[string, byte]
	pageResp, err := query.Paginate(ps, pageReq, func(key []byte, value []byte) error {
		rs = append(rs, mitotypes.NewKV(string(key), value[0]))
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return rs, pageResp, nil
}

func (k kvChainRepo) ExportGenesis() (*types.GenesisChain, error) {
	latestID := k.root.Get(kvChainRepoKeyLatestID)

	var items []*types.GenesisChain_ItemSet
	_, err := query.Paginate(
		prefix.NewStore(k.root, kvChainRepoPrefixItems),
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, value []byte) error {
			items = append(
				items,
				&types.GenesisChain_ItemSet{
					Chain:  string(key),
					Prefix: value,
				},
			)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}

	return &types.GenesisChain{
		LatestId: latestID,
		ItemSet:  items,
	}, nil
}

func (k kvChainRepo) ImportGenesis(genState *types.GenesisChain) error {
	if genState.GetLatestId() == nil {
		k.root.Set(kvChainRepoKeyLatestID, []byte{0x0})
	} else {
		k.root.Set(kvChainRepoKeyLatestID, genState.GetLatestId())
	}

	ps := prefix.NewStore(k.root, kvChainRepoPrefixItems)
	for _, item := range genState.GetItemSet() {
		ps.Set([]byte(item.Chain), item.Prefix)
	}

	return nil
}
