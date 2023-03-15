package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/keeper"
)

type msgServer struct {
	keeper.Keeper
}

// NewMsgServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServer(keeper keeper.Keeper) MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) StartKeygen(goCtx context.Context, msg *MsgStartKeygen) (*MsgStartKeygenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &MsgStartKeygenResponse{}, nil
}

func (k msgServer) SubmitPubkey(goCtx context.Context, msg *MsgSubmitPubkey) (*MsgSubmitPubkeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &MsgSubmitPubkeyResponse{}, nil
}

func (k msgServer) SubmitSignature(goCtx context.Context, msg *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &MsgSubmitSignatureResponse{}, nil
}
