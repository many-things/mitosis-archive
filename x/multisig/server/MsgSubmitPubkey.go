package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/x/multisig/types"
)

const TypeMsgSubmitPubkey = "submit_pubkey"

var _ sdk.Msg = &MsgSubmitPubkey{}

func NewMsgSubmitPubkey(creator string) *MsgSubmitPubkey {
	return &MsgSubmitPubkey{
		Creator: creator,
	}
}

func (msg *MsgSubmitPubkey) Route() string {
	return types.RouterKey
}

func (msg *MsgSubmitPubkey) Type() string {
	return TypeMsgSubmitPubkey
}

func (msg *MsgSubmitPubkey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitPubkey) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitPubkey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
