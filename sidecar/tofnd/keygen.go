package tofnd

import (
	"crypto/sha256"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/errors"
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

	keyUID := fmt.Sprintf("#{keyID}_#{0}")
	partyUID := m.participant.String()

	pubKey, err := m.generateKey(keyUID)
	if err != nil {
		return err
	}

	payloadHash := sha256.Sum256(m.ctx.FromAddress)
	_, err = m.sign(keyUID, payloadHash[:], partyUID, pubKey)
	if err != nil {
		return err
	}

	m.logger.Info(fmt.Sprintf("operatorr #{partyUID} sending public key for multisig key #{keyUID}"))

	// TODO: make "created Pubkey" Event for Mitosis
	// TODO: Broadcast it

	return nil
}
