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
func (k queryServer) Operation(goCtx context.Context, req *QueryOperation) (*QueryOperationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	op, err := k.Keeper.QueryOperation(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &QueryOperationResponse{Operation: op}, nil
}

// Operations // Operation (ctx context.Context, operations *QueryOperations)
func (k queryServer) Operations(goCtx context.Context, req *QueryOperations) (*QueryOperationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	ops, resp, err := k.Keeper.QueryOperations(ctx, req.GetPagination())
	if err != nil {
		return nil, err
	}

	return &QueryOperationsResponse{
		Operations: ops,
		Page:       resp,
	}, nil
}

func (k queryServer) OperationByHash(goCtx context.Context, req *QueryOperationHash) (*QueryOperationHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	op, err := k.Keeper.QueryOperationByHash(ctx, req.Chain, req.TxHash)

	if err != nil {
		return nil, err
	}

	return &QueryOperationHashResponse{
		Operation: op,
	}, nil
}
