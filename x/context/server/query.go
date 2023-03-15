package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/keeper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServer(keeper keeper.Keeper) QueryServer {
	return queryServer{keeper}
}

func (k queryServer) Params(goCtx context.Context, req *QueryParams) (*QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k queryServer) Operation(ctx context.Context, operation *QueryOperation) (*QueryOperationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) Operations(ctx context.Context, operations *QueryOperations) (*QueryOperationsResponse, error) {
	//TODO implement me
	panic("implement me")
}
