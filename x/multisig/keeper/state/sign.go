package state

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
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

func NewKVChainSignRepo(cdc codec.BinaryCodec, root store.KVStore, chainID string) SignRepo {
	return &kvSignRepo{
		cdc:  cdc,
		root: prefix.NewStore(root, append([]byte(chainID), kvSignRepoKey...)),
	}
}

var (
	kvSignRepoLatestID   = []byte{0x01}
	kvSignRepoItemPrefix = []byte{0x02}
)

func (r kvSignRepo) Load(id uint64) (*types.Sign, error) {
	bz := prefix.NewStore(r.root, kvSignRepoItemPrefix).Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "sign")
	}

	sign := new(types.Sign)
	if err := sign.Unmarshal(bz); err != nil {
		return nil, fmt.Errorf("cannot unmarshal sign: %w", err)
	}

	return sign, nil
}

func (r kvSignRepo) Create(sign *types.Sign) (uint64, error) {
	latestIDPrefix := kvSignRepoLatestID
	latestID := sdk.BigEndianToUint64(r.root.Get(latestIDPrefix))

	sign.SigID = latestID
	signBz, err := sign.Marshal()
	if err != nil {
		return 0, sdkerrors.Wrap(errors.ErrNotFound, "cannot find sign")
	}

	prefix.NewStore(r.root, kvSignRepoItemPrefix).Set(sdk.Uint64ToBigEndian(latestID), signBz)
	r.root.Set(latestIDPrefix, sdk.Uint64ToBigEndian(latestID+1))
	return latestID, nil
}

func (r kvSignRepo) Save(sign *types.Sign) error {
	signBz, err := sign.Marshal()
	if err != nil {
		return fmt.Errorf("sign: cannot marshal. %w", err)
	}

	prefix.NewStore(r.root, kvSignRepoItemPrefix).Set(sdk.Uint64ToBigEndian(sign.SigID), signBz)
	return nil
}

func (r kvSignRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, kvSignRepoItemPrefix)

	if !ks.Has(sdk.Uint64ToBigEndian(id)) {
		return sdkerrors.Wrap(errors.ErrNotFound, "sign")
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
			return fmt.Errorf("sign: cannot unmarshal signkey %d. err: %w", key, err)
		}

		results = append(results, mitosistype.NewKV(sdk.BigEndianToUint64(key), sign))
		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("signrepo pagination: %w", err)
	}

	return results, pageResp, nil
}
