package msgconv

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
	"math/big"
)

const EvmOp0RequiredArgsCount = 1
const EvmOp1RequiredArgsCount = 1

type evmFund struct {
	Token common.Address
	Value *big.Int
}

type evmInner struct {
	To    common.Address
	Data  []byte
	Value *big.Int
}

type evmPayload struct {
	Funds []evmFund
	Inner []evmInner
}

func (p evmPayload) Pack() ([]byte, error) {
	packed, err := abi.Arguments{VaultABI.Methods["execute"].Inputs[0]}.Pack(p)
	if err != nil {
		return nil, errors.Wrap(err, "pack evm payload")
	}
	return packed, nil
}

// EvmOp0 has the following arguments:
// 0 - recipient address
// 1 - funds (formatted like `10xdeadbeefdeadbeef,20xdeadbeefdeadbeef`)
func EvmOp0(_ string, args [][]byte, funds []*types.Coin) ([]byte, error) {
	if err := assertArgs(args, EvmOp0RequiredArgsCount); err != nil {
		return nil, err
	}

	recipient := common.HexToAddress(string(args[0]))

	conv := func(c *types.Coin, _ int) evmFund {
		return evmFund{
			Token: common.HexToAddress(c.Denom),
			Value: c.Amount.BigInt(),
		}
	}
	coins := types.Map(funds, conv)

	transferABI, err := sigToABI("transfer(address,uint256)")
	if err != nil {
		return nil, errors.Wrap(err, "get transfer abi")
	}

	inners, err := types.MapErr(
		coins,
		func(t evmFund, i int) (evmInner, error) {
			transferCalldata, err := transferABI.Pack("transfer", recipient, t.Value)
			if err != nil {
				return evmInner{}, err
			}

			inner := evmInner{
				To:    t.Token,
				Data:  transferCalldata,
				Value: big.NewInt(0),
			}

			return inner, nil
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "make calldata")
	}

	payload, err := evmPayload{Funds: coins, Inner: inners}.Pack()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func EvmOp1(_ string, args [][]byte, _ []*types.Coin) ([]byte, error) {
	if err := assertArgs(args, EvmOp1RequiredArgsCount); err != nil {
		return nil, err
	}

	return nil, nil
}
