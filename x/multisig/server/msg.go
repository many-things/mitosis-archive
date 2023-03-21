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
func (m msgServer) StartKeygen(ctx context.Context, keygen *MsgStartKeygen) (*MsgStartKeygenResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)
	chainId, keyId, err := keygen.KeyID.ToInternalVariables()
	if err != nil {
		return nil, err
	}

	// check received keygen is valid
	kgObj, err := m.baseKeeper.QueryKeygen(wctx, chainId, keyId)
	if err != nil {
		return nil, err
	}

	if kgObj.Status > types.Keygen_StatusExecute {
		return nil, errors.Wrap(errors.ErrInvalidRequest, "keygen: cannot start finished keygen")
	} else if !reflect.DeepEqual(keygen.Participants, kgObj.Participants) {
		return nil, errors.Wrap(errors.ErrInvalidRequest, "keygen: invalid participants")
	}

	// Change Keygen Status only Keygen Status Assigned
	if kgObj.Status == types.Keygen_StatusAssign {
		m.baseKeeper.UpdateKeygenStatus(wctx, chainId, keyId, types.Keygen_StatusExecute)
	}
	return &MsgStartKeygenResponse{}, nil
}

func (m msgServer) SubmitPubkey(ctx context.Context, pubkey *MsgSubmitPubkey) (*MsgSubmitPubkeyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) SubmitSignature(ctx context.Context, signature *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	//TODO implement me
	panic("implement me")
}
