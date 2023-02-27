package libs

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ConvertPubKeyToBech32Address(pubKey cryptotypes.PubKey, prefix string) (string, error) {
	bech32, err := sdk.Bech32ifyAddressBytes(prefix, pubKey.Address().Bytes())

	if err != nil {
		return "", err
	}

	return bech32, nil
}
