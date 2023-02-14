package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
)

func (k msgServer) VoteEvent(goCtx context.Context, msg *types.MsgVoteEvent) (*types.MsgVoteEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgVoteEventResponse{}, nil
}
