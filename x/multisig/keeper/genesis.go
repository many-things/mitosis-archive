package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

func (k keeper) ExportGenesis(ctx sdk.Context, chains []byte) (*types.GenesisState, error) {
	genState := new(types.GenesisState)

	for _, chain := range chains {
		chainStr := string(chain)

		keygenKv := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)
		keygenResKv := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)

		keygenExport, err := keygenKv.ExportGenesis()
		if err != nil {
			return nil, err
		}
		keygenResExport, err := keygenResKv.ExportGenesis()
		if err != nil {
			return nil, err
		}

		keygenExport.ResultSet = mitosistype.MapKV(
			keygenResExport,
			func(_ uint64, v *types.KeygenResult, _ int) *types.KeygenResult { return v },
		)
		genState.Keygen.ChainSet = append(genState.Keygen.ChainSet, keygenExport)

		signKv := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)
		signResKv := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainStr)

		signExport, err := signKv.ExportGenesis()
		if err != nil {
			return nil, err
		}

		signResExport, err := signResKv.ExportGenesis()
		if err != nil {
			return nil, err
		}

		signExport.ResultSet = mitosistype.MapKV(
			signResExport,
			func(_ uint64, v *exported.SignResult, _ int) *exported.SignResult { return v },
		)
		genState.Sign.ChainSet = append(genState.Sign.ChainSet, signExport)
	}
	return genState, nil
}

func (k keeper) ImportGenesis(ctx sdk.Context, genState *types.GenesisState) error {
	k.SetParams(ctx, genState.Params)

	for _, keygen := range genState.Keygen.ChainSet {
		kv := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), keygen.Chain)
		rkv := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), keygen.Chain)
		if err := kv.ImportGenesis(keygen); err != nil {
			return err
		}
		if err := rkv.ImportGenesis(
			mitosistype.Map(
				keygen.ResultSet,
				func(t *types.KeygenResult, _ int) mitosistype.KV[uint64, *types.KeygenResult] {
					return mitosistype.NewKV(t.KeyID, t)
				},
			),
		); err != nil {
			return err
		}
	}

	for _, sign := range genState.Sign.ChainSet {
		kv := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), sign.Chain)
		rkv := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), sign.Chain)
		if err := kv.ImportGenesis(sign); err != nil {
			return err
		}
		if err := rkv.ImportGenesis(
			mitosistype.Map(
				sign.ResultSet,
				func(t *exported.SignResult, _ int) mitosistype.KV[uint64, *exported.SignResult] {
					return mitosistype.NewKV(t.SigID, t)
				},
			),
		); err != nil {
			return err
		}
	}

	return nil
}
