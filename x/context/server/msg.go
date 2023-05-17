package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/context/keeper"
)

type msgServer struct {
	keeper.Keeper
}

// NewMsgServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServer(keeper keeper.Keeper) MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ MsgServer = msgServer{}

func (m msgServer) RegisterVault(ctx context.Context, msg *MsgRegisterVault) (*MsgRegisterVaultResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)
	_ = wctx

	return nil, nil
}

func (m msgServer) ClearVault(ctx context.Context, msg *MsgClearVault) (*MsgClearVaultResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)
	_ = wctx

	return nil, nil
}
