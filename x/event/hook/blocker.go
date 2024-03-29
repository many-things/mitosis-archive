package hook

import (
	"fmt"
	"log"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper"
	"github.com/many-things/mitosis/x/event/types"
	abci "github.com/tendermint/tendermint/abci/types"
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
	epoch, err = baseKeeper.CreateSnapshot(ctx, total, powers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(epoch)
	_ = epoch

	// TODO: emit event
}

func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, baseKeeper keeper.Keeper, contextKeeper types.ContextKeeper) []abci.ValidatorUpdate {
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

		for _, kv := range flushed[i] {
			pollID, poll := kv.Key, kv.Value

			switch v := poll.GetPayload().Event.(type) {
			case *types.Event_Req:
				opID, err := contextKeeper.InitOperation(ctx, chain.Key, poll)
				if err != nil {
					panic(err.Error())
				}

				// TODO: log correctly
				log.Println("init operation", opID, "for poll", pollID)
				poll.OpId = opID
			case *types.Event_Res:
				originPoll, err := baseKeeper.QueryPoll(ctx, chain.Key, v.Res.ReqOpId)
				if err != nil {
					panic(err.Error())
				}

				if err := contextKeeper.FinishOperation(ctx, originPoll.OpId, poll); err != nil {
					panic(err.Error())
				}
			default:
				panic("unexpected payload type")
			}
		}
	}

	// TODO: handle flushed event

	// TODO: emit event

	return []abci.ValidatorUpdate{}
}
