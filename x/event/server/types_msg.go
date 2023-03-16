package server

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ = []sdk.Msg{
	&MsgSubmitEvent{},
	&MsgVoteEvent{},
	&MsgRegisterProxy{},
	&MsgClearProxy{},
	&MsgRegisterChain{},
	&MsgUnregisterChain{},
}

func (m *MsgSubmitEvent) ValidateBasic() error         { return nil }
func (m *MsgSubmitEvent) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }

func (m *MsgVoteEvent) ValidateBasic() error         { return nil }
func (m *MsgVoteEvent) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }

func (m *MsgRegisterProxy) ValidateBasic() error { return nil }
func (m *MsgRegisterProxy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Validator.Bytes()}
}

func (m *MsgClearProxy) ValidateBasic() error         { return nil }
func (m *MsgClearProxy) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Validator.Bytes()} }

func (m *MsgRegisterChain) ValidateBasic() error         { return nil }
func (m *MsgRegisterChain) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }

func (m *MsgUnregisterChain) ValidateBasic() error         { return nil }
func (m *MsgUnregisterChain) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{m.Sender} }
