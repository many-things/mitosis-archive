package state

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/queue"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/types"
)

type OperationRepo interface {
	Load(id uint64) (*types.Operation, error)
	Create(op *types.Operation) (uint64, error)
	Save(op *types.Operation) error
	Take(amount uint64) ([]mitotypes.KV[uint64, *types.Operation], error)
	Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Operation], *query.PageResponse, error)

	ExportGenesis() (genState *types.GenesisOperation, err error)
	ImportGenesis(genState *types.GenesisOperation) error
}

var (
	kvOperationRepoQueuePrefix = []byte{0x01}
)

type kvOperationRepo struct {
	cdc   codec.BinaryCodec
	root  store.KVStore
	queue queue.Queue[*types.Operation]
}

func NewKVOperationRepo(cdc codec.BinaryCodec, chain byte, store store.KVStore) OperationRepo {
	root := prefix.NewStore(store, append(kvOperationRepoKey, chain))

	return &kvOperationRepo{
		cdc,
		root,
		queue.NewKVQueue(
			prefix.NewStore(root, kvOperationRepoQueuePrefix),
			func() *types.Operation { return &types.Operation{} },
		),
	}
}

func (k kvOperationRepo) Load(id uint64) (*types.Operation, error) {
	return k.queue.Get(id)
}

func (k kvOperationRepo) Create(op *types.Operation) (uint64, error) {
	op.Id = k.queue.LastIndex()

	ids, err := k.queue.Produce(op)
	if err != nil {
		return 0, err
	}
	if len(ids) != 0 {
		return 0, sdkerrors.Wrap(errors.ErrPanic, "queue.Produce returned more than one id")
	}

	id := ids[0]

	return id, nil
}

func (k kvOperationRepo) Save(op *types.Operation) error {
	return k.queue.Update(op.Id, op)
}

func (k kvOperationRepo) Take(amount uint64) ([]mitotypes.KV[uint64, *types.Operation], error) {
	return k.queue.Consume(amount)
}

func (k kvOperationRepo) Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Operation], *query.PageResponse, error) {
	var kvs []mitotypes.KV[uint64, *types.Operation]

	pageResp, err := k.queue.Paginate(
		pageReq,
		func(op *types.Operation, id uint64) error {
			kvs = append(kvs, mitotypes.NewKV(id, op))
			return nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return kvs, pageResp, nil
}

func (k kvOperationRepo) ExportGenesis() (genState *types.GenesisOperation, err error) {
	genState = &types.GenesisOperation{}

	queueGenesis, err := k.queue.ExportGenesis()
	if err != nil {
		return nil, err
	}

	genState.FirstId = queueGenesis.FirstIndex
	genState.LastId = queueGenesis.LastIndex
	genState.ItemSet = mitotypes.MapKV(
		queueGenesis.Items,
		func(k uint64, v *types.Operation, _ int) *types.GenesisOperation_ItemSet {
			return &types.GenesisOperation_ItemSet{
				Id:        k,
				Operation: v,
			}
		},
	)

	return
}

func (k kvOperationRepo) ImportGenesis(genState *types.GenesisOperation) error {
	queueGenesis := queue.GenesisState[*types.Operation]{
		LastIndex:  genState.LastId,
		FirstIndex: genState.FirstId,
		Items: mitotypes.Map(genState.GetItemSet(), func(t *types.GenesisOperation_ItemSet, _ int) mitotypes.KV[uint64, *types.Operation] {
			return mitotypes.NewKV(t.Id, t.Operation)
		}),
	}
	if err := k.queue.ImportGenesis(queueGenesis); err != nil {
		return err
	}

	return nil
}
