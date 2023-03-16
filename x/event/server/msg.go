package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/event/keeper"
	"github.com/many-things/mitosis/x/event/types"
)

type msgServer struct {
	baseKeeper    keeper.Keeper
	stakingKeeper types.StakingKeeper
}

func NewMsgServer(keeper keeper.Keeper, stakingKeeper types.StakingKeeper) MsgServer {
	return msgServer{keeper, stakingKeeper}
}

func (m msgServer) Submit(ctx context.Context, req *MsgSubmitEvent) (*MsgSubmitResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// convert proxy account to validator account
	val, err := m.baseKeeper.QueryProxyReverse(wctx, req.GetSender())
	if err != nil {
		return nil, err
	}

	pollId, err := m.baseKeeper.SubmitPoll(wctx, val, req.GetChain(), req.GetEvents())
	if err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgSubmitResponse{PollId: pollId}, nil
}

func (m msgServer) Vote(ctx context.Context, req *MsgVoteEvent) (*MsgVoteResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// convert proxy account to validator account
	val, err := m.baseKeeper.QueryProxyReverse(wctx, req.GetSender())
	if err != nil {
		return nil, err
	}

	if err := m.baseKeeper.VotePoll(wctx, val, req.GetChain(), req.GetIds()); err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgVoteResponse{}, nil
}

func (m msgServer) RegisterProxy(ctx context.Context, req *MsgRegisterProxy) (*MsgRegisterProxyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// check validator exists
	if _, found := m.stakingKeeper.GetValidator(wctx, req.GetValidator()); !found {
		return nil, errors.ErrNotFound
	}

	if err := m.baseKeeper.RegisterProxy(wctx, req.GetValidator(), req.GetProxyAccount()); err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgRegisterProxyResponse{}, nil
}

func (m msgServer) ClearProxy(ctx context.Context, req *MsgClearProxy) (*MsgClearProxyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// check validator exists
	if _, found := m.stakingKeeper.GetValidator(wctx, req.GetValidator()); !found {
		return nil, errors.ErrNotFound
	}

	if err := m.baseKeeper.ClearProxy(wctx, req.GetValidator()); err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgClearProxyResponse{}, nil
}

func (m msgServer) RegisterChain(ctx context.Context, req *MsgRegisterChain) (*MsgRegisterChainResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// convert proxy account to validator account to block non-validator execute this message
	if _, err := m.baseKeeper.QueryProxyReverse(wctx, req.GetSender()); err != nil {
		return nil, err
	}

	prefix, err := m.baseKeeper.RegisterChain(wctx, req.GetChain())
	if err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgRegisterChainResponse{ChainPrefix: []byte{prefix}}, nil
}

func (m msgServer) UnregisterChain(ctx context.Context, req *MsgUnregisterChain) (*MsgUnregisterChainResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// convert proxy account to validator account to block non-validator execute this message
	if _, err := m.baseKeeper.QueryProxyReverse(wctx, req.GetSender()); err != nil {
		return nil, err
	}

	if err := m.baseKeeper.UnregisterChain(wctx, req.GetChain()); err != nil {
		return nil, err
	}

	// TODO: Event?

	return &MsgUnregisterChainResponse{}, nil
}
