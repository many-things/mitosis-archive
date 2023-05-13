package server

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	mitotypes "github.com/many-things/mitosis/pkg/types"
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

func (m msgServer) SubmitEvent(ctx context.Context, req *MsgSubmitEvent) (*MsgSubmitResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// basic validation of the request message
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// convert proxy account to validator account
	val, found := m.baseKeeper.QueryProxyReverse(wctx, req.GetSender())
	if !found {
		return nil, fmt.Errorf("cannot found query proxy")
	}

	// make poll candidates
	candidates := mitotypes.Map(
		req.GetEvents(),
		func(t *types.Event, _ int) *types.Poll {
			return &types.Poll{
				Chain:    req.GetChain(),
				Proposer: val,
				Payload:  t,
			}
		},
	)

	// filter requested polls
	newPolls, existPolls, err := m.baseKeeper.FilterNewPolls(wctx, req.GetChain(), candidates)
	if err != nil {
		return nil, fmt.Errorf("FilterNewPolls: %w", err)
	}

	// submits polls
	submitted, err := m.baseKeeper.SubmitPolls(wctx, req.GetChain(), val, newPolls)
	if err != nil {
		return nil, fmt.Errorf("SubmitPolls: %w", err)
	}

	// vote polls
	if err := m.baseKeeper.VotePolls(wctx, req.GetChain(), val, mitotypes.Keys(existPolls)); err != nil {
		return nil, fmt.Errorf("VotePolls: %w", err)
	}

	// [uint64, []byte] -> *types.EventType_SubmitEvent_Poll
	pollConv := func(id uint64, sum []byte, _ int) *types.EventType_SubmitEvent_Poll {
		return &types.EventType_SubmitEvent_Poll{
			Id:        id,
			EventHash: sum,
		}
	}
	// emits [EventType_SubmitEvent]
	event := &types.EventType_SubmitEvent{
		Executor:     val,
		ProxyAccount: req.GetSender(),
		Chain:        req.GetChain(),
		Submitted:    mitotypes.MapKV(submitted, pollConv),
		Voted:        mitotypes.MapKV(existPolls, pollConv),
	}
	if err = wctx.EventManager().EmitTypedEvent(event); err != nil {
		return nil, err
	}

	return &MsgSubmitResponse{
		News:  mitotypes.Keys(submitted),
		Voted: mitotypes.Keys(existPolls),
	}, nil
}

func (m msgServer) RegisterProxy(ctx context.Context, req *MsgRegisterProxy) (*MsgRegisterProxyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// basic validation of the request message
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// check validator exists
	if _, found := m.stakingKeeper.GetValidator(wctx, req.GetValidator()); !found {
		return nil, errors.ErrNotFound
	}

	// register proxy whatever if there's already a proxy registered
	if err := m.baseKeeper.RegisterProxy(wctx, req.GetValidator(), req.GetProxyAccount()); err != nil {
		return nil, err
	}

	// emits [EventType_RegisterProxy]
	event := &types.EventType_RegisterProxy{
		Executor:        req.GetValidator(),
		ProxyRegistered: req.GetProxyAccount(),
	}
	if err := wctx.EventManager().EmitTypedEvent(event); err != nil {
		return nil, err
	}

	return &MsgRegisterProxyResponse{}, nil
}

func (m msgServer) ClearProxy(ctx context.Context, req *MsgClearProxy) (*MsgClearProxyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// basic validation of the request message
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// check validator exists
	if _, found := m.stakingKeeper.GetValidator(wctx, req.GetValidator()); !found {
		return nil, errors.ErrNotFound
	}

	// check proxy account registered
	proxy, found := m.baseKeeper.QueryProxy(wctx, req.GetValidator())
	if !found {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "query proxy")
	}

	// if yes, then clear proxy
	if err := m.baseKeeper.ClearProxy(wctx, req.GetValidator()); err != nil {
		return nil, err
	}

	// emits [EventType_ClearProxy]
	event := &types.EventType_ClearProxy{
		Executor:     req.GetValidator(),
		ProxyCleared: proxy,
	}
	if err := wctx.EventManager().EmitTypedEvent(event); err != nil {
		return nil, err
	}

	return &MsgClearProxyResponse{}, nil
}

// TODO: remove this and replace to separate proposal
func (m msgServer) RegisterChain(ctx context.Context, req *MsgRegisterChain) (*MsgRegisterChainResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// basic validation of the request message
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// convert proxy account to validator account to block non-validator execute this message
	val, found := m.baseKeeper.QueryProxyReverse(wctx, req.GetSender())
	if !found {
		return nil, sdkerrors.Wrap(errors.ErrNotFound, "query proxy reverse")
	}

	// registers chain
	chainPrefix, err := m.baseKeeper.RegisterChain(wctx, req.GetChain())
	if err != nil {
		return nil, err
	}

	// emits [EventType_RegisterChain]
	event := &types.EventType_RegisterChain{
		Executor:     val,
		ProxyAccount: req.GetSender(),
		Chain:        req.GetChain(),
		ChainPrefix:  []byte{chainPrefix},
	}
	if err := wctx.EventManager().EmitTypedEvent(event); err != nil {
		return nil, err
	}

	return &MsgRegisterChainResponse{ChainPrefix: []byte{chainPrefix}}, nil
}

// TODO: remove this and replace to separate proposal
func (m msgServer) UnregisterChain(ctx context.Context, req *MsgUnregisterChain) (*MsgUnregisterChainResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// basic validation of the request message
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	// convert proxy account to validator account to block non-validator execute this message
	val, found := m.baseKeeper.QueryProxyReverse(wctx, req.GetSender())
	if !found {
		return nil, errors.ErrNotFound
	}

	// query chain prefix first
	chainPrefix, err := m.baseKeeper.QueryChain(wctx, req.GetChain())
	if err != nil {
		return nil, err
	}

	// unregisters chain
	if err := m.baseKeeper.UnregisterChain(wctx, req.GetChain()); err != nil {
		return nil, err
	}

	// emits [EventType_UnregisterChain]
	event := &types.EventType_UnregisterChain{
		Executor:     val,
		ProxyAccount: req.GetSender(),
		Chain:        req.GetChain(),
		ChainPrefix:  []byte{chainPrefix},
	}
	if err := wctx.EventManager().EmitTypedEvent(event); err != nil {
		return nil, err
	}

	return &MsgUnregisterChainResponse{}, nil
}
