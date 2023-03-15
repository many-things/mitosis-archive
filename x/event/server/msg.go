package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper"
)

type msgServer struct {
	baseKeeper keeper.Keeper
}

func NewMsgServer(keeper keeper.Keeper) MsgServer {
	return msgServer{keeper}
}

func (m msgServer) Submit(ctx context.Context, req *MsgSubmit) (*MsgSubmitResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: validate request

	pollId, err := m.baseKeeper.SubmitPoll(wctx, req.GetSender(), req.GetChain(), req.GetEvents())
	if err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgSubmitResponse{PollId: pollId}, nil
}

func (m msgServer) Vote(ctx context.Context, req *MsgVote) (*MsgVoteResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: validate request

	if err := m.baseKeeper.VotePoll(wctx, req.GetSender(), req.GetChain(), req.GetIds()); err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgVoteResponse{}, nil
}

func (m msgServer) RegisterProxy(ctx context.Context, req *MsgRegisterProxy) (*MsgRegisterProxyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: validate request

	if err := m.baseKeeper.RegisterProxy(wctx, req.GetValidator(), req.GetProxyAccount()); err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgRegisterProxyResponse{}, nil
}

func (m msgServer) ClearProxy(ctx context.Context, req *MsgClearProxy) (*MsgClearProxyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: validate request

	if err := m.baseKeeper.ClearProxy(wctx, req.GetValidator()); err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgClearProxyResponse{}, nil
}
