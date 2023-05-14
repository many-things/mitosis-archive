package server

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_ sdk.Msg = (*MsgSignerReady)(nil)
	_ sdk.Msg = (*MsgRegisterCosmosSigner)(nil)
	_ sdk.Msg = (*MsgRegisterEVMSigner)(nil)
)

func (m *MsgSignerReady) ValidateBasic() error {
	return nil
}

func (m *MsgSignerReady) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

func (m *MsgRegisterCosmosSigner) ValidateBasic() error {
	return nil
}

func (m *MsgRegisterCosmosSigner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

func (m *MsgRegisterEVMSigner) ValidateBasic() error {
	return nil
}

func (m *MsgRegisterEVMSigner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
