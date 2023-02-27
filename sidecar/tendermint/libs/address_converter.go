package libs

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/wealdtech/go-merkletree/keccak256"
	"golang.org/x/crypto/ripemd160"
)

func ConvertUncompressedSecp256k1ToEth(pubkey []byte) string {
	kHasher := keccak256.New()
	hash := kHasher.Hash(pubkey[1:])
	return "0x" + hex.EncodeToString(hash[len(hash)-20:])
}

func ConvertUncompressedSecp256k1ToBech32(pubkey []byte, prefix string) string {
	p, _ := btcec.ParsePubKey(pubkey)
	compressedPubkey := p.SerializeCompressed()

	sha := sha256.Sum256(compressedPubkey)
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha[:]) // does not error
	v := hasherRIPEMD160.Sum(nil)

	r, _ := bech32.ConvertAndEncode(prefix, v)
	return r
}
