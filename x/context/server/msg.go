package server

import (
	"context"
	sdkerrutils "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	// TODO: verify sender

	if err := m.Keeper.RegisterVault(wctx, msg.Chain, msg.VaultAddr); err != nil {
		return nil, sdkerrutils.Wrapf(sdkerrors.ErrConflict, "failed to register vault. err=%v", err)
	}

	return &MsgRegisterVaultResponse{}, nil
}

func (m msgServer) ClearVault(ctx context.Context, msg *MsgClearVault) (*MsgClearVaultResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: verify sender

	if err := m.Keeper.ClearVault(wctx, msg.Chain); err != nil {
		return nil, sdkerrutils.Wrapf(sdkerrors.ErrConflict, "failed to clear vault. err=%v", err)
	}

	return &MsgClearVaultResponse{}, nil
}
