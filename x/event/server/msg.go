package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper"
	"github.com/many-things/mitosis/x/event/types"
)

type msgServer struct {
	baseKeeper keeper.Keeper
}

func NewMsgServer(keeper keeper.Keeper) MsgServer {
	return msgServer{keeper}
}

func (k msgServer) VoteEvent(gctx context.Context, msg *types.MsgVoteEvent) (*types.MsgVoteEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(gctx)

	// TODO: validate message
	if err := k.baseKeeper.StoreEvent(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgVoteEventResponse{}, nil
}

func (k msgServer) RegisterProxy(gctx context.Context, msg *types.MsgRegisterProxy) (*types.MsgRegisterProxyResponse, error) {
	ctx := sdk.UnwrapSDKContext(gctx)

	// TODO: validate message
	_ = ctx

	return &types.MsgRegisterProxyResponse{}, nil
}
