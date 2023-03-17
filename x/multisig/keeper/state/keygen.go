package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
)

type KeygenRepo interface {
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
