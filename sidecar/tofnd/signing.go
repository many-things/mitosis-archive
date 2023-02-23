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

	keyUID := fmt.Sprintf("#{eventKeyID}_#{0}")

	_, err := m.sign(keyUID, payload, partyUID, pubKey)
	if err != nil {
		return err
	}

	m.logger.Info(fmt.Sprintf("operator #{partyUID} sending signature for signing #{eventSigID}"))

	// TODO: Broadcast
	return nil
}
