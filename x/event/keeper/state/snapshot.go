package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/types"
)

type SnapshotRepo interface {
	Create(powers []mitotypes.KV[sdk.ValAddress, int64], height uint64) (*types.EpochInfo, error)

	PowerOf(epoch *uint64, val sdk.ValAddress) (int64, error)

	LatestPowerOf(val sdk.ValAddress) (int64, error)

	LatestPowers() ([]mitotypes.KV[sdk.ValAddress, int64], error)

	LatestEpoch() (*types.EpochInfo, error)
}

var (
	kvSnapshotRepoLatestEpochKey       = []byte{0x01}
	kvSnapshotRepoValidatorPowerPrefix = []byte{0x02}
	kvSnapshotRepoValidatorSetPrefix   = []byte{0x03}
)

type kvSnapshotRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVSnapshotRepo(cdc codec.BinaryCodec, root store.KVStore) SnapshotRepo {
	return &kvSnapshotRepo{
		cdc:  cdc,
		root: prefix.NewStore(root, kvSnapshotRepoKey),
	}
}

func (r kvSnapshotRepo) valPowerStore() store.KVStore {
	return prefix.NewStore(r.root, kvSnapshotRepoValidatorPowerPrefix)
}

func (r kvSnapshotRepo) valSetStore() store.KVStore {
	return prefix.NewStore(r.root, kvSnapshotRepoValidatorSetPrefix)
}

func (r kvSnapshotRepo) latestEpoch(height uint64) (*types.EpochInfo, error) {
	bz := r.root.Get(kvSnapshotRepoLatestEpochKey)
	if bz == nil {
		return &types.EpochInfo{0, height}, nil
	}

	ei := new(types.EpochInfo)
	if err := ei.Unmarshal(bz); err != nil {
		return nil, err
	}

	return ei, nil
}

func (r kvSnapshotRepo) setLatestEpoch(epoch *types.EpochInfo) error {
	bz, err := epoch.Marshal()
	if err != nil {
		return err
	}

	r.root.Set(kvSnapshotRepoLatestEpochKey, bz)

	return nil
}

func (r kvSnapshotRepo) Create(powers []mitotypes.KV[sdk.ValAddress, int64], height uint64) (*types.EpochInfo, error) {
	var (
		valPowerStore = r.valPowerStore()
		valSetStore   = r.valSetStore()
	)

	latestEpoch, err := r.latestEpoch(height)
	if err != nil {
		return nil, err
	}

	if latestEpoch.GetHeight() < height {
		latestEpoch = &types.EpochInfo{
			Epoch:  latestEpoch.GetEpoch() + 1,
			Height: height,
		}
	}

	for _, power := range powers {
		v, p := power.Key, power.Value

		valPowerStore.Set(
			append(v.Bytes(), sdk.Uint64ToBigEndian(latestEpoch.GetEpoch())...),
			sdk.Uint64ToBigEndian(uint64(p)),
		)
	}

	valSet := types.ValidatorSet{
		Items: mitotypes.MapKV(
			powers,
			func(val sdk.ValAddress, power int64) *types.ValidatorSet_Item {
				return &types.ValidatorSet_Item{
					Validator: val,
					Power:     power,
				}
			},
		),
	}
	valSetBz, err := valSet.Marshal()
	if err != nil {
		return nil, err
	}

	valSetStore.Set(sdk.Uint64ToBigEndian(latestEpoch.GetEpoch()), valSetBz)

	return latestEpoch, nil
}

func (r kvSnapshotRepo) PowerOf(epoch *uint64, val sdk.ValAddress) (int64, error) {
	valPowerStore := prefix.NewStore(r.valPowerStore(), val.Bytes())

	queryReq := &query.PageRequest{Limit: 1}
	if epoch != nil {
		queryReq.Key = sdk.Uint64ToBigEndian(*epoch)
	}

	var power uint64
	_, err := query.Paginate(
		valPowerStore,
		queryReq,
		func(key []byte, value []byte) error {
			power = sdk.BigEndianToUint64(value)
			return nil
		},
	)
	if err != nil {
		return 0, err
	}

	return int64(power), nil
}

func (r kvSnapshotRepo) LatestPowerOf(val sdk.ValAddress) (int64, error) {
	return r.PowerOf(nil, val)
}

func (r kvSnapshotRepo) LatestPowers() ([]mitotypes.KV[sdk.ValAddress, int64], error) {
	valSetStore := r.valSetStore()
	latestEpoch, err := r.latestEpoch(0)
	if err != nil {
		return nil, err
	}
	if latestEpoch.GetHeight() == 0 {
		return nil, nil
	}

	valSetBz := valSetStore.Get(sdk.Uint64ToBigEndian(latestEpoch.GetEpoch()))
	if valSetBz == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "validator set")
	}

	valSet := new(types.ValidatorSet)
	if err := valSet.Unmarshal(valSetBz); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, err.Error())
	}

	kvs := mitotypes.Map(
		valSet.Items,
		func(item *types.ValidatorSet_Item) mitotypes.KV[sdk.ValAddress, int64] {
			return mitotypes.NewKV(item.Validator, item.Power)
		},
	)

	return kvs, nil
}

func (r kvSnapshotRepo) LatestEpoch() (*types.EpochInfo, error) {
	epoch, err := r.latestEpoch(0)
	if err != nil {
		return nil, err
	}
	if epoch.GetEpoch() == 0 {
		return nil, nil
	}
	return epoch, nil
}
