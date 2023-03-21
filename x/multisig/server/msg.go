package server

import (
	"context"
	"github.com/many-things/mitosis/x/multisig/types"
)

type msgServer struct {
	types.Keeper
}

func NewMsgServer(keeper types.Keeper) MsgServer {
	return msgServer{keeper}
}

func (m msgServer) StartKeygen(ctx context.Context, keygen *MsgStartKeygen) (*MsgStartKeygenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) SubmitPubkey(ctx context.Context, pubkey *MsgSubmitPubkey) (*MsgSubmitPubkeyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) SubmitSignature(ctx context.Context, signature *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	//TODO implement me
	panic("implement me")
}
