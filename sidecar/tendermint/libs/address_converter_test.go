package libs

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/assert"
	"testing"
)

func TestConvertUncompressedSecp256k1ToBech32(t *testing.T) {
	expectation := "mito12ssslxcdxyhgth8qm57wnmdy3jv7wyfvxm06g3"

	privKeyStr := "b3086285798d7447192c563e71e15486f8bd9af469b8aae13b14b6d7bca0dccf"
	privBytes, err := hex.DecodeString(privKeyStr)
	if err != nil {
		t.Error(err)
	}
	privKey := &secp256k1.PrivKey{Key: privBytes}

	result, err := sdk.Bech32ifyAddressBytes("mito", privKey.PubKey().Address().Bytes())
	assert.NilError(t, err)
	assert.Equal(t, result, expectation)
}
