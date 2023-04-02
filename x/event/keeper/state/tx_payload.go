package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/x/event/types"
)

type TxPayloadRepo interface {
	Load(pollID uint64) (types.MsgHash, bool)
	Has(pollID uint64) bool
	Save(pollID uint64, hash types.MsgHash) error
	Delete(pollID uint64) error

	ExportGenesis() (*types.GenesisTxPayload_ChainSet, error)
	ImportGenesis(genState *types.GenesisTxPayload_ChainSet) error
}

var (
	kvTxPayloadItemPrefix = []byte{0x01}
)

type kvTxPayloadRepo struct {
	cdc   codec.BinaryCodec
	root  store.KVStore
	chain string
}

func NewKVTxPayloadRepo(cdc codec.BinaryCodec, store store.KVStore, chain string) TxPayloadRepo {
	return kvTxPayloadRepo{
		cdc:   cdc,
		root:  prefix.NewStore(store, append(kvProxyRepoKey, []byte(chain)...)),
		chain: chain,
	}
}

func (k kvTxPayloadRepo) Load(pollID uint64) (types.MsgHash, bool) {
	h := prefix.NewStore(k.root, kvTxPayloadItemPrefix).Get(sdk.Uint64ToBigEndian(pollID))
	return h, h != nil
}

func (k kvTxPayloadRepo) Has(pollID uint64) bool {
	return prefix.NewStore(k.root, kvTxPayloadItemPrefix).Has(sdk.Uint64ToBigEndian(pollID))
}

func (k kvTxPayloadRepo) Save(pollID uint64, hash types.MsgHash) error {
	prefix.NewStore(k.root, kvTxPayloadItemPrefix).Set(sdk.Uint64ToBigEndian(pollID), hash)
	return nil
}

func (k kvTxPayloadRepo) Delete(pollID uint64) error {
	prefix.NewStore(k.root, kvTxPayloadItemPrefix).Delete(sdk.Uint64ToBigEndian(pollID))
	return nil
}

func (k kvTxPayloadRepo) ExportGenesis() (*types.GenesisTxPayload_ChainSet, error) {
	var items []*types.GenesisTxPayload_ItemSet

	_, err := query.Paginate(
		prefix.NewStore(k.root, kvTxPayloadItemPrefix),
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, value []byte) error {
			items = append(
				items,
				&types.GenesisTxPayload_ItemSet{
					PollId:    sdk.BigEndianToUint64(key),
					TxPayload: value,
				})
			return nil
		},
	)

	if err != nil {
		return nil, err
	} else if len(items) == 0 {
		return nil, nil
	}

	return &types.GenesisTxPayload_ChainSet{
		Chain:   []byte(k.chain),
		ItemSet: items,
	}, nil
}

func (k kvTxPayloadRepo) ImportGenesis(genState *types.GenesisTxPayload_ChainSet) error {
	s := prefix.NewStore(k.root, kvTxPayloadItemPrefix)

	for _, item := range genState.GetItemSet() {
		s.Set(sdk.Uint64ToBigEndian(item.PollId), item.TxPayload)
	}

	return nil
}
