package server

import (
	"context"
	"github.com/many-things/mitosis/x/multisig/keeper"
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServer(keeper keeper.Keeper) MsgServer {
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

func (m msgServer) StartSign(ctx context.Context, sign *MsgStartSign) (*MsgStartSignResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) SubmitSignature(ctx context.Context, signature *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	//TODO implement me
	panic("implement me")
}
