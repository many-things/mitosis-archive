package state

import (
	sdkerrutils "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/context/types"
)

type SignerRepo interface {
	Load(chain string) (*types.Signer, error)
	Save(signer *types.Signer) error
}

var (
	kvSignerRepoItemPrefix = []byte{0x02}
)

type kvSignerRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVSignerRepo(cdc codec.BinaryCodec, store store.KVStore) SignerRepo {
	root := prefix.NewStore(store, kvSignerRepoKey)

	return kvSignerRepo{cdc, root}
}

func (k kvSignerRepo) items() store.KVStore {
	return prefix.NewStore(k.root, kvSignerRepoItemPrefix)
}

func (k kvSignerRepo) Load(chain string) (*types.Signer, error) {
	bz := k.items().Get([]byte(chain))
	if bz == nil {
		return nil, sdkerrutils.Wrap(sdkerrors.ErrNotFound, "signer not found")
	}

	signer := new(types.Signer)
	if err := k.cdc.Unmarshal(bz, signer); err != nil {
		return nil, sdkerrutils.Wrap(err, "unmarshal signer")
	}

	return signer, nil
}

func (k kvSignerRepo) Save(signer *types.Signer) error {
	k.items().Set([]byte(signer.Chain), k.cdc.MustMarshal(signer))

	return nil
}
