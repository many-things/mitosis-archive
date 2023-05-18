package msgconv

import (
	"github.com/many-things/mitosis/pkg/types"
	"golang.org/x/crypto/sha3"
)

type ptmplf func(vault string, args ...[]byte) ([]byte, error)
type ntmplf func(vault string, args ...[]byte) ([]byte, []byte, error)

func wrapTmplFKeccak256(f ptmplf) ntmplf {
	return func(vault string, args ...[]byte) ([]byte, []byte, error) {
		payload, err := f(vault, args...)
		if err != nil {
			return nil, nil, err
		}
		return payload, sha3.NewLegacyKeccak256().Sum(payload), nil
	}
}

func wrapTmplFSha256(f ptmplf) ntmplf {
	return func(vault string, args ...[]byte) ([]byte, []byte, error) {
		payload, err := f(vault, args...)
		if err != nil {
			return nil, nil, err
		}
		return payload, sha3.New256().Sum(payload), nil
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
