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

func (k msgServer) VoteEvent(gctx context.Context, msg *MsgVoteEvent) (*MsgVoteEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(gctx)

	// TODO: validate message
	if err := k.baseKeeper.StoreEvent(ctx, msg.GetIncoming(), msg.GetOutgoing()); err != nil {
		return nil, err
	}

	return &MsgVoteEventResponse{}, nil
}

func (k msgServer) RegisterProxy(gctx context.Context, msg *MsgRegisterProxy) (*MsgRegisterProxyResponse, error) {
	ctx := sdk.UnwrapSDKContext(gctx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	validator, err := sdk.ValAddressFromBech32(msg.GetValidator())
	if err != nil {
		panic(err.Error())
	}
	proxyAddr, err := sdk.AccAddressFromBech32(msg.GetProxyAddr())
	if err != nil {
		panic(err.Error())
	}

	if err := k.baseKeeper.RegisterProxy(ctx, validator, proxyAddr); err != nil {
		return nil, err
	}

	return &MsgRegisterProxyResponse{}, nil
}
