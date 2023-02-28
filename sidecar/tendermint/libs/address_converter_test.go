package libs

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/assert"
	"testing"
)

func TestConvertPubKeyToBech32Address(t *testing.T) {
	expectation := "mito17h6ufy9kmpkc8ldzzsltl26y8agm604ae6ea2r"

	privKeyStr := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	privBytes, err := hex.DecodeString(privKeyStr)
	assert.NilError(t, err)
	privKey := &secp256k1.PrivKey{Key: privBytes}

	result, err := sdk.Bech32ifyAddressBytes("mito", privKey.PubKey().Address().Bytes())
	assert.NilError(t, err)
	assert.Equal(t, result, expectation)
}
