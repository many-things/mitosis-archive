package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/pkg/errors"
)

type OperationHashIndexRepo interface {
	// Immutable
	Load(hash []byte) (uint64, error)
	Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[[]byte, uint64], *query.PageResponse, error)

	// Mutable
	Create(hash []byte, opID uint64) error

	// Genesis
	ExportGenesis() (*types.GenesisOperationHashIndex_ChainSet, error)
	ImportGenesis(genState *types.GenesisOperationHashIndex_ChainSet) error
}

var (
	kvOperationHashIndexRepoItemPrefix = []byte{0x01}
)

type kvOperationHashIndexRepo struct {
	cdc     codec.BinaryCodec
	root    store.KVStore
	chainID string
}

func NewKVOperationHashIndexRepo(cdc codec.BinaryCodec, store store.KVStore, chainID string) OperationHashIndexRepo {
	root := prefix.NewStore(store, append(kvOperationRepoIndexByStatusPrefix, []byte(chainID)...))

	return &kvOperationHashIndexRepo{
		cdc:     cdc,
		root:    root,
		chainID: chainID,
	}
}

func (r kvOperationHashIndexRepo) Load(hash []byte) (uint64, error) {
	bz := r.root.Get(hash)
	if bz == nil {
		return 0, errors.New("hash index not found")
	}

	return sdk.BigEndianToUint64(bz), nil
}

func (r kvOperationHashIndexRepo) Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[[]byte, uint64], *query.PageResponse, error) {
	var kvs []mitotypes.KV[[]byte, uint64]

	pageResp, err := query.Paginate(r.root, pageReq, func(id []byte, bz []byte) error {
		kvs = append(kvs, mitotypes.NewKV(id, sdk.BigEndianToUint64(bz)))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return kvs, pageResp, nil
}

func (r kvOperationHashIndexRepo) Create(hash []byte, opID uint64) error {
	r.root.Set(hash, sdk.Uint64ToBigEndian(opID))

	return nil
}

func (r kvOperationHashIndexRepo) ExportGenesis() (*types.GenesisOperationHashIndex_ChainSet, error) {
	genState := &types.GenesisOperationHashIndex_ChainSet{
		ChainID: []byte(r.chainID),
	}

	var state []*types.GenesisOperationHashIndex_ItemSet
	_, err := query.Paginate(
		r.root,
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, bz []byte) error {
			item := &types.GenesisOperationHashIndex_ItemSet{
				Hash: key,
				OpID: bz,
			}

			state = append(state, item)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	genState.ItemSet = state
	return genState, nil
}

func (r kvOperationHashIndexRepo) ImportGenesis(genState *types.GenesisOperationHashIndex_ChainSet) error {
	for _, item := range genState.ItemSet {
		r.root.Set(item.Hash, item.OpID)
	}

	return nil
}
