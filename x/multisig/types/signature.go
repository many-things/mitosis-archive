package types

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/many-things/mitosis/pkg/utils"
)

type Signature []byte

func (sig Signature) ValidateBasic() error {
	_, err := btcec.ParseDERSignature(sig, btcec.S256())

	if err != nil {
		return err
	}

	return nil
}

func (sig Signature) Verify(hash Hash, pubKey PublicKey) bool {
	s, err := btcec.ParseDERSignature(sig, btcec.S256())
	if err != nil {
		return false
	}

	parsedPubKey, err := btcec.ParsePubKey(pubKey, btcec.S256())
	if err != nil {
		return false
	}

	return s.Verify(hash, parsedPubKey)
}

func (sig Signature) unpackSignature() btcec.Signature {
	return *utils.Must(btcec.ParseDERSignature(sig, btcec.S256()))
}
