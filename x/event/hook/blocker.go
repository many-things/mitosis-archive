package hook

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, keeper types.BaseKeeper, stakingKeeper types.StakingKeeper) {

	var (
		total  = stakingKeeper.GetLastTotalPower(ctx)
		powers []mitotypes.KV[sdk.ValAddress, int64]
	)
	stakingKeeper.IterateLastValidatorPowers(
		ctx,
		func(addr sdk.ValAddress, power int64) bool {
			powers = append(powers, mitotypes.NewKV(addr, power))
			return false
		},
	)

	height := uint64(ctx.BlockHeight())
	params := keeper.GetParams(ctx)

	epoch, err := keeper.LatestSnapshotEpoch(ctx)
	if err != nil {
		panic(err.Error())
	}
	if epoch != nil && epoch.GetHeight()+params.GetEpochInterval() < height {
		return
	}

	_, err = keeper.CreateSnapshot(ctx, total, powers)

	// TODO: emit event
}

func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, keeper types.BaseKeeper) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}
