package server

import (
	"context"
	"github.com/many-things/mitosis/x/event/keeper"
)

type msgServer struct {
	baseKeeper keeper.Keeper
}

func NewMsgServer(keeper keeper.Keeper) MsgServer {
	return msgServer{keeper}
}

func (m msgServer) Submit(ctx context.Context, submit *MsgSubmit) (*MsgSubmitResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) Vote(ctx context.Context, vote *MsgVote) (*MsgVoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) RegisterProxy(ctx context.Context, proxy *MsgRegisterProxy) (*MsgRegisterProxyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) ClearProxy(ctx context.Context, proxy *MsgClearProxy) (*MsgClearProxyResponse, error) {
	//TODO implement me
	panic("implement me")
}
