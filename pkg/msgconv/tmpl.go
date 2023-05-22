package msgconv

import (
	"crypto/sha256"
	"github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

func assertArgs(args [][]byte, required int) error {
	if len(args) != required {
		return errors.Errorf("invalid op_args length. it needs to be %d", required)
	}
	return nil
}

type ptmplf func(src, dest, vault string, args [][]byte, funds []*types.Coin) ([]byte, error)
type ntmplf func(src, dest, vault string, args [][]byte, funds []*types.Coin) ([]byte, []byte, error)

func wrapTmplFKeccak256(f ptmplf) ntmplf {
	return func(src, dest, vault string, args [][]byte, funds []*types.Coin) ([]byte, []byte, error) {
		payload, err := f(src, dest, vault, args, funds)
		if err != nil {
			return nil, nil, err
		}

		var digest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(payload)
		hasher.Sum(digest[:0])

		return payload, digest[:], nil
	}
}

func wrapTmplFSha256(f ptmplf) ntmplf {
	return func(src, dest, vault string, args [][]byte, funds []*types.Coin) ([]byte, []byte, error) {
		payload, err := f(src, dest, vault, args, funds)
		if err != nil {
			return nil, nil, err
		}

		hash := sha256.Sum256(payload)
		return payload, hash[:], nil
	}
}

var tmpl = map[types.ChainType]map[uint64]ntmplf{
	types.ChainType_TypeCosmos: {
		0: wrapTmplFSha256(CosmosOp0),
		1: wrapTmplFSha256(CosmosOp1),
	},
	types.ChainType_TypeEvm: {
		0: wrapTmplFKeccak256(EvmOp0),
		1: wrapTmplFKeccak256(EvmOp1),
	},
}
