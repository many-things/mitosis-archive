package types

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/many-things/mitosis/pkg/utils"
)

type KeyID string
type PublicKey []byte
type Hash []byte
type SigID string

const (
	keyIDMinLength = 4
	keyIDMaxLength = 256
)

func (k KeyID) ValidateBasic() error {
	if err := utils.ValidateString(string(k)); err != nil {
		return fmt.Errorf("KeyID: %w", err)
	}

	if !strings.Contains(string(k), "-") {
		return fmt.Errorf("keyID must format like \"string-number\"")
	}

	// length not in betgit ween [keyIDMinLength, keyIDMaxLength]
	if keyIDMinLength > len(k) || keyIDMaxLength < len(k) {
		return fmt.Errorf("keyID length must between [%d, %d]: now %d", keyIDMinLength, keyIDMaxLength, len(k))
	}

	return nil
}

// ToInternalVariables is make keyID into internal chainID / keyID
func (k KeyID) ToInternalVariables() (string, uint64, error) {
	// Expect keyID format as chainName-keyID
	splVal := strings.Split(string(k), "-")
	chainID := strings.Join(splVal[:len(splVal)-1], "-")
	id, err := strconv.ParseUint(splVal[len(splVal)-1], 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("cannot parse KeyID: %w", err)
	}

	return chainID, id, nil
}

func (s SigID) ValidateBasic() error {
	if err := utils.ValidateString(string(s)); err != nil {
		return fmt.Errorf("invalid string: %w", err)
	}

	if !strings.Contains(string(s), "-") {
		return fmt.Errorf("sigID must format like \"string-number\"")
	}

	return nil
}

// ToInternalVariables is make keyID into internal chainID / keyID
func (s SigID) ToInternalVariables() (string, uint64, error) {
	// Expect keyID format as chainName-keyID
	splVal := strings.Split(string(s), "-")
	chainID := strings.Join(splVal[:len(splVal)-1], "-")
	id, err := strconv.ParseUint(splVal[len(splVal)-1], 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("cannot parse SigID: %w", err)
	}

	return chainID, id, nil
}

func (p PublicKey) ValidateBasic() error {
	btcecPubKey, err := btcec.ParsePubKey(p, btcec.S256())
	if err != nil {
		return fmt.Errorf("publickey - cannot parse public key: %w", err)
	}

	if !bytes.Equal(p, btcecPubKey.SerializeCompressed()) {
		return fmt.Errorf("public key is not compressed")
	}

	return nil
}
