package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

// RegisterSignature is register new signature of the sign
func (k keeper) RegisterSignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress, signature types.Signature) error {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	err := signatureRepo.Create(sigId, participant, signature)
	if err != nil {
		return err
	}
	return nil
}

// RemoveSignature is remove the signature of the sign
func (k keeper) RemoveSignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress) error {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	err := signatureRepo.Delete(sigId, participant)
	if err != nil {
		return nil
	}
	return err
}

// QuerySignature is query specific signature
func (k keeper) QuerySignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress) (types.Signature, error) {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	signature, err := signatureRepo.Load(sigId, participant)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// QuerySignatureList is query whole signature of specific sigId
func (k keeper) QuerySignatureList(ctx sdk.Context, chainId string, sigId uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, types.Signature], *query.PageResponse, error) {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainId)

	results, pageResp, err := signatureRepo.Paginate(sigId, page)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
