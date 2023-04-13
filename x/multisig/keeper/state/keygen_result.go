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

type KeygenResultRepo interface {
	Load(id uint64) (*types.KeygenResult, error)
	Create(pubKey *types.KeygenResult) error
	Delete(id uint64) error
	AddKey(id uint64, participant sdk.ValAddress, publicKey exported.PublicKey) error
	RemoveKey(id uint64, participant sdk.ValAddress) error
	HasKeygenResult(id uint64) bool

	Paginate(page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.KeygenResult], *query.PageResponse, error)

	ExportGenesis() ([]mitosistype.KV[uint64, *types.KeygenResult], error)
	ImportGenesis(genState []mitosistype.KV[uint64, *types.KeygenResult]) error
}

type kvKeygenResultRepo struct {
	cdc     codec.BinaryCodec
	root    store.KVStore
	chainID string
}

func NewKVChainKeygenResultRepo(cdc codec.BinaryCodec, root store.KVStore, chainID string) KeygenResultRepo {
	return &kvKeygenResultRepo{
		cdc, prefix.NewStore(root, append([]byte(chainID), kvKeygenResultRepoKey...)), chainID,
	}
}

var (
	kvKeygenResultItemPrefix = []byte{0x01}
)

func (r kvKeygenResultRepo) Load(id uint64) (*types.KeygenResult, error) {
	bz := prefix.NewStore(r.root, kvKeygenResultItemPrefix).Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "keygen")
	}

	pubkey := new(types.KeygenResult)
	if err := pubkey.Unmarshal(bz); err != nil {
		return nil, err
	}

	return pubkey, nil
}

func (r kvKeygenResultRepo) HasKeygenResult(id uint64) bool {
	return prefix.NewStore(r.root, kvKeygenResultItemPrefix).Has(sdk.Uint64ToBigEndian(id))
}

func (r kvKeygenResultRepo) Create(pubKey *types.KeygenResult) error {
	pubKeyBz, err := pubKey.Marshal()
	if err != nil {
		return err
	}

	prefix.NewStore(r.root, kvKeygenResultItemPrefix).Set(sdk.Uint64ToBigEndian(pubKey.KeyID), pubKeyBz)
	return nil
}

func (r kvKeygenResultRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, kvKeygenResultItemPrefix)

	if !ks.Has(sdk.Uint64ToBigEndian(id)) {
		return sdkerrors.Wrap(errors.ErrNotFound, "pubkey")
	}

	ks.Delete(sdk.Uint64ToBigEndian(id))
	return nil
}

func (r kvKeygenResultRepo) AddKey(id uint64, participant sdk.ValAddress, publicKey exported.PublicKey) error {
	ks := prefix.NewStore(r.root, kvKeygenResultItemPrefix)
	bz := ks.Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "pubkey")
	}

	pubKey := new(types.KeygenResult)
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
		pubKey.Items = append(pubKey.Items, &types.KeygenResult_Item{
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

func (r kvKeygenResultRepo) RemoveKey(id uint64, participant sdk.ValAddress) error {
	ks := prefix.NewStore(r.root, kvKeygenResultItemPrefix)
	bz := ks.Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "pubkey")
	}

	pubKey := new(types.KeygenResult)
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

func (r kvKeygenResultRepo) Paginate(page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.KeygenResult], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, kvKeygenResultItemPrefix)

	var results []mitosistype.KV[sdk.ValAddress, *types.KeygenResult]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		pubKey := new(types.KeygenResult)

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

func (r kvKeygenResultRepo) ExportGenesis() (set []mitosistype.KV[uint64, *types.KeygenResult], err error) {
	ks := prefix.NewStore(r.root, kvKeygenResultItemPrefix)

	_, err = query.Paginate(
		ks,
		&query.PageRequest{Limit: query.MaxLimit},
		func(_ []byte, bz []byte) error {
			item := new(types.KeygenResult)
			if err := item.Unmarshal(bz); err != nil {
				return err
			}

			set = append(set, mitosistype.NewKV(item.KeyID, item))
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return
}

func (r kvKeygenResultRepo) ImportGenesis(genState []mitosistype.KV[uint64, *types.KeygenResult]) error {
	ks := prefix.NewStore(r.root, kvKeygenResultItemPrefix)

	for _, item := range genState {
		bz, err := r.cdc.Marshal(item.Value)
		if err != nil {
			return err
		}

		ks.Set(sdk.Uint64ToBigEndian(item.Value.KeyID), bz)
	}

	return nil
}
