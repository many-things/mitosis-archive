package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/multisig/types"
)

const TypeMsgSubmitSignature = "submit_signature"

var _ sdk.Msg = &MsgSubmitSignature{}

func NewMsgSubmitSignature(creator string) *MsgSubmitSignature {
	return &MsgSubmitSignature{
		Creator: creator,
	}
}

func (msg *MsgSubmitSignature) Route() string {
	return types.RouterKey
}

func (msg *MsgSubmitSignature) Type() string {
	return TypeMsgSubmitSignature
}

func (msg *MsgSubmitSignature) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitSignature) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitSignature) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
