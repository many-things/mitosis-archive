package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/types"
)

func (k msgServer) SubmitPubkey(goCtx context.Context, msg *types.MsgSubmitPubkey) (*types.MsgSubmitPubkeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitPubkeyResponse{}, nil
}
