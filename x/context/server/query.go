package server

import (
	"context"
	sdkerrutils "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mitotypes "github.com/many-things/mitosis/pkg/types"

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

func (k queryServer) Vault(ctx context.Context, req *QueryVault) (*QueryVaultResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	vault, err := k.Keeper.QueryVault(wctx, req.Chain)
	if err != nil {
		return nil, sdkerrutils.Wrapf(sdkerrors.ErrConflict, "failed to query vault. err=%v", err)
	}

	return &QueryVaultResponse{Chain: req.Chain, Vault: vault}, nil
}

func (k queryServer) Vaults(ctx context.Context, req *QueryVaults) (*QueryVaultsResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	vaults, pageResp, err := k.Keeper.QueryVaults(wctx, req.Pagination)
	if err != nil {
		return nil, sdkerrutils.Wrapf(sdkerrors.ErrConflict, "failed to query vaults. err=%v", err)
	}

	return &QueryVaultsResponse{
		Vaults: mitotypes.MapKV(
			vaults,
			func(k string, v string, i int) *QueryVaultResponse {
				return &QueryVaultResponse{Chain: k, Vault: v}
			},
		),
		Page: pageResp,
	}, nil
}
