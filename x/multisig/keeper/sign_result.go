package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/pkg/errors"
	"strconv"
)

// RegisterSignature is register new signature of the sign
func (k keeper) RegisterSignResult(ctx sdk.Context, chainID string, signResult *exported.SignResult) error {
	signatureRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := signatureRepo.Save(signResult)
	if err != nil {
		return err
	}
	return nil
}

// RemoveSignature is remove the signature of the sign
func (k keeper) RemoveSignResult(ctx sdk.Context, chainID string, sigID uint64) error {
	signatureRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := signatureRepo.Delete(sigID)
	if err != nil {
		return err
	}

	return nil
}

func (k keeper) HasSignResult(ctx sdk.Context, chainID string, sigID uint64) bool {
	signatureRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)
	return signatureRepo.HasSignResult(sigID)
}

func (k keeper) AddParticipantSignResult(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress, signature exported.Signature) error {
	signRepo := state.NewKVChainSignRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)
	keyRepo := state.NewKVChainKeygenResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)
	signatureRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	sign, err := signRepo.Load(sigID)
	if err != nil {
		return errors.Wrap(err, "load sign")
	}

	keyID, err := strconv.Atoi(sign.KeyID)
	if err != nil {
		return errors.Wrap(err, "conv key id")
	}
	key, err := keyRepo.Load(uint64(keyID))
	if err != nil {
		return errors.Wrap(err, "load key")
	}

	var (
		pubKey exported.PublicKey
		found  = false
	)
	for _, item := range key.GetItems() {
		if item.Participant.Equals(participant) {
			pubKey = item.PubKey
			found = true
			break
		}
	}
	if !found {
		return errors.New("unknown participant")
	}

	if !signature.Verify(sign.GetMessageToSign(), pubKey) {
		return errors.New("verification failed")
	}

	if err := signatureRepo.AddParticipantSignResult(sigID, participant, signature); err != nil {
		return err
	}
	return nil
}

func (k keeper) RemoveParticipantSignResult(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress) error {
	signatureRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	err := signatureRepo.RemoveParticipantSignResult(sigID, participant)
	if err != nil {
		return err
	}
	return nil
}

// QuerySignature is query specific signature
func (k keeper) QuerySignResult(ctx sdk.Context, chainID string, sigID uint64) (*exported.SignResult, error) {
	signatureRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	signature, err := signatureRepo.Load(sigID)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// QuerySignatureList is query whole signature of specific sigID
func (k keeper) QuerySignResultList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.SignResult], *query.PageResponse, error) {
	signatureRepo := state.NewKVChainSignResultRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

	results, pageResp, err := signatureRepo.Paginate(page)
	if err != nil {
		return nil, nil, err
	}

	return results, pageResp, nil
}
