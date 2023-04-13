package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
	"github.com/pkg/errors"
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

	keyChainID, keyID, err := exported.KeyID(sign.KeyID).ToInternalVariables()
	if err != nil {
		return fmt.Errorf("conv key id")
	} else if keyChainID != chainID {
		return fmt.Errorf("not match chain id")
	}
	key, err := keyRepo.Load(keyID)
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

	// Check validator
	if sign.Status == exported.Sign_StatusExecute {
		keygenRepo := state.NewKVChainKeygenRepo(k.cdc, ctx.KVStore(k.storeKey), chainID)

		signResult, err := signatureRepo.Load(sigID)
		if err != nil {
			return errors.Wrap(err, "err during get signature result")
		}
		keygen, err := keygenRepo.Load(keyID)
		if err != nil {
			return errors.Wrap(err, "err during check signature threshold")
		}

		sigThresh := map[string]uint64{}
		sigValue := map[string][]byte{}
		partySize := map[string]uint64{}

		for _, v := range keygen.Participants {
			partySize[v.Address.String()] = uint64(v.Share)
		}

		for _, v := range signResult.Items {
			sigValue[v.Participant.String()] = v.Signature
			signatureKey := string(v.Signature)

			if _, ok := sigThresh[signatureKey]; !ok {
				sigThresh[signatureKey] = partySize[v.Participant.String()]
			} else {
				sigThresh[signatureKey] += partySize[v.Participant.String()]
			}
		}

		var (
			maxValue     uint64
			maxSignature string
		)

		for k, v := range sigThresh {
			if v > maxValue {
				maxValue = v
				maxSignature = k
			}
		}

		if maxValue >= keygen.Threshold {
			signResult.ResultSignature = sigValue[maxSignature]
			if err := signatureRepo.Save(signResult); err != nil {
				return err
			}
			if _, err := k.UpdateSignStatus(ctx, chainID, sigID, exported.Sign_StatusComplete); err != nil {
				return err
			}
		}
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
