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
	Create(powers []mitotypes.KV[sdk.ValAddress, int64]) (uint64, error)

	PowerOf(epoch *uint64, val sdk.ValAddress) (int64, error)

	LatestPowerOf(val sdk.ValAddress) (int64, error)

	LatestPowers() ([]mitotypes.KV[sdk.ValAddress, int64], error)
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
	return prefix.NewStore(r.root, kvSnapshotRepoLatestEpochKey)
}

func (r kvSnapshotRepo) valSetStore() store.KVStore {
	return prefix.NewStore(r.root, kvSnapshotRepoValidatorPowerPrefix)
}

func (r kvSnapshotRepo) latestEpoch() uint64 {
	bz := r.root.Get(kvSnapshotRepoValidatorSetPrefix)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (r kvSnapshotRepo) setLatestEpoch(epoch uint64) {
	r.root.Set(kvSnapshotRepoValidatorSetPrefix, sdk.Uint64ToBigEndian(epoch))
}

func (r kvSnapshotRepo) Create(powers []mitotypes.KV[sdk.ValAddress, int64]) (uint64, error) {
	var (
		valPowerStore = r.valPowerStore()
		valSetStore   = r.valSetStore()

		latestEpoch = r.latestEpoch()
	)

	for _, power := range powers {
		v, p := power.Key, power.Value

		valPowerStore.Set(
			append(v.Bytes(), sdk.Uint64ToBigEndian(latestEpoch)...),
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
		return 0, err
	}

	valSetStore.Set(sdk.Uint64ToBigEndian(latestEpoch), valSetBz)

	r.setLatestEpoch(latestEpoch + 1)

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
	latestEpoch := r.latestEpoch()

	valSetBz := valSetStore.Get(sdk.Uint64ToBigEndian(latestEpoch))
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
