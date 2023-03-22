package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

var (
	_ = []sdk.Msg{
		(*MsgSubmitPubkey)(nil),
		(*MsgSubmitSignature)(nil),
		(*MsgStartKeygen)(nil),
	}
)

func (*MsgSubmitSignature) Type() string { return "MsgSubmitSignature" }
func (*MsgStartKeygen) Type() string     { return "MsgStartKeygen" }
func (*MsgSubmitPubkey) Type() string    { return "MsgSubmitPubkey" }

func (m *MsgSubmitSignature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Participant.Bytes()}
}

func (m *MsgSubmitPubkey) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Participant.Bytes()}
}

func (m *MsgStartKeygen) GetSigners() []sdk.AccAddress {
	return mitotypes.Map(m.Participants, func(t sdk.ValAddress, _ int) sdk.AccAddress { return t.Bytes() })
}

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
