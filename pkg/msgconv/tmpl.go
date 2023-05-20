package msgconv

import (
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

type ptmplf func(vault string, args [][]byte, funds []*types.Coin) ([]byte, error)
type ntmplf func(vault string, args [][]byte, funds []*types.Coin) ([]byte, []byte, error)

func wrapTmplFKeccak256(f ptmplf) ntmplf {
	return func(vault string, args [][]byte, funds []*types.Coin) ([]byte, []byte, error) {
		payload, err := f(vault, args, funds)
		if err != nil {
			return nil, nil, err
		}
		return payload, sha3.NewLegacyKeccak256().Sum(payload), nil
	}
}

func wrapTmplFSha256(f ptmplf) ntmplf {
	return func(vault string, args [][]byte, funds []*types.Coin) ([]byte, []byte, error) {
		payload, err := f(vault, args, funds)
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
