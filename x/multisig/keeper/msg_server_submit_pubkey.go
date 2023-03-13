package keeper

import (
	"context"
	"github.com/many-things/mitosis/x/multisig/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitPubkey(goCtx context.Context, msg *server.MsgSubmitPubkey) (*server.MsgSubmitPubkeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &server.MsgSubmitPubkeyResponse{}, nil
}
