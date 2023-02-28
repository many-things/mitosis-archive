package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) VoteEvent(goCtx context.Context, msg *types.MsgVoteEvent) (*types.MsgVoteEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: validate message
	if err := k.Keeper.StoreEvent(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgVoteEventResponse{}, nil
}
