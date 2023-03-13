package keeper

import (
	"context"
	"github.com/many-things/mitosis/x/context/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContextByAddress(goCtx context.Context, req *server.QueryContextByAddressRequest) (*server.QueryContextByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &server.QueryContextByAddressResponse{}, nil
}
