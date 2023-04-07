package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

func (k keeper) ExportGenesis(ctx sdk.Context, chains []byte) (*types.GenesisState, error) {
	genState := new(types.GenesisState)
	
	for _, chain := range chains {
		chainStr := string(chain)

		keygenKv := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)
		keygenExport, err := keygenKv.ExportGenesis()
		if err != nil {
			return nil, err
		}
		genState.Keygen.ChainSet = append(genState.Keygen.ChainSet, keygenExport)

		pubKeyKv := state.NewKVChainPubKeyRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)
		pubKeyExport, err := pubKeyKv.ExportGenesis()
		if err != nil {
			return nil, err
		}
		genState.PubKey.ChainSet = append(genState.PubKey.ChainSet, pubKeyExport)

		signKv := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)
		signExport, err := signKv.ExportGenesis()
		if err != nil {
			return nil, err
		}
		genState.Sign.ChainSet = append(genState.Sign.ChainSet, signExport)

		signatureKv := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)
		signatureExport, err := signatureKv.ExportGenesis()
		if err != nil {
			return nil, err
		}
		genState.Signature.ChainSet = append(genState.Signature.ChainSet, signatureExport)

	}
	return genState, nil
}

func (k keeper) ImportGenesis(ctx sdk.Context, genState *types.GenesisState) error {
	for _, keygen := range genState.Keygen.ChainSet {
		kv := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), keygen.Chain)
		if err := kv.ImportGenesis(keygen); err != nil {
			return err
		}
	}

	for _, pubkey := range genState.PubKey.ChainSet {
		kv := state.NewKVChainPubKeyRepo(k.cdc, ctx.KVStore(k.storeKey), pubkey.Chain)
		if err := kv.ImportGenesis(pubkey); err != nil {
			return err
		}
	}

	for _, sign := range genState.Sign.ChainSet {
		kv := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), sign.Chain)
		if err := kv.ImportGenesis(sign); err != nil {
			return err
		}
	}

	for _, signature := range genState.Signature.ChainSet {
		kv := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), signature.Chain)
		if err := kv.ImportGenesis(signature); err != nil {
			return err
		}
	}

	return nil
}
