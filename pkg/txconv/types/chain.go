package types

type ChainType string

const (
	ChainTypeCosmos ChainType = "cosmos"
	ChainTypeEvm    ChainType = "evm"
)

type TxConverter func(signer Signer, opID uint64, opArgs ...[]byte) ([]byte, []byte, error)

type ChainInfo struct {
	Type      ChainType
	ChainID   string
	ChainName string

	TxConv TxConverter
}
