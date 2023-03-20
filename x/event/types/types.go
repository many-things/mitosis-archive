package types

import (
	"github.com/tendermint/crypto/sha3"
)

func (e *Event) Hash() ([]byte, error) {
	bz, err := e.Marshal()
	if err != nil {
		return nil, err
	}

	digest := sha3.Sum256(bz)
	return digest[:], nil
}
