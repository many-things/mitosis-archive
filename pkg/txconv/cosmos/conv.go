package cosmos

import (
	sdkerrutils "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	txconvtypes "github.com/many-things/mitosis/pkg/txconv/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

func convert(chainID string, encoder client.TxConfig, config ChainInfoConfig) txconvtypes.TxConverter {
	return func(s txconvtypes.Signer, opID uint64, opArgs ...[]byte) ([]byte, []byte, error) {
		signer := s.(txconvtypes.CosmosSigner)

		op := mitotypes.FindKV(
			config.ops,
			func(k uint64, _ Operation, _ int) bool { return k == opID },
		)
		if op == nil {
			return nil, nil, sdkerrutils.Wrap(sdkerrors.ErrPanic, "operation not found")
		}

		builder := encoder.NewTxBuilder()

		msgs, err := op.Value.Msgs(opArgs...)
		if err != nil {
			return nil, nil, err
		}

		if err := builder.SetMsgs(msgs...); err != nil {
			return nil, nil, err
		}

		builder.SetFeeAmount(config.fee)
		builder.SetGasLimit(op.Value.Gas)

		signMode := signing.SignMode_SIGN_MODE_DIRECT

		signerData := authsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: signer.AccountNumber,
			Sequence:      signer.Sequence,
		}
		sig := signing.SignatureV2{
			PubKey:   &secp256k1.PubKey{Key: signer.PubKey()},
			Data:     &signing.SingleSignatureData{SignMode: signMode},
			Sequence: signer.Sequence,
		}
		if err := builder.SetSignatures(sig); err != nil {
			return nil, nil, err
		}

		// Generate the bytes to be signed.
		bytesToSign, err := encoder.
			SignModeHandler().
			GetSignBytes(
				signing.SignMode_SIGN_MODE_DIRECT,
				signerData,
				builder.GetTx(),
			)
		if err != nil {
			return nil, nil, err
		}

		encodedTx, err := encoder.TxEncoder()(builder.GetTx())
		if err != nil {
			return nil, nil, err
		}

		return encodedTx, bytesToSign, nil
	}
}
