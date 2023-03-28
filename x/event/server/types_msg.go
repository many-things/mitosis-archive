package server

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var _ = []sdk.Msg{
	&MsgSubmitEvent{},
	&MsgRegisterProxy{},
	&MsgClearProxy{},
	&MsgRegisterChain{},
	&MsgUnregisterChain{},
}

func (m *MsgSubmitEvent) ValidateBasic() error {
	// verify sender
	if err := sdk.VerifyAddressFormat(m.GetSender().Bytes()); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "sender %s", m.GetSender())
	}

	// assert events size
	if len(m.GetEvents()) == 0 {
		return sdkerrors.Wrap(errors.ErrInvalidRequest, "msg must have at least 1 event")
	}
	return nil
}
func (m *MsgSubmitEvent) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }

func (m *MsgRegisterProxy) ValidateBasic() error {
	// verify validator
	if err := sdk.VerifyAddressFormat(m.GetValidator().Bytes()); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "validator %s", m.GetValidator())
	}

	// verify proxy account
	if err := sdk.VerifyAddressFormat(m.GetProxyAccount().Bytes()); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "proxy-account %s", m.GetProxyAccount())
	}
	return nil
}
func (m *MsgRegisterProxy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Validator.Bytes()}
}

func (m *MsgClearProxy) ValidateBasic() error {
	// verify validator
	if err := sdk.VerifyAddressFormat(m.GetValidator().Bytes()); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "validator %s", m.GetValidator())
	}

	return nil
}
func (m *MsgClearProxy) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Validator.Bytes()} }

func (m *MsgRegisterChain) ValidateBasic() error {
	// verify sender
	if err := sdk.VerifyAddressFormat(m.GetSender().Bytes()); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "sender %s", m.GetSender())
	}

	return nil
}
func (m *MsgRegisterChain) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }

func (m *MsgUnregisterChain) ValidateBasic() error {
	// verify sender
	if err := sdk.VerifyAddressFormat(m.GetSender().Bytes()); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "sender %s", m.GetSender())
	}

	return nil
}
func (m *MsgUnregisterChain) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }
