package hook

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper"
	"github.com/many-things/mitosis/x/event/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, baseKeeper keeper.Keeper, stakingKeeper types.StakingKeeper) {
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
	params := baseKeeper.GetParams(ctx)

	epoch, err := baseKeeper.LatestSnapshotEpoch(ctx)
	if err != nil {
		if !strings.Contains(err.Error(), "epoch cannot be nil") {
			panic(err.Error())
		}
	}
	if epoch != nil && epoch.GetHeight()+params.GetEpochInterval() > height {
		return
	}

	// TODO: handle error
	epoch, _ = baseKeeper.CreateSnapshot(ctx, total, powers)
	_ = epoch

	// TODO: emit event
}

func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, baseKeeper keeper.Keeper) []abci.ValidatorUpdate {
	params := baseKeeper.GetParams(ctx)

	chains, _, err := baseKeeper.QueryChains(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(err.Error())
	}

	flushed := make([][]mitotypes.KV[uint64, *types.Poll], len(chains))
	for i, chain := range chains {
		flushed[i], err = baseKeeper.FlushPolls(ctx, chain.Key, *params.PollThreshold)
		if err != nil {
			panic(err.Error())
		}
	}

	// TODO: handle flushed event

	// TODO: emit event

	return []abci.ValidatorUpdate{}
}
