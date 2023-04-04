package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/keeper/state"
	"github.com/many-things/mitosis/x/context/types"
)

var _ types.SignerKeeper = keeper{}

func (k keeper) SetReadyToSigner(ctx sdk.Context, chain string) error {
	signerStore := state.NewKVSignerRepo(k.cdc, ctx.KVStore(k.storeKey))

	signer, err := signerStore.Load(chain)
	if err != nil {
		return errors.Wrap(err, "load signer")
	}

	signer.Status = types.Signer_StatusReady
	if err := signerStore.Save(signer); err != nil {
		return errors.Wrap(err, "save signer")
	}

	return nil
}

func (k keeper) RegisterCosmosSigner(ctx sdk.Context, chain string, pubKey []byte, accountNumber uint64) error {
	signerStore := state.NewKVSignerRepo(k.cdc, ctx.KVStore(k.storeKey))

	signer := types.Signer{
		Chain:  chain,
		PubKey: pubKey,
		Status: types.Signer_StatusInit,
		Type:   mitotypes.ChainType_TypeCosmos,
		Payload: &types.Signer_Cosmos{
			Cosmos: &types.CosmosSigner{AccountNumber: accountNumber},
		},
	}
	if err := signerStore.Save(&signer); err != nil {
		return errors.Wrap(err, "save cosmos signer")
	}

	return nil
}

func (k keeper) RegisterEVMSigner(ctx sdk.Context, chain string, pubKey []byte) error {
	signerStore := state.NewKVSignerRepo(k.cdc, ctx.KVStore(k.storeKey))

	signer := types.Signer{
		Chain:  chain,
		PubKey: pubKey,
		Status: types.Signer_StatusInit,
		Type:   mitotypes.ChainType_TypeCosmos,
		Payload: &types.Signer_Evm{
			Evm: &types.EvmSigner{},
		},
	}
	if err := signerStore.Save(&signer); err != nil {
		return errors.Wrap(err, "save evm signer")
	}

	return nil
}
