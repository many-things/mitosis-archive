package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/x/multisig/types"
)

type KeygenRepo interface {
	Load(id uint64) (*types.Keygen, error)
	Save(keygen *types.Keygen) error
	Delete(id uint64) error

	Paginate(page *query.PageRequest) ([]types.KV[uint64, *types.Keygen], *query.PageResponse, error)

	// TODO: Implement Genesis Tool
}

type kvKeygenRepo struct {
	cdc     codec.BinaryCodec
	root    store.KVStore
	chainID string
}

func NewKVChainKeygenRepo(cdc codec.BinaryCodec, root store.KVStore, chainID string) KeygenRepo {
	return &kvKeygenRepo{
		cdc, root, chainID,
	}
}

var (
	kvKeygenRepoLatestId    = []byte{0x01}
	kvKeygenRepoItemsPrefix = []byte{0x02}
)

func (r kvKeygenRepo) getPrefix(prefix []byte) []byte {
	return append([]byte(r.chainID), prefix...)
}

func (r kvKeygenRepo) Load(id uint64) (*types.Keygen, error) {
	bz := prefix.NewStore(r.root, r.getPrefix(kvKeygenRepoItemsPrefix)).Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return nil, nil
	}

	keygen := new(types.Keygen)
	if err := keygen.Unmarshal(bz); err != nil {
		return nil, err
	}

	return keygen, nil
}

func (r kvKeygenRepo) Save(keygen *types.Keygen) error {
	latestId := sdk.BigEndianToUint64(r.root.Get(r.getPrefix(kvKeygenRepoLatestId)))
	latestIdBz := sdk.Uint64ToBigEndian(latestId)

	keygen.KeyId = latestId
	keygenBz, err := keygen.Marshal()
	if err != nil {
		return err
	}

	prefix.NewStore(r.root, r.getPrefix(kvKeygenRepoItemsPrefix)).Set(latestIdBz, keygenBz)
	r.root.Set(r.getPrefix(kvKeygenRepoLatestId), sdk.Uint64ToBigEndian(latestId+1))

	return nil
}

func (r kvKeygenRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, r.getPrefix(kvKeygenRepoItemsPrefix))
	bz := ks.Get(sdk.Uint64ToBigEndian(id))

	// check for obj is exists and valid
	var keygen types.Keygen
	if err := keygen.Unmarshal(bz); err != nil {
		return err
	}

	ks.Delete(sdk.Uint64ToBigEndian(id))
	return nil
}

func (r kvKeygenRepo) Paginate(page *query.PageRequest) ([]types.KV[uint64, *types.Keygen], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, r.getPrefix(kvKeygenRepoItemsPrefix))

	var results []types.KV[uint64, *types.Keygen]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		keygen := new(types.Keygen)

		if err := keygen.Unmarshal(value); err != nil {
			return err
		}
		results = append(results, types.NewKV(sdk.BigEndianToUint64(key), keygen))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
