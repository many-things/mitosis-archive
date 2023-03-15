package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/multisig/types"
)

const TypeMsgStartKeygen = "start_keygen"

var _ sdk.Msg = &MsgStartKeygen{}

func NewMsgStartKeygen(creator string) *MsgStartKeygen {
	return &MsgStartKeygen{
		Creator: creator,
	}
}

func (msg *MsgStartKeygen) Route() string {
	return types.RouterKey
}

func (msg *MsgStartKeygen) Type() string {
	return TypeMsgStartKeygen
}

func (msg *MsgStartKeygen) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartKeygen) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartKeygen) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
