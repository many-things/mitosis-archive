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
	Load(id uint64) (types.Poll, error)
	LoadByHash(hash []byte) (types.Poll, error)

	Save(poll types.Poll) error

	Delete(id uint64) error
	DeleteByHash(hash []byte) error

	Paginate(page *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error)
}

var (
	kvPollRepoKeyLatestId = []byte{0x01}
	kvPollRepoItemsPrefix = []byte{0x02}
	kvPollRepoHashPrefix  = []byte{0x03}
)

type kvPollRepo struct {
	cdc  codec.Codec
	root store.KVStore
}

func NewKVPollRepo(cdc codec.Codec, store store.KVStore) PollRepo {
	return kvPollRepo{cdc, store}
}

func (k kvPollRepo) Load(id uint64) (types.Poll, error) {
	bz := prefix.NewStore(k.root, kvPollRepoItemsPrefix).Get(sdk.Uint64ToBigEndian(id))

	var poll types.Poll
	if err := poll.Unmarshal(bz); err != nil {
		return types.Poll{}, err
	}

	return poll, nil
}

func (k kvPollRepo) LoadByHash(hash []byte) (types.Poll, error) {
	id := prefix.NewStore(k.root, kvPollRepoHashPrefix).Get(hash)
	return k.Load(sdk.BigEndianToUint64(id))
}

func (k kvPollRepo) Save(poll types.Poll) error {
	latestId := sdk.BigEndianToUint64(k.root.Get(kvPollRepoKeyLatestId))
	latestIdBz := sdk.Uint64ToBigEndian(latestId)

	poll.Id = latestId
	pollBz, err := poll.Marshal()
	if err != nil {
		return err
	}
	evtHash, err := poll.GetPayload().Hash()
	if err != nil {
		return err
	}

	prefix.NewStore(k.root, kvPollRepoItemsPrefix).Set(latestIdBz, pollBz)
	prefix.NewStore(k.root, kvPollRepoHashPrefix).Set(evtHash, latestIdBz)

	latestId++
	k.root.Set(kvPollRepoKeyLatestId, sdk.Uint64ToBigEndian(latestId))

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
