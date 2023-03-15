package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/many-things/mitosis/x/event/types"
)

func RegisterServerLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

func RegisterServerInterfaces(reg cdctypes.InterfaceRegistry) {
	reg.RegisterImplementations((*sdk.Msg)(nil))

	msgservice.RegisterMsgServiceDesc(reg, &_Msg_serviceDesc)
}

func init() {
	RegisterServerLegacyAminoCodec(types.Amino)
	cryptocodec.RegisterCrypto(types.Amino)
	sdk.RegisterLegacyAminoCodec(types.Amino)
}
