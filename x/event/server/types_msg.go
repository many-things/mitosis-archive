package server

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ = []sdk.Msg{
	&MsgSubmit{},
	&MsgVote{},
	&MsgRegisterProxy{},
	&MsgClearProxy{},
}

func (m *MsgSubmit) ValidateBasic() error         { return nil }
func (m *MsgSubmit) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }

func (m *MsgVote) ValidateBasic() error         { return nil }
func (m *MsgVote) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }

func (m *MsgRegisterProxy) ValidateBasic() error { return nil }
func (m *MsgRegisterProxy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Validator.Bytes()}
}

func (m *MsgClearProxy) ValidateBasic() error         { return nil }
func (m *MsgClearProxy) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Validator.Bytes()} }
