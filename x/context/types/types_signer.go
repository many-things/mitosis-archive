package types

import (
	sdkerrutils "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/pkg/txconv"
	txconvtypes "github.com/many-things/mitosis/pkg/txconv/types"
)

func (s *Signer) ValidateBasic() error {
	chain := txconv.Converter.FindChain(s.GetChain())
	if chain == nil {
		return sdkerrutils.Wrap(sdkerrors.ErrInvalidType, "unregistered chain")
	}

	if s.Type != chain.Value.Type {
		return sdkerrutils.Wrap(sdkerrors.ErrInvalidType, "signer type does not match chain type")
	}

	return nil
}

func (s *Signer) TxConvSigner() txconvtypes.Signer {
	switch v := any(s.GetPayload()).(type) {
	case Signer_Cosmos:
		return txconvtypes.NewCosmosSigner(
			s.GetPubKey(),
			v.Cosmos.GetPrefix(),
			v.Cosmos.GetAccountNumber(),
			v.Cosmos.GetSequenceNumber(),
		)
	case Signer_Evm:
		return txconvtypes.NewEvmSigner(
			s.GetPubKey(),
			v.Evm.Nonce,
		)
	default:
		panic("unknown signer payload type")
	}
}
