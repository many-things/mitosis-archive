package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	baseKeeper keeper.Keeper
}

func NewQueryServer(keeper keeper.Keeper) QueryServer {
	return queryServer{keeper}
}

func (k queryServer) Params(gcx context.Context, req *QueryParams) (*QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(gcx)

	return &QueryParamsResponse{Params: k.baseKeeper.GetParams(ctx)}, nil
}

func (k queryServer) Poll(ctx context.Context, req *QueryPoll) (*QueryPollResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	poll, err := k.baseKeeper.QueryPoll(wctx, req.GetChain(), req.GetId())
	if err != nil {
		return nil, err
	}

	return &QueryPollResponse{Poll: poll}, nil
}

func (k queryServer) Polls(ctx context.Context, req *QueryPolls) (*QueryPollsResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	set, pageResp, err := k.baseKeeper.QueryPolls(wctx, req.GetChain(), req.GetPagination())
	if err != nil {
		return nil, err
	}

	return &QueryPollsResponse{
		Polls: mitotypes.Values(set),
		Page:  pageResp,
	}, nil
}

func (k queryServer) Proxy(ctx context.Context, req *QueryProxy) (*QueryProxyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	proxy, err := k.baseKeeper.QueryProxy(wctx, req.GetValidator())
	if err != nil {
		return nil, err
	}

	return &QueryProxyResponse{Validator: req.GetValidator(), ProxyAccount: proxy}, nil
}

func (k queryServer) ProxyReverse(ctx context.Context, req *QueryProxyReverse) (*QueryProxyReverseResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	validator, err := k.baseKeeper.QueryProxyReverse(wctx, req.GetProxyAccount())
	if err != nil {
		return nil, err
	}

	return &QueryProxyReverseResponse{Validator: validator}, nil
}

func (k queryServer) Proxies(ctx context.Context, req *QueryProxies) (*QueryProxiesResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	set, pageResp, err := k.baseKeeper.QueryProxies(wctx, req.GetPagination())
	if err != nil {
		return nil, err
	}

	return &QueryProxiesResponse{
		Proxies: mitotypes.Map(set, func(k sdk.ValAddress, v sdk.AccAddress) *QueryProxyResponse {
			return &QueryProxyResponse{
				Validator:    k,
				ProxyAccount: v,
			}
		}),
		Page: pageResp,
	}, nil
}

func (k queryServer) Chain(ctx context.Context, req *QueryChain) (*QueryChainResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	prefix, err := k.baseKeeper.QueryChain(wctx, req.GetChain())
	if err != nil {
		return nil, err
	}

	return &QueryChainResponse{Chain: req.GetChain(), Prefix: []byte{prefix}}, nil
}

func (k queryServer) Chains(ctx context.Context, req *QueryChains) (*QueryChainsResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	set, pageResp, err := k.baseKeeper.QueryChains(wctx, req.GetPagination())
	if err != nil {
		return nil, err
	}

	return &QueryChainsResponse{
		Chains: mitotypes.Map(set, func(k string, v byte) *QueryChainResponse {
			return &QueryChainResponse{
				Chain:  k,
				Prefix: []byte{v},
			}
		}),
		Page: pageResp,
	}, nil
}
