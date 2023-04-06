package state

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/types"
	"golang.org/x/exp/slices"
)

type PubKeyRepo interface {
	Load(id uint64) (*types.PubKey, error)
	Create(pubKey *types.PubKey) error
	Delete(id uint64) error
	AddKey(id uint64, participant sdk.ValAddress, publicKey exported.PublicKey) error
	RemoveKey(id uint64, participant sdk.ValAddress) error
	HasPubKey(id uint64) bool

	Paginate(page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.PubKey], *query.PageResponse, error)
}

type kvPubkeyRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVChainPubKeyRepo(cdc codec.BinaryCodec, root store.KVStore, chainID string) PubKeyRepo {
	return &kvPubkeyRepo{
		cdc, prefix.NewStore(root, append([]byte(chainID), kvPubKeyRepoKey...)),
	}
}

var (
	kvPubKeyItemPrefix = []byte{0x01}
)

func (r kvPubkeyRepo) Load(id uint64) (*types.PubKey, error) {
	bz := prefix.NewStore(r.root, kvPubKeyItemPrefix).Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "pubkey")
	}

	pubkey := new(types.PubKey)
	if err := pubkey.Unmarshal(bz); err != nil {
		return nil, err
	}

	return pubkey, nil
}

func (r kvPubkeyRepo) HasPubKey(id uint64) bool {
	return prefix.NewStore(r.root, kvPubKeyItemPrefix).Has(sdk.Uint64ToBigEndian(id))
}

func (r kvPubkeyRepo) Create(pubKey *types.PubKey) error {
	pubKeyBz, err := pubKey.Marshal()
	if err != nil {
		return err
	}

	prefix.NewStore(r.root, kvPubKeyItemPrefix).Set(sdk.Uint64ToBigEndian(pubKey.KeyID), pubKeyBz)
	return nil
}

func (r kvPubkeyRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, kvPubKeyItemPrefix)

	if !ks.Has(sdk.Uint64ToBigEndian(id)) {
		return sdkerrors.Wrap(errors.ErrNotFound, "pubkey")
	}

	ks.Delete(sdk.Uint64ToBigEndian(id))
	return nil
}

func (r kvPubkeyRepo) AddKey(id uint64, participant sdk.ValAddress, publicKey exported.PublicKey) error {
	ks := prefix.NewStore(r.root, kvPubKeyItemPrefix)
	bz := ks.Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "pubkey")
	}

	pubKey := new(types.PubKey)
	if err := pubKey.Unmarshal(bz); err != nil {
		return err
	}

	exist := false
	for _, item := range pubKey.Items {
		if item.Participant.Equals(participant) {
			exist = true
			item.PubKey = publicKey
			break
		}
	}
	if !exist {
		pubKey.Items = append(pubKey.Items, &types.PubKey_Item{
			Participant: participant,
			PubKey:      publicKey,
		})
	}

	newBz, err := pubKey.Marshal()
	if err != nil {
		return err
	}

	ks.Set(sdk.Uint64ToBigEndian(id), newBz)
	return nil
}

func (r kvPubkeyRepo) RemoveKey(id uint64, participant sdk.ValAddress) error {
	ks := prefix.NewStore(r.root, kvPubKeyItemPrefix)
	bz := ks.Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "pubkey")
	}

	pubKey := new(types.PubKey)
	if err := pubKey.Unmarshal(bz); err != nil {
		return err
	}

	targetID := -1
	for i, item := range pubKey.Items {
		if item.Participant.Equals(participant) {
			targetID = i
			break
		}
	}
	if targetID < 0 {
		return sdkerrors.Wrap(errors.ErrNotFound, "pubkey item")
	}
	pubKey.Items = slices.Delete(pubKey.Items, targetID, targetID+1)

	newBz, err := pubKey.Marshal()
	if err != nil {
		return err
	}

	ks.Set(sdk.Uint64ToBigEndian(id), newBz)
	return nil
}

func (r kvPubkeyRepo) Paginate(page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.PubKey], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, kvPubKeyItemPrefix)

	var results []mitosistype.KV[sdk.ValAddress, *types.PubKey]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		pubKey := new(types.PubKey)

		if err := pubKey.Unmarshal(value); err != nil {
			return err
		}

		addr := sdk.ValAddress(key)
		results = append(results, mitosistype.NewKV(addr, pubKey))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
