package types

import (
	"strconv"
	"strings"
)

type KeyID string
type PublicKey []byte
type Hash []byte
type Signature []byte

// TODO: add ValidatorBasic

// ToInternalVariables is make keyId into internal chainId / keyId
func (k KeyID) ToInternalVariables() (string, uint64, error) {
	// Expect KeyID format as chainName-KeyID
	splVal := strings.Split(string(k), "-")
	chainId := strings.Join(splVal[:len(splVal)-1], "-")
	id, err := strconv.ParseUint(splVal[len(splVal)-1], 10, 64)
	if err != nil {
		return "", 0, err
	}

	return chainId, id, nil
}
