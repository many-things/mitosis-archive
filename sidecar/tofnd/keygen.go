package tofnd

import (
	"crypto/sha256"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Gather utils stuffs kinda this
func anyOf[T any](source []T, predicate func(T) bool) bool {
	for i := range source {
		if predicate(source[i]) {
			return true
		}
	}

	return false
}

func (m Manager) ProcessKeygenStarted(participants []sdk.ValAddress, eventKeyID string) error {
	if !anyOf(participants, m.isParticipant) {
		return nil
	}

	partyUID := m.participant.String()

	pubKey, err := m.generateKey(eventKeyID)
	if err != nil {
		return err
	}

	payloadHash := sha256.Sum256(m.ctx.FromAddress)
	_, err = m.sign(eventKeyID, payloadHash[:], partyUID, pubKey)
	if err != nil {
		return err
	}

	m.logger.Info(fmt.Sprintf("operator %s sending public key for multisig key %s", partyUID, eventKeyID))

	// TODO: make "created Pubkey" Event for Mitosis
	// TODO: Broadcast it

	return nil
}
