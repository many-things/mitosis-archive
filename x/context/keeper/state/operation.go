package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/queue"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/pkg/errors"
)

var _ queue.Message = (*ID)(nil)

type ID uint64

func (i *ID) Marshal() ([]byte, error) {
	return sdk.Uint64ToBigEndian(uint64(*i)), nil
}

func (i *ID) Unmarshal(bz []byte) error {
	*i = ID(sdk.BigEndianToUint64(bz))
	return nil
}

type OperationRepo interface {
	// immut

	Load(id uint64) (*types.Operation, error)
	MustLoad(id uint64) *types.Operation
	Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Operation], *query.PageResponse, error)
	PaginateStatus(status types.Operation_Status, pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Operation], *query.PageResponse, error)

	// mut

	Create(op *types.Operation) (uint64, error)
	Save(op *types.Operation) error
	Shift(id uint64, status types.Operation_Status) error

	// genesis

	ExportGenesis() (genState *types.GenesisOperation, err error)
	ImportGenesis(genState *types.GenesisOperation) error
}

var (
	kvOperationRepoLastIDKey           = []byte{0x01}
	kvOperationRepoItemsPrefix         = []byte{0x02}
	kvOperationRepoIndexByStatusPrefix = []byte{0x03}
)

type kvOperationRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVOperationRepo(cdc codec.BinaryCodec, store store.KVStore) OperationRepo {
	root := prefix.NewStore(store, kvOperationRepoKey)

	return &kvOperationRepo{cdc, root}
}

func (k kvOperationRepo) items() store.KVStore {
	return prefix.NewStore(k.root, kvOperationRepoItemsPrefix)
}

func (k kvOperationRepo) indexes(status types.Operation_Status) store.KVStore {
	return prefix.NewStore(
		k.root,
		append(
			kvOperationRepoIndexByStatusPrefix,
			sdk.Uint64ToBigEndian(uint64(status))...,
		),
	)
}

func (k kvOperationRepo) lastID() uint64 {
	bz := k.root.Get(kvOperationRepoLastIDKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k kvOperationRepo) setLastID(id uint64) {
	k.root.Set(kvOperationRepoLastIDKey, sdk.Uint64ToBigEndian(id))
}

func (k kvOperationRepo) Load(id uint64) (*types.Operation, error) {
	bz := k.items().Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return nil, errors.New("operation not found")
	}

	operation := new(types.Operation)
	if err := k.cdc.Unmarshal(bz, operation); err != nil {
		return nil, errors.Wrap(err, "unmarshal operation")
	}

	return operation, nil
}

func (k kvOperationRepo) MustLoad(id uint64) *types.Operation {
	operation, err := k.Load(id)
	if err != nil {
		panic(err)
	}

	return operation
}

func (k kvOperationRepo) Create(op *types.Operation) (uint64, error) {
	op.ID = k.lastID()

	opID := sdk.Uint64ToBigEndian(op.ID)
	nextID := op.ID + 1

	k.items().Set(opID, k.cdc.MustMarshal(op))
	k.indexes(types.Operation_StatusPending).Set(opID, []byte{})
	k.setLastID(nextID)

	return op.ID, nil
}

func (k kvOperationRepo) Save(op *types.Operation) error {
	k.items().Set(sdk.Uint64ToBigEndian(op.ID), k.cdc.MustMarshal(op))

	return nil
}

func (k kvOperationRepo) Shift(id uint64, status types.Operation_Status) error {
	idBz := sdk.Uint64ToBigEndian(id)
	op := k.MustLoad(id)
	k.indexes(op.Status).Delete(idBz)
	k.indexes(status).Set(idBz, []byte{})

	return nil
}

func (k kvOperationRepo) Paginate(pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Operation], *query.PageResponse, error) {
	var kvs []mitotypes.KV[uint64, *types.Operation]

	pageResp, err := query.Paginate(
		k.items(),
		pageReq,
		func(id []byte, bz []byte) error {
			op := new(types.Operation)
			if err := k.cdc.Unmarshal(bz, op); err != nil {
				return err
			}

			kvs = append(
				kvs,
				mitotypes.NewKV(sdk.BigEndianToUint64(id), op),
			)

			return nil
		},
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "paginate")
	}

	return kvs, pageResp, nil
}

func (k kvOperationRepo) PaginateStatus(status types.Operation_Status, pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Operation], *query.PageResponse, error) {
	var kvs []mitotypes.KV[uint64, *types.Operation]

	pageResp, err := query.Paginate(
		k.indexes(status),
		pageReq,
		func(idBz []byte, _ []byte) error {
			id := sdk.BigEndianToUint64(idBz)
			op := k.MustLoad(id)

			kvs = append(kvs, mitotypes.NewKV(id, op))
			return nil
		},
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "paginate")
	}

	return kvs, pageResp, nil
}

func (k kvOperationRepo) ExportGenesis() (genState *types.GenesisOperation, err error) {
	genState = &types.GenesisOperation{}

	genState.LastId = k.lastID()

	_, err = query.Paginate(
		k.items(),
		&query.PageRequest{Limit: query.MaxLimit},
		func(id []byte, bz []byte) error {
			op := new(types.Operation)
			if err := k.cdc.Unmarshal(bz, op); err != nil {
				return errors.Wrap(err, "unmarshal operation")
			}

			genState.ItemSet = append(
				genState.ItemSet,
				&types.GenesisOperation_ItemSet{
					ID:        sdk.BigEndianToUint64(id),
					Operation: op,
				},
			)
			return nil
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "paginate items")
	}

	for s := range types.Operation_Status_name {
		items, _, err := k.PaginateStatus(
			types.Operation_Status(s),
			&query.PageRequest{Limit: query.MaxLimit},
		)
		if err != nil {
			return nil, errors.Wrap(err, "paginate status")
		}

		for _, item := range items {
			genState.IndexSet = append(
				genState.IndexSet,
				&types.GenesisOperation_IndexSet{
					Status: types.Operation_Status(s),
					OpID:   item.Key,
				},
			)
		}
	}

	return
}

func (k kvOperationRepo) ImportGenesis(genState *types.GenesisOperation) error {
	k.setLastID(genState.LastId)

	itemStore := k.items()
	for _, item := range genState.ItemSet {
		itemStore.Set(sdk.Uint64ToBigEndian(item.ID), k.cdc.MustMarshal(item.Operation))
	}

	for _, item := range genState.IndexSet {
		k.indexes(item.Status).Set(sdk.Uint64ToBigEndian(item.OpID), []byte{})
	}

	return nil
}
