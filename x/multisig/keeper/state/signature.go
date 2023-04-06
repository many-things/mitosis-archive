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

type SignatureRepo interface {
	Load(id uint64) (*exported.SignSignature, error)
	Save(signature *exported.SignSignature) error
	Delete(id uint64) error
	AddParticipantSignature(id uint64, participant sdk.ValAddress, signature exported.Signature) error
	RemoveParticipantSignature(id uint64, participant sdk.ValAddress) error
	HasSignSignature(id uint64) bool

	Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.SignSignature], *query.PageResponse, error)

	// TODO: Implement Genesis Tool
}

type kvSignatureRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

var (
	kvSignatureRepoItemPrefix = []byte{0x01}
)

func NewKVChainSignatureRepo(cdc codec.BinaryCodec, root store.KVStore, chainID string) SignatureRepo {
	return &kvSignatureRepo{
		cdc:  cdc,
		root: prefix.NewStore(root, append([]byte(chainID), kvSignatureRepoKey...)),
	}
}

func (r kvSignatureRepo) Load(id uint64) (*exported.SignSignature, error) {
	bz := prefix.NewStore(r.root, kvSignatureRepoItemPrefix).Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	signSignature := new(exported.SignSignature)
	if err := signSignature.Unmarshal(bz); err != nil {
		return nil, err
	}

	return signSignature, nil
}

func (r kvSignatureRepo) Save(signature *exported.SignSignature) error {
	bz, err := signature.Marshal()
	if err != nil {
		return err
	}

	prefix.NewStore(r.root, kvSignatureRepoItemPrefix).Set(sdk.Uint64ToBigEndian(signature.SigID), bz)
	return nil
}

func (r kvSignatureRepo) HasSignSignature(id uint64) bool {
	return prefix.NewStore(r.root, kvSignatureRepoItemPrefix).Has(sdk.Uint64ToBigEndian(id))
}

func (r kvSignatureRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, kvSignatureRepoItemPrefix)
	idBz := sdk.Uint64ToBigEndian(id)

	if !ks.Has(idBz) {
		return sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	ks.Delete(idBz)
	return nil
}

func (r kvSignatureRepo) AddParticipantSignature(id uint64, participant sdk.ValAddress, signature exported.Signature) error {
	ks := prefix.NewStore(r.root, kvSignatureRepoItemPrefix)
	idBz := sdk.Uint64ToBigEndian(id)
	bz := ks.Get(idBz)

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	signSignature := new(exported.SignSignature)
	if err := signSignature.Unmarshal(bz); err != nil {
		return err
	}

	exists := false
	for _, item := range signSignature.Items {
		if item.Participant.Equals(participant) {
			exists = true
			item.Signature = signature
			break
		}
	}
	if !exists {
		signSignature.Items = append(signSignature.Items, &exported.SignSignature_Item{
			Participant: participant,
			Signature:   signature,
		})
	}

	sigBz, err := signSignature.Marshal()
	if err != nil {
		return err
	}
	ks.Set(idBz, sigBz)
	return nil
}

func (r kvSignatureRepo) RemoveParticipantSignature(id uint64, participant sdk.ValAddress) error {
	ks := prefix.NewStore(r.root, kvSignatureRepoItemPrefix)
	idBz := sdk.Uint64ToBigEndian(id)
	bz := ks.Get(idBz)

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "sign_signature")
	}

	signSignature := new(exported.SignSignature)
	if err := signSignature.Unmarshal(bz); err != nil {
		return err
	}

	targetID := -1
	for i, item := range signSignature.Items {
		if item.Participant.Equals(participant) {
			targetID = i
			break
		}
	}

	if targetID < 0 {
		return sdkerrors.Wrap(errors.ErrNotFound, "signature")
	}
	signSignature.Items = slices.Delete(signSignature.Items, targetID, targetID+1)

	sigBz, err := signSignature.Marshal()
	if err != nil {
		return err
	}
	ks.Set(idBz, sigBz)
	return nil
}

func (r kvSignatureRepo) Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.SignSignature], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, kvSignatureRepoItemPrefix)

	var results []mitosistype.KV[uint64, *exported.SignSignature]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		signSignature := new(exported.SignSignature)
		if err := signSignature.Unmarshal(value); err != nil {
			return err
		}

		results = append(results, mitosistype.NewKV(sdk.BigEndianToUint64(key), signSignature))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
