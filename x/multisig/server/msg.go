package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
	"reflect"
)

type msgServer struct {
	baseKeeper keeper.Keeper
}

func NewMsgServer(keeper keeper.Keeper) MsgServer {
	return msgServer{keeper}
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
		return nil, errors.Wrap(errors.ErrInvalidRequest, "keygen: cannot start finished keygen")
	} else if !reflect.DeepEqual(msg.Participants, kgObj.Participants) {
		return nil, errors.Wrap(errors.ErrInvalidRequest, "keygen: invalid participants")
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
	pubKey := types.PubKey{
		Chain:       chainID,
		KeyID:       keyID,
		Participant: msg.Participant,
		PubKey:      msg.PubKey,
	}

	if err := m.baseKeeper.RegisterPubKey(wctx, chainID, &pubKey); err != nil {
		return nil, err
	}

	return &MsgSubmitPubkeyResponse{}, nil
}

// SubmitSignature is handle submit signature events
func (m msgServer) SubmitSignature(ctx context.Context, msg *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	chainID, sigId, err := msg.GetSigID().ToInternalVariables()
	if err != nil {
		return nil, err
	}

	wctx := sdk.UnwrapSDKContext(ctx)
	if err := m.baseKeeper.RegisterSignature(wctx, chainID, sigId, msg.Participant, msg.Signature); err != nil {
		return nil, err
	}
	return &MsgSubmitSignatureResponse{}, nil
}
