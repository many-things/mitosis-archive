package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/many-things/mitosis/x/multisig/types"
)

// RegisterPubKey is register pubKey for specific chain and keyID. keyID included in pubKey instance
func (k keeper) RegisterPubKey(ctx sdk.Context, chainID string, pubKey *types.PubKey) error {
	pubKeyRepo := state.NewKVChainPubKeyRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := pubKeyRepo.Create(pubKey)
	if err != nil {
		return err
	}

	return nil
}

// RemovePubKey is unregistered(delete) pubKey for specific chain, keyID and Participant.
func (k keeper) RemovePubKey(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress) error {
	pubKeyRepo := state.NewKVChainPubKeyRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := pubKeyRepo.Delete(keyID, participant)
	if err != nil {
		return err
	}

	return nil
}

// QueryPubKey is query specific pubKey via specific chain, keyID and Participant.
func (k keeper) QueryPubKey(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress) (*types.PubKey, error) {
	pubKeyRepo := state.NewKVChainPubKeyRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	pubKey, err := pubKeyRepo.Load(keyID, participant)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

// QueryPubKeyList is query whole pubKey via specific chain and keyID
func (k keeper) QueryPubKeyList(ctx sdk.Context, chainID string, keyID uint64, pageReq *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.PubKey], *query.PageResponse, error) {
	pubKeyRepo := state.NewKVChainPubKeyRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	results, pageResp, err := pubKeyRepo.Paginate(keyID, pageReq)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
