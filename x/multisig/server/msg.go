package server

import (
	"context"
	"reflect"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
)

type msgServer struct {
	baseKeeper    keeper.Keeper
	contextKeeper types.ContextKeeper
	eventKeeper   types.EventKeeper
}

func NewMsgServer(keeper keeper.Keeper, contextKeeper types.ContextKeeper, eventKeeper types.EventKeeper) MsgServer {
	return msgServer{keeper, contextKeeper, eventKeeper}
}

// StartKeygen is handle MsgStartKeygen message
func (m msgServer) StartKeygen(ctx context.Context, msg *MsgStartKeygen) (*MsgStartKeygenResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	chainID, keyID, err := msg.KeyID.ToInternalVariables()
	if err != nil {
		return nil, err
	}

	// check received keygen is valid
	wctx := sdk.UnwrapSDKContext(ctx)
	kgObj, err := m.baseKeeper.QueryKeygen(wctx, chainID, keyID)
	if err != nil {
		return nil, err
	}

	if kgObj.Status > types.Keygen_StatusExecute {
		return nil, sdkerrors.Wrap(errors.ErrInvalidRequest, "keygen: cannot start finished keygen")
	} else if !reflect.DeepEqual(
		msg.Participants,
		mitosistype.Map(
			kgObj.Participants,
			func(p *types.Keygen_Participant, _ int) sdk.ValAddress { return p.Address },
		),
	) {
		return nil, sdkerrors.Wrap(errors.ErrInvalidRequest, "keygen: invalid participants")
	}

	// Change Keygen Status only Keygen Status Assigned
	if kgObj.Status == types.Keygen_StatusAssign {
		_, err := m.baseKeeper.UpdateKeygenStatus(wctx, chainID, keyID, types.Keygen_StatusExecute)
		if err != nil {
			return nil, err
		}
	}
	return &MsgStartKeygenResponse{}, nil
}

// SubmitPubkey is handle generated PublicKey
func (m msgServer) SubmitPubkey(ctx context.Context, msg *MsgSubmitPubkey) (*MsgSubmitPubkeyResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	chainID, keyID, err := msg.KeyID.ToInternalVariables()
	if err != nil {
		return nil, err
	}

	wctx := sdk.UnwrapSDKContext(ctx)

	if !m.baseKeeper.HasKeygenResult(wctx, chainID, keyID) {
		pubKey := types.KeygenResult{
			Chain: chainID,
			KeyID: keyID,
			Items: []*types.KeygenResult_Item{{
				Participant: msg.Participant,
				PubKey:      msg.PubKey,
			}},
		}

		if err := m.baseKeeper.RegisterKeygenResult(wctx, chainID, &pubKey); err != nil {
			return nil, err
		}
	} else {
		if err := m.baseKeeper.AddParticipantKeygenResult(wctx, chainID, keyID, msg.Participant, msg.PubKey); err != nil {
			return nil, err
		}
	}

	return &MsgSubmitPubkeyResponse{}, nil
}

// SubmitSignature is handle submit signature events
func (m msgServer) SubmitSignature(ctx context.Context, msg *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	chainID, sigID, err := msg.GetSigID().ToInternalVariables()
	if err != nil {
		return nil, err
	}

	wctx := sdk.UnwrapSDKContext(ctx)

	sign, err := m.baseKeeper.QuerySign(wctx, chainID, sigID)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "query sign")
	}
	if sign.Status == exported.Sign_StatusSuccess || sign.Status == exported.Sign_StatusFailed {
		return nil, sdkerrors.Wrap(errors.ErrConflict, "sign already finished / failed")
	}

	origin, found := m.eventKeeper.QueryProxyReverse(wctx, msg.Sender)
	if !found {
		return nil, sdkerrors.Wrap(errors.ErrKeyNotFound, "proxy origin not found")
	}
	if !msg.Participant.Equals(origin) {
		return nil, sdkerrors.Wrap(errors.ErrInvalidRequest, "participant address does not match with proxy origin")
	}

	if m.baseKeeper.HasSignResult(wctx, chainID, sigID) {
		if err := m.baseKeeper.AddParticipantSignResult(wctx, chainID, sigID, msg.Participant, msg.Signature); err != nil {
			return nil, err
		}
	} else {
		signature := exported.SignResult{
			Chain: chainID,
			SigID: sigID,
			Items: []*exported.SignResult_Item{{
				Participant: msg.Participant,
				Signature:   msg.Signature,
			}},
		}

		if err := m.baseKeeper.RegisterSignResult(wctx, chainID, &signature); err != nil {
			return nil, err
		}
	}

	signResult, err := m.baseKeeper.QuerySignResult(wctx, chainID, sigID)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "query sign result")
	}

	keygenChain, keygenID, err := exported.KeyID(sign.KeyID).ToInternalVariables()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "malformed keyID")
	}

	keygen, err := m.baseKeeper.QueryKeygen(wctx, keygenChain, keygenID)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "query keygen")
	}

	// check all participants are signed (TODO: threshold)
	for _, item := range signResult.Items {
		for _, p := range keygen.Participants {
			if !item.Participant.Equals(p.Address) {
				return &MsgSubmitSignatureResponse{}, nil
			}
		}
	}

	// TODO: Handle
	if sign.Status == exported.Sign_StatusExecute || sign.Status == exported.Sign_StatusAssign {
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
			if err := m.baseKeeper.SetResultSignature(wctx, chainID, sigID, sigValue[maxSignature]); err != nil {
				return nil, sdkerrors.Wrap(err, "update result signature")
			}
			if _, err := m.baseKeeper.UpdateSignStatus(wctx, chainID, sigID, exported.Sign_StatusSuccess); err != nil {
				return nil, sdkerrors.Wrap(err, "update sign status")
			}

			if err := m.contextKeeper.FinishSignOperation(wctx, sign.OpId, sigValue[maxSignature]); err != nil {
				return nil, sdkerrors.Wrap(err, "finish sign operation")
			}
		}
	}

	return &MsgSubmitSignatureResponse{}, nil
}
