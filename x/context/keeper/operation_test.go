package keeper_test

import (
	crand "crypto/rand"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

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
		args := [][]byte{
			testutils.GenAccAddress(t),
			testutils.GenAccAddress(t),
			[]byte("100000uosmo"),
		}

		evt.Event = &types.Event_Req{
			Req: &types.TxReqEvent{
				DestChain: "osmosis-2",
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
	_ = txconv.Converter.RegisterCosmosChain("osmosis-2", "osmosis", encoder)

	opID, err := k.InitOperation(ctx, chain, &poll)
	assert.NilError(t, err)
	hashRepo := state.NewKVOperationHashIndexRepo(cdc, ctx.KVStore(storeKey), poll.Chain)
	hashIndex, err := hashRepo.Load(poll.Payload.TxHash)

	assert.NilError(t, err)
	assert.Equal(t, opID, hashIndex)

	// Check typed event emitted
	evt := ctx.EventManager().Events()[0]
	expectEvt, _ := sdk.TypedEventToEvent(&ctxType.EventOperationInitialized{
		PollID:      poll.Id,
		OperationID: opID,
	})
	assert.DeepEqual(t, evt, expectEvt)
}

func Test_StartSignOperation(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	opRepo := state.NewKVOperationRepo(cdc, ctx.KVStore(storeKey))

	// No Operation Exists
	err := k.StartSignOperation(ctx, 1, 1)
	assert.Error(t, err, "operation not found")

	bz := make([]byte, 32)
	_, _ = crand.Read(bz)

	// Genearte Operations
	op := &ctxType.Operation{
		Chain:         "1",
		ID:            0,
		PollID:        0,
		Status:        ctxType.Operation_StatusPending,
		SignerPubkey:  testutils.GenPublicKey(t),
		TxPayload:     bz,
		TxBytesToSign: bz,
		Result:        nil,
	}

	opID, _ := opRepo.Create(op)
	err = k.StartSignOperation(ctx, opID, 1)
	assert.NilError(t, err)

	changedOp, _ := opRepo.Load(opID)
	assert.DeepEqual(t, changedOp.Status, ctxType.Operation_StatusInitSign)

	// Check Emitted Sign
	emitEvt := ctx.EventManager().Events()[0]
	expectEvt, err := sdk.TypedEventToEvent(&ctxType.EventOperationSigningStarted{
		OperationID: opID,
		SignID:      1,
		Signer:      op.SignerPubkey,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, emitEvt, expectEvt)
}

func TestKeeper_FinishSignOperation(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	opRepo := state.NewKVOperationRepo(cdc, ctx.KVStore(storeKey))

	// No Operation Exists
	signature := []byte("signature")
	err := k.FinishSignOperation(ctx, 1, signature)
	assert.Error(t, err, "operation not found")

	bz := make([]byte, 32)
	_, _ = crand.Read(bz)

	// Genearte Operations
	op := &ctxType.Operation{
		Chain:         "1",
		ID:            0,
		PollID:        0,
		Status:        ctxType.Operation_StatusInitSign,
		SignerPubkey:  testutils.GenPublicKey(t),
		TxPayload:     bz,
		TxBytesToSign: bz,
		Result:        nil,
		SigID:         1,
	}

	opID, _ := opRepo.Create(op)
	err = k.FinishSignOperation(ctx, opID, signature)
	assert.NilError(t, err)

	changedOp, _ := opRepo.Load(opID)
	assert.DeepEqual(t, changedOp.Status, ctxType.Operation_StatusFinishSign)

	// Check Emitted events
	emitEvt := ctx.EventManager().Events()[0]
	expectEvt, err := sdk.TypedEventToEvent(&ctxType.EventOperationSigningFinished{
		OperationID: opID,
		SignID:      1,
		Signer:      op.SignerPubkey,
		Signature:   signature,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, emitEvt, expectEvt)
}

func Test_FinishOperation(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	opRepo := state.NewKVOperationRepo(cdc, ctx.KVStore(storeKey))

	valAddr := testutils.GenValAddress(t)
	chain := "1"

	totalPower := sdk.NewInt(1)
	confirmed := sdk.NewInt(1)

	poll := &types.Poll{
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
		Payload: mockEvent(t, true),
	}

	// must payload type is Response
	err := k.FinishOperation(ctx, 1, poll)
	assert.Error(t, err, "invalid event payload type: panic")

	// No Operation Exists
	poll.Payload = mockEvent(t, false)
	err = k.FinishOperation(ctx, 1, poll)
	assert.Error(t, err, "operation not found")

	bz := make([]byte, 32)
	_, _ = crand.Read(bz)

	op := &ctxType.Operation{
		Chain:         "1",
		ID:            0,
		PollID:        0,
		Status:        ctxType.Operation_StatusFinishSign,
		SignerPubkey:  testutils.GenPublicKey(t),
		TxPayload:     bz,
		TxBytesToSign: bz,
		Result:        nil,
		SigID:         1,
	}

	opID, _ := opRepo.Create(op)
	err = k.FinishOperation(ctx, opID, poll)
	assert.NilError(t, err)

	changedOp, _ := opRepo.Load(opID)
	assert.DeepEqual(t, changedOp.Status, ctxType.Operation_StatusFinalized)

	// Check Emitted events
	emitEvt := ctx.EventManager().Events()[0]
	expectEvt, err := sdk.TypedEventToEvent(&ctxType.EventOperationFinalized{
		OperationID: opID,
		ReqPollID:   op.PollID,
		RespPollID:  poll.Id,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, emitEvt, expectEvt)
}

func Test_QueryOperation(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	opRepo := state.NewKVOperationRepo(cdc, ctx.KVStore(storeKey))

	_, err := k.QueryOperation(ctx, 1)
	assert.Error(t, err, "operation not found")

	bz := make([]byte, 32)
	_, _ = crand.Read(bz)

	op := &ctxType.Operation{
		Chain:         "1",
		ID:            0,
		PollID:        0,
		Status:        ctxType.Operation_StatusFinishSign,
		SignerPubkey:  testutils.GenPublicKey(t),
		TxPayload:     bz,
		TxBytesToSign: bz,
		Result:        nil,
		SigID:         1,
	}

	opID, _ := opRepo.Create(op)
	res, err := k.QueryOperation(ctx, opID)
	assert.NilError(t, err)
	assert.DeepEqual(t, op, res)
}

func Test_QueryOperationByStatus(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	opRepo := state.NewKVOperationRepo(cdc, ctx.KVStore(storeKey))

	emptyRes, _, err := k.QueryOperationsByStatus(ctx, ctxType.Operation_StatusInitSign, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	assert.Equal(t, len(emptyRes), 0)

	bz := make([]byte, 32)
	_, _ = crand.Read(bz)

	op := &ctxType.Operation{
		Chain:         "1",
		ID:            0,
		PollID:        0,
		Status:        ctxType.Operation_StatusPending,
		SignerPubkey:  testutils.GenPublicKey(t),
		TxPayload:     bz,
		TxBytesToSign: bz,
		Result:        nil,
		SigID:         1,
	}

	_, _ = opRepo.Create(op)
	res, _, err := k.QueryOperationsByStatus(ctx, ctxType.Operation_StatusPending, &query.PageRequest{Limit: query.MaxLimit})
	assert.NilError(t, err)
	assert.DeepEqual(t, res[0], op)
}

func Test_QueryOperationByHash(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	chainID := "osmosis-1"
	opRepo := state.NewKVOperationRepo(cdc, ctx.KVStore(storeKey))
	opHashRepo := state.NewKVOperationHashIndexRepo(cdc, ctx.KVStore(storeKey), chainID)

	bz := make([]byte, 32)
	_, _ = crand.Read(bz)

	_, err := k.QueryOperationByHash(ctx, chainID, bz)
	assert.Error(t, err, "hash index not found")

	op := &ctxType.Operation{
		Chain:         chainID,
		ID:            0,
		PollID:        0,
		Status:        ctxType.Operation_StatusPending,
		SignerPubkey:  testutils.GenPublicKey(t),
		TxPayload:     bz,
		TxBytesToSign: bz,
		Result:        nil,
		SigID:         1,
	}

	opID, _ := opRepo.Create(op)
	_ = opHashRepo.Create(bz, opID)
	res, err := k.QueryOperationByHash(ctx, chainID, bz)
	assert.NilError(t, err)
	assert.DeepEqual(t, res, op)
}
