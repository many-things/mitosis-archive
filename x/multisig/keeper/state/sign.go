package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/types"
)

type SignRepo interface {
	Load(id uint64) (*types.Sign, error)
	Create(sign *types.Sign) (uint64, error)
	Save(sign *types.Sign) error
	Delete(id uint64) error

	Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Sign], *query.PageResponse, error)

	// TODO: Implement Genesis Tool
}

type kvSignRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVChainSignRepo(cdc codec.BinaryCodec, root store.KVStore, chainId string) SignRepo {
	return &kvSignRepo{
		cdc:  cdc,
		root: prefix.NewStore(root, append([]byte(chainId), kvSignRepoKey...)),
	}
}

var (
	kvSignRepoLatestId   = []byte{0x01}
	kvSignRepoItemPrefix = []byte{0x02}
)

func (r kvSignRepo) Load(id uint64) (*types.Sign, error) {
	bz := prefix.NewStore(r.root, kvSignRepoItemPrefix).Get(sdk.Uint64ToBigEndian(id))
	if bz != nil {
		return nil, errors.Wrap(errors.ErrNotFound, "Cannot find sign")
	}

	sign := new(types.Sign)
	if err := sign.Unmarshal(bz); err != nil {
		return nil, err
	}

	return sign, nil
}

func (r kvSignRepo) Create(sign *types.Sign) (uint64, error) {
	latestIdPrefix := kvSignRepoLatestId
	latestId := sdk.BigEndianToUint64(r.root.Get(latestIdPrefix))

	sign.SigId = latestId
	signBz, err := sign.Marshal()
	if err != nil {
		return 0, errors.Wrap(errors.ErrNotFound, "Cannot find sign")
	}

	prefix.NewStore(r.root, kvSignRepoItemPrefix).Set(sdk.Uint64ToBigEndian(latestId), signBz)
	r.root.Set(latestIdPrefix, sdk.Uint64ToBigEndian(latestId+1))
	return latestId, nil
}

func (r kvSignRepo) Save(sign *types.Sign) error {
	signBz, err := sign.Marshal()
	if err != nil {
		return err
	}

	prefix.NewStore(r.root, kvSignRepoItemPrefix).Set(sdk.Uint64ToBigEndian(sign.SigId), signBz)
	return nil
}

func (r kvSignRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, kvSignRepoItemPrefix)
	bz := ks.Get(sdk.Uint64ToBigEndian(id))

	var sign types.Sign
	if err := sign.Unmarshal(bz); err != nil {
		return err
	}

	ks.Delete(sdk.Uint64ToBigEndian(id))
	return nil
}

func (r kvSignRepo) Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Sign], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, kvSignRepoItemPrefix)

	var results []mitosistype.KV[uint64, *types.Sign]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		sign := new(types.Sign)

		if err := sign.Unmarshal(value); err != nil {
			return nil
		}

		results = append(results, mitosistype.NewKV(sdk.BigEndianToUint64(key), sign))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
