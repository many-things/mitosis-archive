package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func RegisterServerLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStartKeygen{}, "multisig/StartKeygen", nil)
	cdc.RegisterConcrete(&MsgSubmitPubkey{}, "multisig/SubmitPubkey", nil)
	cdc.RegisterConcrete(&MsgSubmitSignature{}, "multisig/SubmitSignature", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartKeygen{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitPubkey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitSignature{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterServerLegacyAminoCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
