package keeper

import (
	"context"
	"github.com/many-things/mitosis/x/multisig/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StartKeygen(goCtx context.Context, msg *server.MsgStartKeygen) (*server.MsgStartKeygenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &server.MsgStartKeygenResponse{}, nil
}
