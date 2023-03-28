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

// Operation (ctx context.Context, operation *QueryOperation)
func (k queryServer) Operation(_ context.Context, _ *QueryOperation) (*QueryOperationResponse, error) {
	// TODO implement me
	panic("implement me")
}

// Operations // Operation (ctx context.Context, operations *QueryOperations)
func (k queryServer) Operations(_ context.Context, _ *QueryOperations) (*QueryOperationsResponse, error) {
	// TODO implement me
	panic("implement me")
}
