package server

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_ sdk.Msg = (*MsgRegisterVault)(nil)
	_ sdk.Msg = (*MsgClearVault)(nil)
)

func (m *MsgRegisterVault) ValidateBasic() error {
	return nil
}

func (m *MsgRegisterVault) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

func (m *MsgClearVault) ValidateBasic() error {
	return nil
}

func (m *MsgClearVault) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
