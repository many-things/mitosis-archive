package types

import (
	"fmt"
	"github.com/many-things/mitosis/pkg/utils"
	"strconv"
	"strings"
)

type KeyID string
type PublicKey []byte
type Hash []byte
type Signature []byte

var (
	KeyIDMinLength = 4
	KeyIDMaxLength = 256
)

func (k KeyID) ValidateBasic() error {
	if err := utils.ValidateString(string(k)); err != nil {
		return err
	}

	// length not in between [KeyIDMinLength, KeyIDMaxLength]
	if KeyIDMinLength > len(k) || KeyIDMaxLength < len(k) {
		return fmt.Errorf("KeyID length must between [%d, %d]: now %d", KeyIDMinLength, KeyIDMaxLength, len(k))
	}

	return nil
}

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
