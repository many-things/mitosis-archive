package server

func (*MsgSubmitSignature) Type() string { return "MsgSubmitSignature" }
func (*MsgStartKeygen) Type() string     { return "MsgStartKeygen" }
func (*MsgSubmitPubkey) Type() string    { return "MsgSubmitPubkey" }
