package types

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func ConvertUncompressedSecp256k1ToEth(pubkey []byte) string {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(pubkey[1:])
	hash := hasher.Sum(nil)
	return "0x" + hex.EncodeToString(hash[len(hash)-20:])
}

func ConvertUncompressedSecp256k1ToBech32(pubkey []byte, prefix string) string {
	p, _ := btcec.ParsePubKey(pubkey)
	compressedPubkey := p.SerializeCompressed()

	sha := sha3.Sum256(compressedPubkey)
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha[:]) // does not error
	v := hasherRIPEMD160.Sum(nil)

	r, _ := bech32.ConvertAndEncode(prefix, v)
	return r
}
