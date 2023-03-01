package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper"
	"github.com/many-things/mitosis/x/event/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	baseKeeper keeper.Keeper
}

func NewQueryServer(keeper keeper.Keeper) types.QueryServer {
	return queryServer{keeper}
}

func (k queryServer) Params(gcx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(gcx)

	return &types.QueryParamsResponse{Params: k.baseKeeper.GetParams(ctx)}, nil
}

func (k queryServer) VoteStatus(gcx context.Context, req *types.QueryVoteStatusRequest) (*types.QueryVoteStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(gcx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryVoteStatusResponse{}, nil
}

func (k queryServer) Proxy(gcx context.Context, req *types.QueryProxyRequest) (*types.QueryProxyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(gcx)

	// TODO: process the query
	_ = ctx

	return &types.QueryProxyResponse{}, nil
}