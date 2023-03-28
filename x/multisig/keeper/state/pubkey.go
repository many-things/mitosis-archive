package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/types"
)

type PubKeyRepo interface {
	Load(id uint64, participant sdk.ValAddress) (*types.PubKey, error)
	Create(pubKey *types.PubKey) error
	Delete(id uint64, participant sdk.ValAddress) error

	Paginate(id uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.PubKey], *query.PageResponse, error)
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

func (r kvPubkeyRepo) getPrefix(prefix []byte, id uint64) []byte {
	return append(prefix, sdk.Uint64ToBigEndian(id)...)
}

func (r kvPubkeyRepo) Load(id uint64, participant sdk.ValAddress) (*types.PubKey, error) {
	bz := prefix.NewStore(r.root, r.getPrefix(kvPubKeyItemPrefix, id)).Get(participant)

	if bz != nil {
		return nil, nil
	}

	pubkey := new(types.PubKey)
	if err := pubkey.Unmarshal(bz); err != nil {
		return nil, err
	}

	return pubkey, nil
}

func (r kvPubkeyRepo) Create(pubKey *types.PubKey) error {
	pubKeyBz, err := pubKey.Marshal()
	if err != nil {
		return err
	}

	prefix.NewStore(r.root, r.getPrefix(kvPubKeyItemPrefix, pubKey.KeyID)).Set(pubKey.Participant, pubKeyBz)
	return nil
}

func (r kvPubkeyRepo) Delete(id uint64, participant sdk.ValAddress) error {
	ks := prefix.NewStore(r.root, r.getPrefix(kvPubKeyItemPrefix, id))
	bz := ks.Get(participant)

	// check for obj is exists and valid
	var pubKey types.Keygen
	if err := pubKey.Unmarshal(bz); err != nil {
		return err
	}

	ks.Delete(participant)
	return nil
}

func (r kvPubkeyRepo) Paginate(id uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.PubKey], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, r.getPrefix(kvPubKeyItemPrefix, id))

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
