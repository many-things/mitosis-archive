package keeper

import "github.com/many-things/mitosis/x/context/server"

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) server.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ server.MsgServer = msgServer{}
