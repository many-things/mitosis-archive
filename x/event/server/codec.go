package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterServerLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgVoteEvent{}, "event/VoteEvent", nil)
	cdc.RegisterConcrete(&MsgRegisterProxy{}, "event/RegisterProxy", nil)
}

func RegisterServerInterfaces(reg cdctypes.InterfaceRegistry) {
	reg.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVoteEvent{},
		&MsgRegisterProxy{},
	)

	msgservice.RegisterMsgServiceDesc(reg, &_Msg_serviceDesc)
}
