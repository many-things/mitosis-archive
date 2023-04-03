package types

type CosmosSigner struct {
	pubKey        []byte
	prefix        string
	address       string
	AccountNumber uint64
	Sequence      uint64
}

func NewCosmosSigner(pubKey []byte, prefix string, accountNumber, sequence uint64) Signer {
	return CosmosSigner{
		pubKey:        pubKey,
		prefix:        prefix,
		address:       ConvertUncompressedSecp256k1ToBech32(pubKey, prefix),
		AccountNumber: accountNumber,
		Sequence:      sequence,
	}
}

func (s CosmosSigner) Type() ChainType {
	return ChainTypeCosmos
}

func (s CosmosSigner) PubKey() []byte {
	return s.pubKey
}

func (s CosmosSigner) Address() string {
	return s.address
}

type EvmSigner struct {
	pubKey  []byte
	address string
	Nonce   uint64
}

func NewEvmSigner(pubKey []byte, nonce uint64) Signer {
	return EvmSigner{
		pubKey:  pubKey,
		address: ConvertUncompressedSecp256k1ToEth(pubKey),
		Nonce:   nonce,
	}
}

func (s EvmSigner) Type() ChainType {
	return ChainTypeEvm
}

func (s EvmSigner) PubKey() []byte {
	return s.pubKey
}

func (s EvmSigner) Address() string {
	return s.address
}

type Signer interface {
	Type() ChainType
	PubKey() []byte
	Address() string
}
