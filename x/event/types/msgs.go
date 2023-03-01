package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/pkg/utils"
)

const (
	TypeMsgVoteEvent     = "vote_event"
	TypeMsgRegisterProxy = "register_proxy"
)

var (
	_ sdk.Msg = &MsgVoteEvent{}
	_ sdk.Msg = &MsgRegisterProxy{}
)

func NewMsgVoteEvent(voter string) *MsgVoteEvent {
	return &MsgVoteEvent{
		Voter: voter,
	}
}

func (msg *MsgVoteEvent) Route() string {
	return RouterKey
}

func (msg *MsgVoteEvent) Type() string {
	return TypeMsgVoteEvent
}

func (msg *MsgVoteEvent) GetSigners() []sdk.AccAddress {
	voter, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{voter}
}

func (msg *MsgVoteEvent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteEvent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid voter address (%s)", err)
	}
	return nil
}

func (msg *MsgRegisterProxy) Route() string {
	return RouterKey
}

func (msg *MsgRegisterProxy) Type() string {
	return TypeMsgRegisterProxy
}

func (msg *MsgRegisterProxy) GetSigners() []sdk.AccAddress {
	validator := utils.Unwrap1(sdk.ValAddressFromBech32, msg.Validator)
	return []sdk.AccAddress{sdk.AccAddress(validator)}
}

func (msg *MsgRegisterProxy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterProxy) ValidateBasic() error {
	if _, err := sdk.ValAddressFromBech32(msg.Validator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Proxy); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid proxy address (%s)", err)
	}
	return nil
}