package msgconv

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"strings"
)

var VaultABI = mustParseABI(`[
   {
      "inputs":[
         {
            "components":[
               {
                  "components":[
                     {
                        "internalType":"address",
                        "name":"token",
                        "type":"address"
                     },
                     {
                        "internalType":"uint256",
                        "name":"value",
                        "type":"uint256"
                     }
                  ],
                  "internalType":"struct Vault.ExecuteFund[]",
                  "name":"funds",
                  "type":"tuple[]"
               },
               {
                  "components":[
                     {
                        "internalType":"address",
                        "name":"to",
                        "type":"address"
                     },
                     {
                        "internalType":"uint256",
                        "name":"value",
                        "type":"uint256"
                     },
                     {
                        "internalType":"bytes",
                        "name":"data",
                        "type":"bytes"
                     }
                  ],
                  "internalType":"struct Vault.ExecuteCalldata[]",
                  "name":"inner",
                  "type":"tuple[]"
               }
            ],
            "internalType":"struct Vault.ExecutePayload",
            "name":"_payload",
            "type":"tuple"
         },
         {
            "internalType":"bytes",
            "name":"_signature",
            "type":"bytes"
         }
      ],
      "name":"execute",
      "outputs":[],
      "stateMutability":"nonpayable",
      "type":"function"
   }
]`)

func mustParseABI(abiStr string) abi.ABI {
	i, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		panic(errors.Wrap(err, "must parse given abi"))
	}
	return i
}

func sigToABI(sig string) (abi.ABI, error) {
	sel, err := abi.ParseSelector(sig)
	if err != nil {
		return abi.ABI{}, errors.Wrap(err, "parse selector")
	}

	tmp, err := jsoniter.Marshal([]abi.SelectorMarshaling{sel})
	if err != nil {
		return abi.ABI{}, errors.Wrap(err, "marshal selector to abi")
	}

	i, err := abi.JSON(bytes.NewReader(tmp))
	if err != nil {
		return abi.ABI{}, errors.Wrap(err, "parse abi")
	}

	return i, nil
}
