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
	"golang.org/x/exp/slices"
)

type SignResultRepo interface {
	Load(id uint64) (*exported.SignResult, error)
	Save(signature *exported.SignResult) error
	Delete(id uint64) error
	AddParticipantSignResult(id uint64, participant sdk.ValAddress, signature exported.Signature) error
	RemoveParticipantSignResult(id uint64, participant sdk.ValAddress) error
	HasSignResult(id uint64) bool

	Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.SignResult], *query.PageResponse, error)

	ExportGenesis() ([]mitosistype.KV[uint64, *exported.SignResult], error)
	ImportGenesis(genState []mitosistype.KV[uint64, *exported.SignResult]) error
}

type kvSignResultRepo struct {
	cdc     codec.BinaryCodec
	root    store.KVStore
	chainID string
}

var (
	kvSignResultRepoItemPrefix = []byte{0x01}
)

func NewKVChainSignResultRepo(cdc codec.BinaryCodec, root store.KVStore, chainID string) SignResultRepo {
	return &kvSignResultRepo{
		cdc:     cdc,
		root:    prefix.NewStore(root, append([]byte(chainID), kvSignResultRepoKey...)),
		chainID: chainID,
	}
}

func (r kvSignResultRepo) Load(id uint64) (*exported.SignResult, error) {
	bz := prefix.NewStore(r.root, kvSignResultRepoItemPrefix).Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	signSignResult := new(exported.SignResult)
	if err := signSignResult.Unmarshal(bz); err != nil {
		return nil, err
	}

	return signSignResult, nil
}

func (r kvSignResultRepo) Save(signature *exported.SignResult) error {
	bz, err := signature.Marshal()
	if err != nil {
		return err
	}

	prefix.NewStore(r.root, kvSignResultRepoItemPrefix).Set(sdk.Uint64ToBigEndian(signature.SigID), bz)
	return nil
}

func (r kvSignResultRepo) HasSignResult(id uint64) bool {
	return prefix.NewStore(r.root, kvSignResultRepoItemPrefix).Has(sdk.Uint64ToBigEndian(id))
}

func (r kvSignResultRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, kvSignResultRepoItemPrefix)
	idBz := sdk.Uint64ToBigEndian(id)

	if !ks.Has(idBz) {
		return sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	ks.Delete(idBz)
	return nil
}

func (r kvSignResultRepo) AddParticipantSignResult(id uint64, participant sdk.ValAddress, signature exported.Signature) error {
	ks := prefix.NewStore(r.root, kvSignResultRepoItemPrefix)
	idBz := sdk.Uint64ToBigEndian(id)
	bz := ks.Get(idBz)

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	signSignResult := new(exported.SignResult)
	if err := signSignResult.Unmarshal(bz); err != nil {
		return err
	}

	exists := false
	for _, item := range signSignResult.Items {
		if item.Participant.Equals(participant) {
			exists = true
			item.Signature = signature
			break
		}
	}
	if !exists {
		signSignResult.Items = append(signSignResult.Items, &exported.SignResult_Item{
			Participant: participant,
			Signature:   signature,
		})
	}

	sigBz, err := signSignResult.Marshal()
	if err != nil {
		return err
	}
	ks.Set(idBz, sigBz)
	return nil
}

func (r kvSignResultRepo) RemoveParticipantSignResult(id uint64, participant sdk.ValAddress) error {
	ks := prefix.NewStore(r.root, kvSignResultRepoItemPrefix)
	idBz := sdk.Uint64ToBigEndian(id)
	bz := ks.Get(idBz)

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	signSignResult := new(exported.SignResult)
	if err := signSignResult.Unmarshal(bz); err != nil {
		return err
	}

	targetID := -1
	for i, item := range signSignResult.Items {
		if item.Participant.Equals(participant) {
			targetID = i
			break
		}
	}

	if targetID < 0 {
		return sdkerrors.Wrap(errors.ErrNotFound, "signature")
	}
	signSignResult.Items = slices.Delete(signSignResult.Items, targetID, targetID+1)

	sigBz, err := signSignResult.Marshal()
	if err != nil {
		return err
	}
	ks.Set(idBz, sigBz)
	return nil
}

func (r kvSignResultRepo) Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.SignResult], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, kvSignResultRepoItemPrefix)

	var results []mitosistype.KV[uint64, *exported.SignResult]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		signSignResult := new(exported.SignResult)
		if err := signSignResult.Unmarshal(value); err != nil {
			return err
		}

		results = append(results, mitosistype.NewKV(sdk.BigEndianToUint64(key), signSignResult))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}

func (r kvSignResultRepo) ExportGenesis() (set []mitosistype.KV[uint64, *exported.SignResult], err error) {
	ks := prefix.NewStore(r.root, kvSignResultRepoItemPrefix)

	_, err = query.Paginate(
		ks,
		&query.PageRequest{Limit: query.MaxLimit},
		func(_ []byte, value []byte) error {
			item := new(exported.SignResult)

			if err := item.Unmarshal(value); err != nil {
				return err
			}

			set = append(set, mitosistype.NewKV(item.SigID, item))
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return
}

func (r kvSignResultRepo) ImportGenesis(genState []mitosistype.KV[uint64, *exported.SignResult]) error {
	ks := prefix.NewStore(r.root, kvSignResultRepoItemPrefix)

	for _, item := range genState {
		bz, err := r.cdc.Marshal(item.Value)
		if err != nil {
			return err
		}

		ks.Set(sdk.Uint64ToBigEndian(item.Value.SigID), bz)
	}

	return nil
}
