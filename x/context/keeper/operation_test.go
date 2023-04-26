package keeper_test

import (
	crand "crypto/rand"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/app"
	"github.com/many-things/mitosis/pkg/testutils"
	"github.com/many-things/mitosis/pkg/txconv"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/context/keeper/state"
	ctxType "github.com/many-things/mitosis/x/context/types"
	"github.com/many-things/mitosis/x/event/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/rand"
	"gotest.tools/assert"
)

func mockEvent(t *testing.T, isReq bool) *types.Event {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)

	evt := &types.Event{
		Height: rand.Uint64(),
		TxHash: bz,
		EvtIdx: rand.Uint32(),
	}
	if isReq {
		evt.Event = &types.Event_Req{
			Req: &types.TxReqEvent{
				DestChain: "osmosis-1",
				DestAddr:  bz,
				OpId:      0,
				OpArgs:    [][]byte{bz},
				Funds: []*mitotypes.Coin{
					{
						Denom:   "uosmo",
						Amount:  mitotypes.Ref(sdk.NewInt(rand.Int63())),
						Decimal: 18,
					},
				},
			},
		}
	} else {
		evt.Event = &types.Event_Res{
			Res: &types.TxResEvent{
				ReqEvtId: 0,
				Ok:       true,
				Result:   bz,
			},
		}
	}

	return evt
}

func Test_InitOperation(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	valAddr := testutils.GenValAddress(t)
	chain := "1"

	totalPower := sdk.NewInt(1)
	confirmed := sdk.NewInt(1)

	poll := types.Poll{
		Chain:    chain,
		Id:       1,
		OpId:     0,
		Epoch:    0,
		Proposer: valAddr,
		Status:   types.Poll_StatusPending,
		Tally: &types.Tally{
			TotalPower: &totalPower,
			Confirmed:  &confirmed,
		},
		Payload: mockEvent(t, false),
	}

	// wrong payload: response
	_, err := k.InitOperation(ctx, chain, &poll)
	assert.Error(t, err, "invalid event payload type: panic")

	poll.Payload = mockEvent(t, true)

	// signer not found
	_, err = k.InitOperation(ctx, chain, &poll)
	assert.Error(t, err, fmt.Sprintf("signer not found for chain %s: not found", poll.Chain))

	signerRepo := state.NewKVSignerRepo(cdc, ctx.KVStore(storeKey))
	_ = signerRepo.Save(&ctxType.Signer{
		Chain:  poll.GetPayload().GetReq().DestChain,
		PubKey: testutils.GenPublicKey(t),
		Status: ctxType.Signer_StatusReady,
		Type:   mitotypes.ChainType_TypeCosmos,
		Payload: &ctxType.Signer_Cosmos{
			Cosmos: &ctxType.CosmosSigner{
				Prefix:         "osmo",
				AccountNumber:  0,
				SequenceNumber: 0,
			},
		},
	})

	// support chain not found
	_, err = k.InitOperation(ctx, chain, &poll)
	assert.Error(t, err, "convert event to sign target: panic")

	encoder := app.MakeEncodingConfig().TxConfig
	_ = txconv.Converter.RegisterEvmChain(chain, "ethereum")
	_ = txconv.Converter.RegisterCosmosChain("osmosis-1", "osmosis", encoder)

	_, err = k.InitOperation(ctx, chain, &poll)
	assert.Error(t, err, "convert event to sign target: panic")
}
