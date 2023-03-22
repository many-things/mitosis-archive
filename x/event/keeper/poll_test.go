package keeper_test

import (
	crand "crypto/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/event/types"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

var (
	Chains = []string{
		"mitosis-1",
		"mitosis-2",
		"mitosis-3",
		"mitosis-4",
		"mitosis-5",
		"mitosis-6",
	}
	Denoms = []string{
		"mito1",
		"mito2",
		"mito3",
		"mito4",
		"mito5",
		"mito6",
	}
)

func mockEvent(t *testing.T) *types.Event {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)

	evt := &types.Event{
		Height: rand.Uint64(),
		TxHash: bz,
		EvtIdx: rand.Uint32(),
	}
	if rand.Int()%2 == 0 {
		evt.Event = &types.Event_Req{
			Req: &types.TxReqEvent{
				DestChain: Chains[rand.Int()%len(Chains)],
				DestAddr:  bz,
				OpId:      rand.Uint32(),
				OpArgs:    [][]byte{bz},
				Funds: []*mitotypes.Coin{
					{
						Denom:   Denoms[rand.Int()%len(Denoms)],
						Amount:  mitotypes.Ref(sdk.NewInt(rand.Int63())),
						Decimal: 18,
					},
				},
			},
		}
	} else {
		evt.Event = &types.Event_Res{
			Res: &types.TxResEvent{
				ReqEvtId: rand.Uint64(),
				Ok:       rand.Int()%2 == 0,
				Result:   bz,
			},
		}
	}

	return evt
}

func TestPoll(t *testing.T) {
	k, ctx := testkeeper.EventKeeper(t)

	ctx = ctx.WithBlockHeight(123)

	_, err := k.RegisterChain(ctx, "osmosis-1")
	require.NoError(t, err)

	vals := mitotypes.Map(
		make([]byte, 2),
		func(_ byte) sdk.ValAddress {
			bz := make([]byte, 32)
			_, err = crand.Read(bz)
			require.NoError(t, err)
			return bz
		},
	)

	epoch, err := k.CreateSnapshot(
		ctx, sdk.NewInt(100),
		mitotypes.Map(
			vals,
			func(val sdk.ValAddress) mitotypes.KV[sdk.ValAddress, int64] {
				return mitotypes.NewKV(val, int64(100))
			},
		),
	)

	require.NoError(t, err)
	_ = epoch

	events := mitotypes.Map(
		make([]byte, 20),
		func(_ byte) *types.Event { return mockEvent(t) },
	)

	polls := mitotypes.Map(
		events,
		func(evt *types.Event) *types.Poll {
			return &types.Poll{
				Chain:    "osmosis-1",
				Proposer: vals[0],
				Payload:  evt,
			}
		},
	)

	submitted, err := k.SubmitPolls(ctx, "osmosis-1", vals[0], polls)
	require.NoError(t, err)

	newPolls, existPolls, err := k.FilterNewPolls(ctx, "osmosis-1", polls)
	require.NoError(t, err)
	require.Equal(t, submitted, existPolls)
	require.Equal(t, newPolls, []*types.Poll(nil))

	require.NoError(t, k.VotePolls(ctx, "osmosis-1", vals[1], mitotypes.Keys(existPolls)))

	pollsResp, _, err := k.QueryPolls(ctx, "osmosis-1", &query.PageRequest{Limit: query.MaxLimit})
	require.NoError(t, err)
	require.Equal(
		t,
		mitotypes.Map(
			make([]byte, len(polls)),
			func(_ byte) uint64 { return 200 },
		),
		mitotypes.MapKV(
			pollsResp,
			func(k uint64, v *types.Poll) uint64 { return v.Tally.Confirmed.Uint64() },
		),
	)
}
