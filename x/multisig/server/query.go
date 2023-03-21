package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	types.Keeper
}

func NewQueryServer(keeper types.Keeper) QueryServer {
	return queryServer{keeper}
}

func (k queryServer) Params(goCtx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
