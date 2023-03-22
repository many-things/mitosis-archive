package server

import sdk "github.com/cosmos/cosmos-sdk/types"

func (*MsgSubmitSignature) Type() string { return "MsgSubmitSignature" }
func (*MsgStartKeygen) Type() string     { return "MsgStartKeygen" }
func (*MsgSubmitPubkey) Type() string    { return "MsgSubmitPubkey" }

func (m *MsgStartKeygen) ValidateBasic() error {
	if err := m.GetKeyID().ValidateBasic(); err != nil {
		return err
	}

	// if one of participant
	for _, participant := range m.GetParticipants() {
		if err := sdk.VerifyAddressFormat(participant); err != nil {
			return err
		}
	}

	return nil
}

func (m *MsgSubmitPubkey) ValidateBasic() error {
	if err := m.GetKeyID().ValidateBasic(); err != nil {
		return err
	}

	if err := sdk.VerifyAddressFormat(m.GetParticipant()); err != nil {
		return err
	}

	if err := m.GetPubKey().ValidateBasic(); err != nil {
		return err
	}

	return nil
}

func (m *MsgSubmitSignature) ValidateBasic() error {
	if err := m.GetSigID().ValidateBasic(); err != nil {
		return err
	}

	if err := sdk.VerifyAddressFormat(m.GetParticipant()); err != nil {
		return err
	}

	return nil
}
