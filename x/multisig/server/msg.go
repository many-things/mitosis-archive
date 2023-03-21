package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/multisig/types"
	"reflect"
)

type msgServer struct {
	baseKeeper types.Keeper
}

func NewMsgServer(keeper types.Keeper) MsgServer {
	return msgServer{keeper}
}

// StartKeygen is handle MsgStartKeygen message
func (m msgServer) StartKeygen(ctx context.Context, msg *MsgStartKeygen) (*MsgStartKeygenResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	chainId, keyId, err := msg.KeyID.ToInternalVariables()
	if err != nil {
		return nil, err
	}

	// check received keygen is valid
	wctx := sdk.UnwrapSDKContext(ctx)
	kgObj, err := m.baseKeeper.QueryKeygen(wctx, chainId, keyId)
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
		m.baseKeeper.UpdateKeygenStatus(wctx, chainId, keyId, types.Keygen_StatusExecute)
	}
	return &MsgStartKeygenResponse{}, nil
}

// SubmitPubkey is handle generated PublicKey
func (m msgServer) SubmitPubkey(ctx context.Context, msg *MsgSubmitPubkey) (*MsgSubmitPubkeyResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	chainId, keyId, err := msg.KeyID.ToInternalVariables()
	if err != nil {
		return nil, err
	}

	wctx := sdk.UnwrapSDKContext(ctx)
	pubKey := types.PubKey{
		Chain:       chainId,
		KeyId:       keyId,
		Participant: msg.Participant,
		PubKey:      msg.PubKey,
	}

}

func (m msgServer) SubmitSignature(ctx context.Context, signature *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	//TODO implement me
	panic("implement me")
}
