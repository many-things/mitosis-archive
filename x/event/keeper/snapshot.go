package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
	"github.com/pkg/errors"
)

var _ types.SnapshotKeeper = keeper{}

func (k keeper) CreateSnapshot(ctx sdk.Context, total sdkmath.Int, powers []mitotypes.KV[sdk.ValAddress, int64]) (*types.EpochInfo, error) {
	height := uint64(ctx.BlockHeight())
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	epoch, err := snapshotRepo.Create(total, powers, height)
	if err != nil {
		return nil, err
	}

	return epoch, nil
}

func (k keeper) TotalPowerOf(ctx sdk.Context, _ *uint64) (int64, error) {
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	powers, err := snapshotRepo.LatestPowers()
	if err != nil {
		return 0, errors.Wrap(err, "fetch latest powers")
	}

	total := int64(0)
	for _, power := range mitotypes.Values(powers) {
		total += power
	}
	return total, nil
}

func (k keeper) VotingPowerOf(ctx sdk.Context, _ *uint64, val sdk.ValAddress) (int64, error) {
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	epochInfo, err := snapshotRepo.LatestEpoch()
	if err != nil {
		return 0, err
	}

	power, err := snapshotRepo.PowerOf(epochInfo.GetEpoch(), val)
	if err != nil {
		return 0, err
	}
	return power, nil
}

func (k keeper) LatestSnapshotEpoch(ctx sdk.Context) (*types.EpochInfo, error) {
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	epoch, err := snapshotRepo.LatestEpoch()
	if err != nil {
		return nil, err
	}
	return epoch, nil
}
