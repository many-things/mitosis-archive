package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/types"
)

func (k msgServer) StartKeygen(goCtx context.Context, msg *types.MsgStartKeygen) (*types.MsgStartKeygenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgStartKeygenResponse{}, nil
}
