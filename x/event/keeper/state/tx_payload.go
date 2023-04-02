package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/many-things/mitosis/x/event/types"
)

type TxPayloadRepo interface {
	Load(pollID uint64) (types.MsgHash, error)
	Create(pollID uint64, hash types.MsgHash) error
	Save(pollID uint64, hash types.MsgHash) error
	Delete(pollID uint64) error

	ExportGenesis() (*types.GenesisProxy, error)
	ImportGenesis(genState *types.GenesisProxy) error
}

var (
	kvTxPayloadItemPrefix = []byte{0x01}
)

type kvTxPayloadRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVTxPayloadRepo(cdc codec.BinaryCodec, store store.KVStore, chain string) TxPayloadRepo {
	return kvTxPayloadRepo{
		cdc:  cdc,
		root: prefix.NewStore(store, append(kvProxyRepoKey, []byte(chain)...)),
	}
}

func (k kvTxPayloadRepo) Load(pollID uint64) (types.MsgHash, error) {
	return nil, nil
}

func (k kvTxPayloadRepo) Create(pollID uint64, hash types.MsgHash) error {
	return nil
}

func (k kvTxPayloadRepo) Save(pollID uint64, hash types.MsgHash) error {
	return nil
}

func (k kvTxPayloadRepo) Delete(pollID uint64) error {
	return nil
}

func (k kvTxPayloadRepo) ExportGenesis() (*types.GenesisProxy, error) {
	return nil, nil
}

func (k kvTxPayloadRepo) ImportGenesis(genState *types.GenesisPayload) {

}