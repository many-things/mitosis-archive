package msgconv

import (
	"github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
)

func ToMsgs(chain string, vault string, opID uint64, opArgs [][]byte, funds []*types.Coin) (payload []byte, bytesToSign []byte, err error) {
	chainType := types.ChainType_TypeUnspecified
	switch chain {
	case "osmo-test-5":
		chainType = types.ChainType_TypeCosmos
	case "evm-5":
		chainType = types.ChainType_TypeEvm
	}
	if chainType == types.ChainType_TypeUnspecified {
		return nil, nil, errors.Errorf("unknown chain type: %s", chain)
	}

	return tmpl[chainType][opID](vault, opArgs, funds)
}
