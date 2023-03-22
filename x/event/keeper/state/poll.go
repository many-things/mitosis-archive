package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/types"
)

type PollRepo interface {
	Load(id uint64) (*types.Poll, error)
	LoadByHash(hash []byte) (*types.Poll, error)

	IsVoted(id uint64, addr sdk.ValAddress) bool
	SetVoted(id uint64, addr sdk.ValAddress)

	Create(poll types.Poll) (uint64, error)
	Save(poll types.Poll) error

	Delete(id uint64) error
	DeleteByHash(hash []byte) error

	Paginate(page *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error)

	// ExportGenesis returns the entire module's state
	ExportGenesis() (genState *types.GenesisPoll_ChainSet, err error)

	// ImportGenesis sets the entire module's state
	ImportGenesis(genState *types.GenesisPoll_ChainSet) error
}

var (
	kvPollRepoKeyLatestId = []byte{0x01}
	kvPollRepoItemsPrefix = []byte{0x02}
	kvPollRepoHashPrefix  = []byte{0x03}
	kvPollRepoVotePrefix  = []byte{0x04}
)

type kvPollRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVPollRepo(cdc codec.BinaryCodec, chain byte, store store.KVStore) PollRepo {
	return kvPollRepo{
		cdc,
		prefix.NewStore(store, append(kvPollRepoKey, chain)),
	}
}

func (k kvPollRepo) Load(id uint64) (*types.Poll, error) {
	bz := prefix.NewStore(k.root, kvPollRepoItemsPrefix).Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return nil, nil
	}

	poll := new(types.Poll)
	if err := poll.Unmarshal(bz); err != nil {
		return nil, err
	}

	return poll, nil
}

func (k kvPollRepo) LoadByHash(hash []byte) (*types.Poll, error) {
	id := prefix.NewStore(k.root, kvPollRepoHashPrefix).Get(hash)
	if id == nil {
		return nil, nil
	}
	return k.Load(sdk.BigEndianToUint64(id))
}

func (k kvPollRepo) IsVoted(id uint64, addr sdk.ValAddress) bool {
	return prefix.NewStore(k.root, append(kvPollRepoVotePrefix, sdk.Uint64ToBigEndian(id)...)).Has(addr.Bytes())
}

func (k kvPollRepo) SetVoted(id uint64, addr sdk.ValAddress) {
	prefix.NewStore(k.root, append(kvPollRepoVotePrefix, sdk.Uint64ToBigEndian(id)...)).Set(addr.Bytes(), []byte{})
}

func (k kvPollRepo) Create(poll types.Poll) (uint64, error) {
	itemStore := prefix.NewStore(k.root, kvPollRepoItemsPrefix)
	hashStore := prefix.NewStore(k.root, kvPollRepoHashPrefix)

	latestId := sdk.BigEndianToUint64(k.root.Get(kvPollRepoKeyLatestId))
	latestIdBz := sdk.Uint64ToBigEndian(latestId)

	poll.Id = latestId
	pollBz, err := poll.Marshal()
	if err != nil {
		return 0, err
	}
	evtHash, err := poll.GetPayload().Hash()
	if err != nil {
		return 0, err
	}

	itemStore.Set(latestIdBz, pollBz)
	hashStore.Set(evtHash, latestIdBz)

	latestId++
	k.root.Set(kvPollRepoKeyLatestId, sdk.Uint64ToBigEndian(latestId))

	return poll.Id, nil
}

func (k kvPollRepo) Save(poll types.Poll) error {
	itemStore := prefix.NewStore(k.root, kvPollRepoItemsPrefix)

	pollBz, err := poll.Marshal()
	if err != nil {
		return err
	}

	itemStore.Set(sdk.Uint64ToBigEndian(poll.Id), pollBz)

	return nil
}

func (k kvPollRepo) Delete(id uint64) error {
	ps := prefix.NewStore(k.root, kvPollRepoItemsPrefix)
	bz := ps.Get(sdk.Uint64ToBigEndian(id))

	var poll types.Poll
	if err := poll.Unmarshal(bz); err != nil {
		return err
	}

	evtHash, err := poll.GetPayload().Hash()
	if err != nil {
		return err
	}

	ps.Delete(sdk.Uint64ToBigEndian(id))
	prefix.NewStore(k.root, kvPollRepoHashPrefix).Delete(evtHash)
	return nil
}

func (k kvPollRepo) DeleteByHash(hash []byte) error {
	id := prefix.NewStore(k.root, kvPollRepoHashPrefix).Get(hash)
	return k.Delete(sdk.BigEndianToUint64(id))
}

func (k kvPollRepo) Paginate(page *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error) {
	ps := prefix.NewStore(k.root, kvPollRepoItemsPrefix)

	var rs []mitotypes.KV[uint64, *types.Poll]
	pageResp, err := query.Paginate(ps, page, func(key []byte, value []byte) error {
		poll := new(types.Poll)
		if err := poll.Unmarshal(value); err != nil {
			return err
		}

		rs = append(rs, mitotypes.NewKV(sdk.BigEndianToUint64(key), poll))
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return rs, pageResp, nil
}

func (k kvPollRepo) exportItemSet(store store.KVStore) (itemSet []*types.GenesisPoll_ItemSet, err error) {
	_, err = query.Paginate(
		store,
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, value []byte) error {
			poll := new(types.Poll)
			if err := poll.Unmarshal(value); err != nil {
				return err
			}

			itemSet = append(itemSet, &types.GenesisPoll_ItemSet{
				Id:   sdk.BigEndianToUint64(key),
				Poll: poll,
			})

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return
}

func (k kvPollRepo) exportHashSet(store store.KVStore) (hashSet []*types.GenesisPoll_HashSet, err error) {
	_, err = query.Paginate(
		store,
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, value []byte) error {
			hashSet = append(hashSet, &types.GenesisPoll_HashSet{
				Hash: key,
				Id:   sdk.BigEndianToUint64(value),
			})
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return
}

func (k kvPollRepo) ExportGenesis() (genState *types.GenesisPoll_ChainSet, err error) {
	// initialize
	genState = &types.GenesisPoll_ChainSet{}

	// load latest id
	genState.LatestId = sdk.BigEndianToUint64(k.root.Get(kvPollRepoKeyLatestId))

	// load item set
	genState.ItemSet, err = k.exportItemSet(prefix.NewStore(k.root, kvPollRepoItemsPrefix))
	if err != nil {
		return nil, err
	}

	// load hash set
	genState.HashSet, err = k.exportHashSet(prefix.NewStore(k.root, kvPollRepoHashPrefix))
	if err != nil {
		return nil, err
	}

	return
}

func (k kvPollRepo) ImportGenesis(genState *types.GenesisPoll_ChainSet) error {
	// save latest id
	k.root.Set(kvPollRepoKeyLatestId, sdk.Uint64ToBigEndian(genState.GetLatestId()))

	// save item set
	pis := prefix.NewStore(k.root, kvPollRepoItemsPrefix)
	for _, item := range genState.GetItemSet() {
		bz, err := item.GetPoll().Marshal()
		if err != nil {
			return err
		}
		pis.Set(sdk.Uint64ToBigEndian(item.GetId()), bz)
	}

	// save hash set
	phs := prefix.NewStore(k.root, kvPollRepoHashPrefix)
	for _, item := range genState.GetHashSet() {
		phs.Set(item.GetHash(), sdk.Uint64ToBigEndian(item.GetId()))
	}

	return nil
}
