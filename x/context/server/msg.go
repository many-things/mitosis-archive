package server

import (
	"context"
	sdkerrutils "cosmossdk.io/errors"
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

func (m msgServer) SignerReady(ctx context.Context, msg *MsgSignerReady) (*MsgSignerReadyResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: validate sender

	if err := m.Keeper.SetReadyToSigner(wctx, msg.GetChain()); err != nil {
		return nil, sdkerrutils.Wrap(err, "internal")
	}

	return &MsgSignerReadyResponse{}, nil
}

func (m msgServer) RegisterCosmosSigner(ctx context.Context, msg *MsgRegisterCosmosSigner) (*MsgRegisterCosmosSignerResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: validate sender

	if err := m.Keeper.RegisterCosmosSigner(wctx, msg.GetChain(), msg.GetPubKey(), msg.GetAccountNumber()); err != nil {
		return nil, sdkerrutils.Wrap(err, "internal")
	}

	return &MsgRegisterCosmosSignerResponse{}, nil
}

func (m msgServer) RegisterEVMSigner(ctx context.Context, msg *MsgRegisterEVMSigner) (*MsgRegisterEVMSignerResponse, error) {
	wctx := sdk.UnwrapSDKContext(ctx)

	// TODO: validate sender

	if err := m.Keeper.RegisterEVMSigner(wctx, msg.GetChain(), msg.GetPubKey()); err != nil {
		return nil, sdkerrutils.Wrap(err, "internal")
	}

	return &MsgRegisterEVMSignerResponse{}, nil
}
