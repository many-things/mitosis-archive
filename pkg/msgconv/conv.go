package msgconv

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
)

var chainReg = map[string]types.ChainType{
	"osmo-test-5": types.ChainType_TypeCosmos,
	"evm-5":       types.ChainType_TypeEvm,
}

var evmEncoderABI abi.ABI

func init() {
	encodeABI, err := sigToABI("test(string)")
	if err != nil {
		panic(err.Error())
	}
	evmEncoderABI = encodeABI
}

func decodeEvmArgs(args [][]byte) ([][]byte, error) {
	convArgs, err := types.MapErr(
		args,
		func(t []byte, i int) ([]byte, error) {
			unpacked, err := evmEncoderABI.Methods["test"].Inputs.Unpack(t)
			if err != nil {
				return nil, errors.Wrap(err, "failed to unpack ABI")
			}

			return []byte(unpacked[0].(string)), nil
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert ABI")
	}
	return convArgs, nil
}

func ToMsgs(src, dest string, vault string, opID uint64, opArgs [][]byte, funds []*types.Coin) (payload []byte, bytesToSign []byte, err error) {
	srcChainType, ok := chainReg[src]
	if !ok {
		return nil, nil, errors.New("unknown src chain type")
	}
	if srcChainType == types.ChainType_TypeEvm {
		decoded, err := decodeEvmArgs(opArgs)
		if err != nil {
			return nil, nil, err
		}
		opArgs = decoded
	}

	destChainType, ok := chainReg[dest]
	if !ok {
		return nil, nil, errors.New("unknown dest chain type")
	}

	return tmpl[destChainType][opID](dest, vault, opArgs, funds)
}
