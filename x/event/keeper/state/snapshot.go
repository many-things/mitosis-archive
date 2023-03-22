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
	"github.com/pkg/errors"
	"sort"
)

type SnapshotRepo interface {
	Create(total sdk.Int, powers []mitotypes.KV[sdk.ValAddress, int64], height uint64) (*types.EpochInfo, error)

	PowerOf(epoch uint64, val sdk.ValAddress) (int64, error)

	LatestPowers() ([]mitotypes.KV[sdk.ValAddress, int64], error)

	LatestEpoch() (*types.EpochInfo, error)

	// ExportGenesis returns the entire module's state
	ExportGenesis() (genState *types.GenesisSnapshot, err error)

	// ImportGenesis sets the entire module's state
	ImportGenesis(genState *types.GenesisSnapshot) error
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

func (r kvSnapshotRepo) latestEpoch() (*types.EpochInfo, error) {
	bz := r.root.Get(kvSnapshotRepoLatestEpochKey)
	if bz == nil {
		return nil, nil
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

func (r kvSnapshotRepo) Create(total sdk.Int, powers []mitotypes.KV[sdk.ValAddress, int64], height uint64) (*types.EpochInfo, error) {
	var (
		valPowerStore = r.valPowerStore()
		valSetStore   = r.valSetStore()
	)

	latestEpoch, err := r.latestEpoch()
	if err != nil {
		return nil, err
	}

	nextEpoch := &types.EpochInfo{
		Epoch:      0, // start from zero for first epoch
		Height:     height,
		TotalPower: &total,
	}
	if latestEpoch != nil {
		nextEpoch.Epoch = latestEpoch.GetEpoch() + 1
	}

	if err := r.setLatestEpoch(nextEpoch); err != nil {
		return nil, err
	}

	sort.Slice(
		powers,
		func(i, j int) bool {
			return powers[i].Key.String() > powers[j].Key.String()
		},
	)
	for _, power := range powers {
		v, p := power.Key, power.Value

		valPowerStore.Set(
			append(v.Bytes(), sdk.Uint64ToBigEndian(nextEpoch.GetEpoch())...),
			sdk.Uint64ToBigEndian(uint64(p)),
		)
	}

	valSet := types.ValidatorSet{
		Items: mitotypes.MapKV(
			powers,
			func(val sdk.ValAddress, power int64, _ int) *types.ValidatorSet_Item {
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

	valSetStore.Set(sdk.Uint64ToBigEndian(nextEpoch.GetEpoch()), valSetBz)

	return nextEpoch, nil
}

func (r kvSnapshotRepo) PowerOf(epoch uint64, val sdk.ValAddress) (int64, error) {
	valPowerStore := prefix.NewStore(r.valPowerStore(), val.Bytes())

	bz := valPowerStore.Get(sdk.Uint64ToBigEndian(epoch))
	if bz == nil {
		return 0, errors.New("power not found")
	}

	return int64(sdk.BigEndianToUint64(bz)), nil
}

func (r kvSnapshotRepo) LatestPowers() ([]mitotypes.KV[sdk.ValAddress, int64], error) {
	valSetStore := r.valSetStore()
	latestEpoch, err := r.latestEpoch()
	if err != nil {
		return nil, err
	}
	if latestEpoch == nil {
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
		func(item *types.ValidatorSet_Item, _ int) mitotypes.KV[sdk.ValAddress, int64] {
			return mitotypes.NewKV(item.Validator, item.Power)
		},
	)

	return kvs, nil
}

func (r kvSnapshotRepo) LatestEpoch() (*types.EpochInfo, error) {
	epoch, err := r.latestEpoch()
	if err != nil {
		return nil, err
	}
	if epoch == nil {
		return nil, errors.New("epoch cannot be nil")
	}
	return epoch, nil
}

func (r kvSnapshotRepo) ExportGenesis() (genState *types.GenesisSnapshot, err error) {
	epoch, err := r.latestEpoch()
	if err != nil {
		return nil, err
	}
	if epoch == nil {
		return
	}

	genState = &types.GenesisSnapshot{}
	genState.LatestEpoch = epoch

	powerStore := r.valPowerStore()

	var powerSet []*types.GenesisSnapshot_PowerSet
	_, err = query.Paginate(
		powerStore,
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, value []byte) error {
			div := len(key) - 8
			addr := sdk.ValAddress(key[:div])
			epoch := sdk.BigEndianToUint64(key[div:])

			powerSet = append(powerSet, &types.GenesisSnapshot_PowerSet{
				Validator: addr,
				Epoch:     epoch,
				Power:     sdk.BigEndianToUint64(value),
			})
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	if len(powerSet) > 0 {
		genState.PowerSet = powerSet
	}

	return
}

func (r kvSnapshotRepo) ImportGenesis(genState *types.GenesisSnapshot) error {
	if genState.GetLatestEpoch() != nil {
		if err := r.setLatestEpoch(genState.GetLatestEpoch()); err != nil {
			return err
		}
	}

	var (
		setStore   = r.valSetStore()
		powerStore = r.valPowerStore()
	)

	sets := make(map[uint64]*types.ValidatorSet)
	for _, power := range genState.GetPowerSet() {
		if _, ok := sets[power.GetEpoch()]; ok {
			sets[power.GetEpoch()].Items = append(
				sets[power.GetEpoch()].Items,
				&types.ValidatorSet_Item{
					Validator: power.GetValidator(),
					Power:     int64(power.GetPower()),
				},
			)
		} else {
			sets[power.GetEpoch()] = &types.ValidatorSet{
				Items: []*types.ValidatorSet_Item{{
					Validator: power.GetValidator(),
					Power:     int64(power.GetPower()),
				}},
			}
		}

		powerStore.Set(
			append(power.GetValidator().Bytes(), sdk.Uint64ToBigEndian(power.GetEpoch())...),
			sdk.Uint64ToBigEndian(power.Power),
		)
	}

	for epoch, set := range sets {
		sort.Slice(
			set.Items,
			func(i, j int) bool {
				return set.Items[i].Validator.String() > set.Items[j].Validator.String()
			},
		)
		bz, err := set.Marshal()
		if err != nil {
			return err
		}
		setStore.Set(sdk.Uint64ToBigEndian(epoch), bz)
	}

	return nil
}
