package keeper

import (
	"context"
	"github.com/many-things/mitosis/x/multisig/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitSignature(goCtx context.Context, msg *server.MsgSubmitSignature) (*server.MsgSubmitSignatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &server.MsgSubmitSignatureResponse{}, nil
}
