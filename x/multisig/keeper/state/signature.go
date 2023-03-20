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

type SignatureRepo interface {
	Load(id uint64, participant sdk.ValAddress) (types.Signature, error)
	Save(id uint64, participant sdk.ValAddress, signature types.Signature) error
	Delete(id uint64, participant sdk.ValAddress) error

	Paginate(id uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, types.Signature], *query.PageResponse, error)

	// TODO: Implement Genesis Tool
}

type kvSignatureRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

var (
	kvSignatureRepoItemPrefix = []byte{0x01}
)

func NewKVChainSignatureRepo(cdc codec.BinaryCodec, root store.KVStore, chainId string) SignatureRepo {
	return &kvSignatureRepo{
		cdc:  cdc,
		root: prefix.NewStore(root, append([]byte(chainId), kvSignatureRepoKey...)),
	}
}

func (r kvSignatureRepo) getPrefix(prefix []byte, id uint64) []byte {
	return append(prefix, sdk.Uint64ToBigEndian(id)...)
}

func (r kvSignatureRepo) Load(id uint64, participant sdk.ValAddress) (types.Signature, error) {
	bz := prefix.NewStore(r.root, r.getPrefix(kvSignatureRepoItemPrefix, id)).Get(participant)

	if bz != nil { // TODO: not found exception
		return nil, nil
	}

	return bz, nil
}

func (r kvSignatureRepo) Save(id uint64, participant sdk.ValAddress, signature types.Signature) error {
	prefix.NewStore(r.root, r.getPrefix(kvSignatureRepoItemPrefix, id)).Set(participant, signature)
	return nil
}

func (r kvSignatureRepo) Delete(id uint64, participant sdk.ValAddress) error {
	ks := prefix.NewStore(r.root, r.getPrefix(kvSignatureRepoItemPrefix, id))
	bz := ks.Get(participant)

	if bz == nil { // TODO: not found exception
		return nil
	}

	ks.Delete(participant)
	return nil
}

func (r kvSignatureRepo) Paginate(id uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, types.Signature], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, r.getPrefix(kvSignatureRepoItemPrefix, id))

	var results []mitosistype.KV[sdk.ValAddress, types.Signature]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		results = append(results, mitosistype.NewKV(sdk.ValAddress(key), types.Signature(value)))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
