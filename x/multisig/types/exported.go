package types

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/many-things/mitosis/pkg/utils"
	"strconv"
	"strings"
)

type KeyID string
type PublicKey []byte
type Hash []byte
type Signature []byte
type SigID string

var (
	KeyIDMinLength = 4
	KeyIDMaxLength = 256
)

func (k KeyID) ValidateBasic() error {
	if err := utils.ValidateString(string(k)); err != nil {
		return err
	}

	if !strings.Contains(string(k), "-") {
		return fmt.Errorf("keyID must format like \"string-number\"")
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

func (s SigID) ValidateBasic() error {
	if err := utils.ValidateString(string(s)); err != nil {
		return err
	}

	if !strings.Contains(string(s), "-") {
		return fmt.Errorf("sigID must format like \"string-number\"")
	}

	return nil
}

// ToInternalVariables is make keyId into internal chainId / keyId
func (s SigID) ToInternalVariables() (string, uint64, error) {
	// Expect KeyID format as chainName-KeyID
	splVal := strings.Split(string(s), "-")
	chainId := strings.Join(splVal[:len(splVal)-1], "-")
	id, err := strconv.ParseUint(splVal[len(splVal)-1], 10, 64)
	if err != nil {
		return "", 0, err
	}

	return chainId, id, nil
}

func (p PublicKey) ValidateBasic() error {
	btcecPubKey, err := btcec.ParsePubKey(p, btcec.S256())
	if err != nil {
		return err
	}

	if !bytes.Equal(p, btcecPubKey.SerializeCompressed()) {
		return fmt.Errorf("public key is not compressed")
	}

	return nil
}
