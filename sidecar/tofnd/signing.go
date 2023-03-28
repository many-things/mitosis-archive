package tofnd

import (
	"fmt"

	"github.com/many-things/mitosis/sidecar/types"
)

func (m *Manager) ProcessSingingStarted(eventKeyID string, eventSigID uint64, payload types.Hash, pubKeys map[string]types.PublicKey) error {
	partyUID := m.participant.String()

	pubKey, ok := pubKeys[partyUID]
	if !ok {
		return nil // TODO: Return Notfound Message
	}

	keyUID := fmt.Sprintf("%s_%d", eventKeyID, eventSigID)

	_, err := m.sign(keyUID, payload, partyUID, pubKey)
	if err != nil {
		return err
	}

	m.logger.Info(fmt.Sprintf("operator %s sending signature for signing %d", partyUID, eventSigID))

	// TODO: Broadcast
	return nil
}
