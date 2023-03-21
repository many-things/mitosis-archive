package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

var _ types.SnapshotKeeper = keeper{}

func (k keeper) CreateSnapshot(ctx sdk.Context, snapshot []mitotypes.KV[sdk.ValAddress, int64]) (*types.EpochInfo, error) {
	height := uint64(ctx.BlockHeight())
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	epoch, err := snapshotRepo.Create(snapshot, height)
	if err != nil {
		return nil, err
	}

	return epoch, nil
}

func (k keeper) LatestSnapshotEpoch(ctx sdk.Context) (*types.EpochInfo, error) {
	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	epoch, err := snapshotRepo.LatestEpoch()
	if err != nil {
		return nil, err
	}
	return epoch, nil
}
