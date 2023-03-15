package server

import "github.com/many-things/mitosis/x/context/keeper"

type msgServer struct {
	keeper.Keeper
}

// NewMsgServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServer(keeper keeper.Keeper) MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ MsgServer = msgServer{}
