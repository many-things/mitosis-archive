package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/testutils"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/context/keeper/state"
	ctxType "github.com/many-things/mitosis/x/context/types"
	"gotest.tools/assert"
	"testing"
)

func Test_SetReadyToSigner(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	signerStore := state.NewKVSignerRepo(cdc, ctx.KVStore(storeKey))

	signer := &ctxType.Signer{
		Chain:  "osmosis-1",
		PubKey: testutils.GenPublicKey(t),
		Status: ctxType.Signer_StatusInit,
		Type:   mitotypes.ChainType_TypeCosmos,
		Payload: &ctxType.Signer_Cosmos{
			Cosmos: &ctxType.CosmosSigner{AccountNumber: 0},
		},
	}
	_ = signerStore.Save(signer)

	err := k.SetReadyToSigner(ctx, signer.Chain)
	assert.NilError(t, err)

	updatedSigner, _ := signerStore.Load("osmosis-1")
	assert.Equal(t, updatedSigner.Status, ctxType.Signer_StatusReady)
}

func Test_RegisterCosmosSigner(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	signerStore := state.NewKVSignerRepo(cdc, ctx.KVStore(storeKey))

	pubKey := testutils.GenPublicKey(t)
	err := k.RegisterCosmosSigner(ctx, "osmosis-1", pubKey, 1)
	assert.NilError(t, err)

	savedStore, _ := signerStore.Load("osmosis-1")
	assert.DeepEqual(t, savedStore.PubKey, []byte(pubKey))

	emitEvt := ctx.EventManager().Events()[0]
	expectEvt, _ := sdk.TypedEventToEvent(&ctxType.EventSignerRegistered{
		ChainType: mitotypes.ChainType_TypeCosmos,
		Pubkey:    pubKey,
	})
	assert.DeepEqual(t, emitEvt, expectEvt)
}

func Test_RegisterEVMSigner(t *testing.T) {
	k, ctx, cdc, storeKey, _ := testkeeper.ContextKeeper(t)
	signerStore := state.NewKVSignerRepo(cdc, ctx.KVStore(storeKey))

	pubKey := testutils.GenPublicKey(t)
	err := k.RegisterEVMSigner(ctx, "1", pubKey)
	assert.NilError(t, err)

	savedStore, _ := signerStore.Load("1")
	assert.DeepEqual(t, savedStore.PubKey, []byte(pubKey))

	emitEvt := ctx.EventManager().Events()[0]
	expectEvt, _ := sdk.TypedEventToEvent(&ctxType.EventSignerRegistered{
		ChainType: mitotypes.ChainType_TypeEvm,
		Pubkey:    pubKey,
	})
	assert.DeepEqual(t, emitEvt, expectEvt)
}
