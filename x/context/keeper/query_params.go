package keeper

import (
	"context"
	"github.com/many-things/mitosis/x/context/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Params(goCtx context.Context, req *server.QueryParamsRequest) (*server.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &server.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
