package testutils

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"github.com/btcsuite/btcd/btcec"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
)

func GenValAddress(t *testing.T) sdk.ValAddress {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)
	return bz
}

func GenAccAddress(t *testing.T) sdk.AccAddress {
	bz := make([]byte, 32)
	_, err := crand.Read(bz)
	require.NoError(t, err)
	return bz
}

func GenPublicKey(t *testing.T) types.PublicKey {
	curve := btcec.S256()
	key, err := ecdsa.GenerateKey(curve, crand.Reader)
	require.NoError(t, err)

	pubKey := btcec.PublicKey{
		Curve: curve,
		X:     key.PublicKey.X,
		Y:     key.PublicKey.Y,
	}

	return pubKey.SerializeCompressed()
}
