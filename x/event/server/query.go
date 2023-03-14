package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k queryServer) Poll(ctx context.Context, poll *QueryPoll) (*QueryPollResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) Polls(ctx context.Context, polls *QueryPolls) (*QueryPollsResponse, error) {
	//TODO implement me
	panic("implement me")
}
