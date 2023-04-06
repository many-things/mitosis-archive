package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
)

// RegisterSignature is register new signature of the sign
func (k keeper) RegisterSignature(ctx sdk.Context, chainID string, signSignature *exported.SignSignature) error {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := signatureRepo.Save(signSignature)
	if err != nil {
		return err
	}
	return nil
}

// RemoveSignature is remove the signature of the sign
func (k keeper) RemoveSignature(ctx sdk.Context, chainID string, sigID uint64) error {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := signatureRepo.Delete(sigID)
	if err != nil {
		return err
	}

	return nil
}

func (k keeper) HasSignature(ctx sdk.Context, chainID string, sigID uint64) bool {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)
	return signatureRepo.HasSignSignature(sigID)
}

func (k keeper) AddParticipantSignature(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress, signature exported.Signature) error {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := signatureRepo.AddParticipantSignature(sigID, participant, signature)
	if err != nil {
		return err
	}
	return nil
}

func (k keeper) RemoveParticipantSignature(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress) error {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := signatureRepo.RemoveParticipantSignature(sigID, participant)
	if err != nil {
		return err
	}
	return nil
}

// QuerySignature is query specific signature
func (k keeper) QuerySignature(ctx sdk.Context, chainID string, sigID uint64) (*exported.SignSignature, error) {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	signature, err := signatureRepo.Load(sigID)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// QuerySignatureList is query whole signature of specific sigID
func (k keeper) QuerySignatureList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.SignSignature], *query.PageResponse, error) {
	signatureRepo := state.NewKVChainSignatureRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	results, pageResp, err := signatureRepo.Paginate(page)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
