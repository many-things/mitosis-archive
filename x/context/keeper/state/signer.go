package state

import (
	sdkerrutils "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/context/types"
)

type SignerRepo interface {
	Load(id uint64) (*types.Signer, error)
	LoadByChain(chain string) (*types.Signer, error)
	Create(signer *types.Signer) (uint64, error)
	Save(signer *types.Signer) error
}

var (
	kvSignerRepoLastIDKey   = []byte{0x01}
	kvSignerRepoItemPrefix  = []byte{0x02}
	kvSignerRepoChainPrefix = []byte{0x03}
)

type kvSignerRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVSignerRepo(cdc codec.BinaryCodec, store store.KVStore) SignerRepo {
	root := prefix.NewStore(store, kvSignerRepoKey)

	return kvSignerRepo{cdc, root}
}

func (k kvSignerRepo) chains() store.KVStore {
	return prefix.NewStore(k.root, kvSignerRepoChainPrefix)
}

func (k kvSignerRepo) items() store.KVStore {
	return prefix.NewStore(k.root, kvSignerRepoItemPrefix)
}

func (k kvSignerRepo) lastID() uint64 {
	bz := k.root.Get(kvSignerRepoLastIDKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k kvSignerRepo) Load(id uint64) (*types.Signer, error) {
	bz := k.items().Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return nil, sdkerrutils.Wrap(sdkerrors.ErrNotFound, "signer not found")
	}

	signer := new(types.Signer)
	if err := k.cdc.Unmarshal(bz, signer); err != nil {
		return nil, sdkerrutils.Wrap(err, "unmarshal signer")
	}

	return signer, nil
}

func (k kvSignerRepo) LoadByChain(chain string) (*types.Signer, error) {
	bz := k.chains().Get([]byte(chain))
	if bz == nil {
		return nil, sdkerrutils.Wrap(sdkerrors.ErrNotFound, "signer not found")
	}
	return k.Load(sdk.BigEndianToUint64(bz))
}

func (k kvSignerRepo) Create(signer *types.Signer) (uint64, error) {
	signer.ID = k.lastID()

	signerID := sdk.Uint64ToBigEndian(signer.ID)
	nextID := sdk.Uint64ToBigEndian(signer.ID + 1)

	k.items().Set(signerID, k.cdc.MustMarshal(signer))
	k.chains().Set([]byte(signer.Chain), signerID)
	k.root.Set(kvSignerRepoLastIDKey, nextID)

	return signer.ID, nil
}

func (k kvSignerRepo) Save(signer *types.Signer) error {
	k.items().Set(sdk.Uint64ToBigEndian(signer.ID), k.cdc.MustMarshal(signer))
	return nil
}
