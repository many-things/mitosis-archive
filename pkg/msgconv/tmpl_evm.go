package msgconv

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
	"strings"
)

type evmFund struct {
	Token string
	Value string
}

type evmInner struct {
	To    string
	Data  string
	Value string
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
func EvmOp0(_ string, args ...[]byte) ([]byte, error) {
	recipient := string(args[0])
	funds := types.Map(
		strings.Split(string(args[1]), ","),
		func(t string, i int) evmFund {
			ts := strings.Split(t, "0x")

			return evmFund{
				Token: ts[0],
				Value: fmt.Sprintf("0x%s", ts[1]),
			}
		},
	)

	transferABI, err := sigToABI("transfer(address,uint256)")
	if err != nil {
		return nil, errors.Wrap(err, "get transfer abi")
	}

	calldata, err := types.MapErr(
		funds,
		func(t evmFund, i int) (evmInner, error) {
			calldata, err := transferABI.Pack("transfer", recipient, t.Value)
			if err != nil {
				return evmInner{}, err
			}

			return evmInner{To: t.Token, Data: hex.EncodeToString(calldata), Value: "0"}, nil
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "make calldata")
	}

	payload, err := evmPayload{Funds: funds, Inner: calldata}.Pack()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func EvmOp1(_ string, _ ...[]byte) ([]byte, error) {
	return nil, nil
}
