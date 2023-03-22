package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/queue"
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

	Flush(threshold sdk.Dec) ([]mitotypes.KV[uint64, *types.Poll], error)

	Paginate(page *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error)

	// ExportGenesis returns the entire module's state
	ExportGenesis() (genState *types.GenesisPoll_ChainSet, err error)

	// ImportGenesis sets the entire module's state
	ImportGenesis(genState *types.GenesisPoll_ChainSet) error
}

var (
	kvPollRepoQueuePrefix = []byte{0x01}
	kvPollRepoHashPrefix  = []byte{0x02}
	kvPollRepoVotePrefix  = []byte{0x03}
)

type kvPollRepo struct {
	cdc   codec.BinaryCodec
	root  store.KVStore
	queue queue.Queue[*types.Poll]
}

func NewKVPollRepo(cdc codec.BinaryCodec, chain byte, store store.KVStore) PollRepo {
	root := prefix.NewStore(store, append(kvPollRepoKey, chain))

	return kvPollRepo{
		cdc,
		root,
		queue.NewKVQueue(
			prefix.NewStore(root, kvPollRepoQueuePrefix),
			func() *types.Poll { return &types.Poll{} },
		),
	}
}

func (k kvPollRepo) Load(id uint64) (*types.Poll, error) {
	return k.queue.Get(id)
}

func (k kvPollRepo) LoadByHash(hash []byte) (*types.Poll, error) {
	id := prefix.NewStore(k.root, kvPollRepoHashPrefix).Get(hash)
	if id == nil {
		return nil, nil
	}
	return k.queue.Get(sdk.BigEndianToUint64(id))
}

func (k kvPollRepo) IsVoted(id uint64, addr sdk.ValAddress) bool {
	return prefix.NewStore(k.root, append(kvPollRepoVotePrefix, sdk.Uint64ToBigEndian(id)...)).Has(addr.Bytes())
}

func (k kvPollRepo) SetVoted(id uint64, addr sdk.ValAddress) {
	prefix.NewStore(k.root, append(kvPollRepoVotePrefix, sdk.Uint64ToBigEndian(id)...)).Set(addr.Bytes(), []byte{})
}

func (k kvPollRepo) Create(poll types.Poll) (uint64, error) {
	hashStore := prefix.NewStore(k.root, kvPollRepoHashPrefix)

	poll.Id = k.queue.LastIndex()

	ids, err := k.queue.Produce(&poll)
	if err != nil {
		return 0, err
	}
	if len(ids) != 1 {
		return 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "queue.Produce returned more than one id")
	}

	evtHash, err := poll.GetPayload().Hash()
	if err != nil {
		return 0, err
	}

	id := ids[0]

	hashStore.Set(evtHash, sdk.Uint64ToBigEndian(id))

	return id, nil
}

func (k kvPollRepo) Save(poll types.Poll) error {
	return k.queue.Update(poll.GetId(), &poll)
}

func (k kvPollRepo) Flush(threshold sdk.Dec) ([]mitotypes.KV[uint64, *types.Poll], error) {
	checker := func(poll *types.Poll, u uint64) (bool, error) {
		return poll.GetTally().Passed(threshold), nil
	}
	passed, err := k.queue.ConsumeUntil(checker)
	if err != nil {
		return nil, err
	}
	return passed, nil
}

func (k kvPollRepo) Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error) {
	var kvs []mitotypes.KV[uint64, *types.Poll]

	pageResp, err := k.queue.Paginate(
		pageReq,
		func(poll *types.Poll, u uint64) error {
			kvs = append(kvs, mitotypes.NewKV(u, poll))
			return nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return kvs, pageResp, nil
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

	// load item set
	queueGenesis, err := k.queue.ExportGenesis()
	if err != nil {
		return nil, err
	}

	genState.FirstId = queueGenesis.FirstIndex
	genState.LastId = queueGenesis.LastIndex
	genState.ItemSet = mitotypes.MapKV(
		queueGenesis.Items,
		func(k uint64, v *types.Poll, _ int) *types.GenesisPoll_ItemSet {
			return &types.GenesisPoll_ItemSet{
				Id:   k,
				Poll: v,
			}
		},
	)

	// load hash set
	genState.HashSet, err = k.exportHashSet(prefix.NewStore(k.root, kvPollRepoHashPrefix))
	if err != nil {
		return nil, err
	}

	return
}

func (k kvPollRepo) ImportGenesis(genState *types.GenesisPoll_ChainSet) error {
	// save item set
	queueGenesis := queue.GenesisState[*types.Poll]{
		LastIndex:  genState.LastId,
		FirstIndex: genState.FirstId,
		Items: mitotypes.Map(
			genState.GetItemSet(),
			func(t *types.GenesisPoll_ItemSet, _ int) mitotypes.KV[uint64, *types.Poll] {
				return mitotypes.NewKV(t.GetId(), t.GetPoll())
			},
		),
	}
	if err := k.queue.ImportGenesis(queueGenesis); err != nil {
		return err
	}

	// save hash set
	phs := prefix.NewStore(k.root, kvPollRepoHashPrefix)
	for _, item := range genState.GetHashSet() {
		phs.Set(item.GetHash(), sdk.Uint64ToBigEndian(item.GetId()))
	}

	return nil
}
