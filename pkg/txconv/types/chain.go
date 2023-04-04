package types

import mitotypes "github.com/many-things/mitosis/pkg/types"

type TxConverter func(signer Signer, opID uint64, opArgs ...[]byte) ([]byte, []byte, error)

type ChainInfo struct {
	Type      mitotypes.ChainType
	ChainID   string
	ChainName string

	TxConv TxConverter
}
