package server

import (
	"context"
	crand "crypto/rand"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/testutils"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/stretchr/testify/assert"

	"math/rand"
	"testing"

	"github.com/many-things/mitosis/x/context/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	ctxtypes "github.com/many-things/mitosis/x/context/types"
	evttypes "github.com/many-things/mitosis/x/event/types"
	"github.com/stretchr/testify/require"
)

func mockEvent(t *testing.T, isReq bool) *evttypes.Event {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)

	evt := &evttypes.Event{
		Height: rand.Uint64(),
		TxHash: bz,
		EvtIdx: rand.Uint32(),
	}
	if isReq {
		args := [][]byte{
			[]byte(testutils.GenAccAddress(t).String()),
		}

		evt.Event = &evttypes.Event_Req{
			Req: &evttypes.TxReqEvent{
				DestChain: "osmo-test-5",
				DestAddr:  bz,
				OpId:      0,
				OpArgs:    args,
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
		evt.Event = &evttypes.Event_Res{
			Res: &evttypes.TxResEvent{
				ReqOpId: 0,
				Ok:      true,
				Result:  bz,
			},
		}
	}

	return evt
}

func setupQueryServer(t testing.TB) (keeper.Keeper, QueryServer, context.Context) {
	k, ctx, _, _, _ := testkeeper.ContextKeeper(t)
	return k, NewQueryServer(k), sdk.WrapSDKContext(ctx)
}

func TestParamsQuery(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	params := ctxtypes.DefaultParams()
	k.SetParams(wctx, params)

	response, err := s.Params(wctx, &QueryParams{})
	require.NoError(t, err)
	require.Equal(t, &QueryParamsResponse{Params: params}, response)
}

func TestOperation(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	_ = s

	wctx := ctx.(sdk.Context)
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-5", "osmo1deaddeaddeaddeaddead"))

	valAddr := testutils.GenValAddress(t)
	chain := "1"

	totalPower := sdk.NewInt(1)
	confirmed := sdk.NewInt(1)

	poll := evttypes.Poll{
		Chain:    chain,
		Id:       1,
		OpId:     0,
		Epoch:    0,
		Proposer: valAddr,
		Status:   evttypes.Poll_StatusPending,
		Tally: &evttypes.Tally{
			TotalPower: &totalPower,
			Confirmed:  &confirmed,
		},
		Payload: mockEvent(t, true),
	}

	opID, err := k.InitOperation(wctx, chain, &poll)
	require.Nil(t, err)

	op, err := k.QueryOperation(wctx, opID)
	require.Nil(t, err)

	assert.NotNil(t, op.TxPayload)
	assert.NotNil(t, op.TxBytesToSign)
}

func TestQueryVault(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-1", "osmo1deaddeaddeaddeaddead"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-2", "osmo1c0ffeec0ffeec0ffee"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-3", "osmo1beefbeefbeefbeefbeef"))

	vaultResp, err := s.Vault(wctx, &QueryVault{Chain: "osmo-test-1"})
	require.Nil(t, err)

	assert.Equal(t, vaultResp.Chain, "osmo-test-1")
	assert.Equal(t, vaultResp.Vault, "osmo1deaddeaddeaddeaddead")
}

func TestQueryVaults(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-1", "osmo1deaddeaddeaddeaddead"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-2", "osmo1c0ffeec0ffeec0ffee"))
	require.Nil(t, k.RegisterVault(wctx, "osmo-test-3", "osmo1beefbeefbeefbeefbeef"))

	vaultsResp, err := s.Vaults(wctx, &QueryVaults{Pagination: nil})
	require.Nil(t, err)

	assert.Equal(t, vaultsResp, &QueryVaultsResponse{
		Vaults: []*QueryVaultResponse{
			{
				Chain: "osmo-test-1",
				Vault: "osmo1deaddeaddeaddeaddead",
			}, {
				Chain: "osmo-test-2",
				Vault: "osmo1c0ffeec0ffeec0ffee",
			}, {
				Chain: "osmo-test-3",
				Vault: "osmo1beefbeefbeefbeefbeef",
			},
		},
		Page: &query.PageResponse{Total: 3},
	})
}
