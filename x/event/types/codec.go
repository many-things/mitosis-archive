package types

import "github.com/cosmos/cosmos-sdk/codec"

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)
