package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

// RegisterKeygenResult is register pubKey for specific chain and keyID. keyID included in pubKey instance
func (k keeper) RegisterKeygenResult(ctx sdk.Context, chainID string, pubKey *types.KeygenResult) error {
	pubKeyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := pubKeyRepo.Create(pubKey)
	if err != nil {
		return err
	}

	return nil
}

// DeleteKeygenResult is unregistered(delete) pubKey for specific chain, keyID and Participant.
func (k keeper) DeleteKeygenResult(ctx sdk.Context, chainID string, keyID uint64) error {
	pubKeyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := pubKeyRepo.Delete(keyID)
	if err != nil {
		return err
	}

	return nil
}

// AddParticipantKeygenResult add participant PublicKey if pubkey Initialized.
func (k keeper) AddParticipantKeygenResult(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress, publicKey exported.PublicKey) error {
	pubKeyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)
	return pubKeyRepo.AddKey(keyID, participant, publicKey)
}

// RemoveParticipantKeygenResult remove participant PublicKey if pubkey existed.
func (k keeper) RemoveParticipantKeygenResult(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress) error {
	pubKeyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)
	return pubKeyRepo.RemoveKey(keyID, participant)
}

// QueryKeygenResult is query specific pubKey via specific chain, keyID and Participant.
func (k keeper) QueryKeygenResult(ctx sdk.Context, chainID string, keyID uint64) (*types.KeygenResult, error) {
	pubKeyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	pubKey, err := pubKeyRepo.Load(keyID)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

// HasKeygenResult is check specific chains pubKey is exists
func (k keeper) HasKeygenResult(ctx sdk.Context, chainID string, keyID uint64) bool {
	pubKeyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	return pubKeyRepo.HasKeygenResult(keyID)
}

// QueryKeygenResultList is query whole pubKey via specific chain and keyID
func (k keeper) QueryKeygenResultList(ctx sdk.Context, chainID string, pageReq *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.KeygenResult], *query.PageResponse, error) {
	pubKeyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	results, pageResp, err := pubKeyRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
