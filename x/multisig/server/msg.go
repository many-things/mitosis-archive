package server

import (
	"context"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"reflect"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
)

type msgServer struct {
	baseKeeper  keeper.Keeper
	eventKeeper types.EventKeeper
}

func NewMsgServer(keeper keeper.Keeper, eventKeeper types.EventKeeper) MsgServer {
	return msgServer{keeper, eventKeeper}
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

	return &MsgSubmitSignatureResponse{}, nil
}
