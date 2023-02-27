package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVoteEvent = "vote_event"

var _ sdk.Msg = &MsgVoteEvent{}

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
