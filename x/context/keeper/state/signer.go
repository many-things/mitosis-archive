package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/many-things/mitosis/x/context/types"
)

type SignerRepo interface {
	Load(id uint64) (*types.Signer, error)
	Create(signer *types.Signer) (uint64, error)
	Save(id uint64, signer *types.Signer) error
}

type kvSignerRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVSignerRepo(cdc codec.BinaryCodec, store store.KVStore) SignerRepo {
	return kvSignerRepo{
		cdc,
		prefix.NewStore(store, kvSignerRepoKey),
	}
}

func (k kvSignerRepo) Load(id uint64) (*types.Signer, error) {
	_ = id
	return nil, nil
}

func (k kvSignerRepo) Create(signer *types.Signer) (uint64, error) {
	_ = signer
	return 0, nil
}

func (k kvSignerRepo) Save(id uint64, signer *types.Signer) error {
	_, _ = id, signer
	return nil
}
