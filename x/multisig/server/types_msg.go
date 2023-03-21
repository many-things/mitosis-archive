package server

func (*MsgSubmitSignature) Type() string { return "MsgSubmitSignature" }
func (*MsgStartKeygen) Type() string     { return "MsgStartKeygen" }
func (*MsgSubmitPubkey) Type() string    { return "MsgSubmitPubkey" }

func (m *MsgStartKeygen) ValidateBasic() error {
	return nil
}

func (m *MsgSubmitPubkey) ValidateBasic() error {
	return nil
}

func (m *MsgSubmitSignature) ValidateBasic() error {
	return nil
}
