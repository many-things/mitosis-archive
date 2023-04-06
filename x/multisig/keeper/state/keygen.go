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
	"github.com/many-things/mitosis/x/multisig/types"
)

type KeygenRepo interface {
	Load(id uint64) (*types.Keygen, error)
	Create(keygen *types.Keygen) (uint64, error)
	Save(keygen *types.Keygen) error
	Delete(id uint64) error

	Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error)
	ExportGenesis() (*types.GenesisKeygen_ChainSet, error)
	ImportGenesis(genState *types.GenesisKeygen_ChainSet) error
}

type kvKeygenRepo struct {
	cdc     codec.BinaryCodec
	root    store.KVStore
	chainID string
}

func NewKVChainKeygenRepo(cdc codec.BinaryCodec, root store.KVStore, chainID string) KeygenRepo {
	return &kvKeygenRepo{
		cdc:     cdc,
		root:    prefix.NewStore(root, append([]byte(chainID), kvKeygenRepoKey...)),
		chainID: chainID,
	}
}

var (
	kvKeygenRepolatestID    = []byte{0x01}
	kvKeygenRepoItemsPrefix = []byte{0x02}
)

func (r kvKeygenRepo) Load(id uint64) (*types.Keygen, error) {
	bz := prefix.NewStore(r.root, kvKeygenRepoItemsPrefix).Get(sdk.Uint64ToBigEndian(id))
	if bz == nil {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "keygen")
	}

	keygen := new(types.Keygen)
	if err := keygen.Unmarshal(bz); err != nil {
		return nil, err
	}

	return keygen, nil
}

// Create is create new keygen with new id
func (r kvKeygenRepo) Create(keygen *types.Keygen) (uint64, error) {
	latestID := sdk.BigEndianToUint64(r.root.Get(kvKeygenRepolatestID))
	latestIDBz := sdk.Uint64ToBigEndian(latestID)

	keygen.KeyID = latestID
	keygenBz, err := keygen.Marshal()
	if err != nil {
		return 0, err
	}

	prefix.NewStore(r.root, kvKeygenRepoItemsPrefix).Set(latestIDBz, keygenBz)
	r.root.Set(kvKeygenRepolatestID, sdk.Uint64ToBigEndian(latestID+1))

	return latestID, nil
}

func (r kvKeygenRepo) Save(keygen *types.Keygen) error {
	keygenBz, err := keygen.Marshal()
	if err != nil {
		return err
	}

	store := prefix.NewStore(r.root, kvKeygenRepoItemsPrefix)
	keyIDBz := sdk.Uint64ToBigEndian(keygen.KeyID)
	if !store.Has(keyIDBz) {
		return sdkerrors.Wrap(errors.ErrNotFound, "keygen")
	}

	store.Set(keyIDBz, keygenBz)
	return nil
}

func (r kvKeygenRepo) Delete(id uint64) error {
	ks := prefix.NewStore(r.root, kvKeygenRepoItemsPrefix)
	bz := ks.Get(sdk.Uint64ToBigEndian(id))

	if bz == nil {
		return sdkerrors.Wrap(errors.ErrNotFound, "keygen")
	}

	// check for obj is exists and valid
	var keygen types.Keygen
	if err := keygen.Unmarshal(bz); err != nil {
		return err
	}

	ks.Delete(sdk.Uint64ToBigEndian(id))
	return nil
}

func (r kvKeygenRepo) Paginate(page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error) {
	ks := prefix.NewStore(r.root, kvKeygenRepoItemsPrefix)

	var results []mitosistype.KV[uint64, *types.Keygen]
	pageResp, err := query.Paginate(ks, page, func(key []byte, value []byte) error {
		keygen := new(types.Keygen)

		if err := keygen.Unmarshal(value); err != nil {
			return err
		}
		results = append(results, mitosistype.NewKV(sdk.BigEndianToUint64(key), keygen))
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}

func (r kvKeygenRepo) ExportGenesis() (*types.GenesisKeygen_ChainSet, error) {
	ks := prefix.NewStore(r.root, kvKeygenRepoItemsPrefix)

	genState := &types.GenesisKeygen_ChainSet{
		Chain:  []byte(r.chainID),
		LastId: sdk.BigEndianToUint64(r.root.Get(kvKeygenRepolatestID)),
	}

	var keygenSet []*types.Keygen
	_, err := query.Paginate(
		ks,
		&query.PageRequest{Limit: query.MaxLimit},
		func(_ []byte, value []byte) error {
			keygen := new(types.Keygen)
			if err := keygen.Unmarshal(value); err != nil {
				return err
			}

			keygenSet = append(keygenSet, keygen)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return genState, nil
}

func (r kvKeygenRepo) ImportGenesis(genState *types.GenesisKeygen_ChainSet) error {
	ks := prefix.NewStore(r.root, kvKeygenRepoItemsPrefix)

	r.root.Set(kvKeygenRepolatestID, sdk.Uint64ToBigEndian(genState.LastId))
	for _, item := range genState.GetItemSet() {
		bz, err := item.Marshal()
		if err != nil {
			return err
		}
		ks.Set(sdk.Uint64ToBigEndian(item.KeyID), bz)
	}

	return nil
}
