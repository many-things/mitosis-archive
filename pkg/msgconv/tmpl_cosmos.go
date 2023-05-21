package msgconv

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"text/template"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/msgconv/osmo"
	"github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type cosmosPayload struct {
	Vault string
	Args  []string
}

func MustParse(name string, text string) *template.Template {
	t, err := template.New(name).Parse(text)
	if err != nil {
		panic(errors.Wrap(err, "parse template"))
	}
	return t
}

const CosmosOp0RequiredArgsCount = 1
const CosmosOp1RequiredArgsCount = 3

var CosmosOp0Tmpl = MustParse("cosmos-op-0", `[
	{
		"bank": {
			"send": {
				"to_address": "{{index .Args 0}}",
				"amount": {{index .Args 1}}
			}
		}
	}
]`)

// CosmosOp0 has the following arguments:
// 0 - recipient address
func CosmosOp0(chain, vault string, args [][]byte, funds []*types.Coin) ([]byte, error) {
	if err := assertArgs(args, CosmosOp0RequiredArgsCount); err != nil {
		return nil, err
	}

	toAddr := string(args[0])

	deref := func(c *types.Coin, _ int) sdk.Coin {
		cc := c.ToSDK()
		cc.Denom = AssetMappingReverse[cc.Denom][chain]
		return cc
	}
	fundsBz, err := sdk.Coins(types.Map(funds, deref)).MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "marshal coins to json")
	}

	rendered := bytes.NewBuffer([]byte{})
	err = CosmosOp0Tmpl.Execute(rendered, cosmosPayload{
		Vault: vault,
		Args: []string{
			toAddr,
			string(fundsBz),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "execute template")
	}

	renderedBz := rendered.Bytes()
	rendered.Reset()
	if err := json.Compact(rendered, renderedBz); err != nil {
		return nil, errors.Wrap(err, "compact json")
	}

	return rendered.Bytes(), nil
}

var CosmosOp1Tmpl = MustParse("cosmos-op-1", `[
	{
		"stargate": {
			"type_url": "/osmosis.poolmanager.v1beta1.MsgSwapExactAmountIn",
			"value": "{{index .Args 0}}"
		}
	},
	{
		"bank": {
			"send": {
				"to_address": "{{index .Args 1}}",
				"amount": {{index .Args 2}}
			}
		}
	}
]`)

// CosmosOp1 has the following arguments:
// 0 - recipient address
// 1 - swap target denom
// 2 - swap minimum amount
func CosmosOp1(chain, vault string, args [][]byte, funds []*types.Coin) ([]byte, error) {
	if err := assertArgs(args, CosmosOp1RequiredArgsCount); err != nil {
		return nil, err
	} else if len(funds) == 0 {
		return nil, fmt.Errorf("CosmoOp1: must request with funds")
	}

	msgSwap := osmo.MsgSwapExactAmountIn{
		Sender: vault,
		Routes: []*osmo.SwapAmountInRoute{{
			PoolId:        16, // FIXME: hardcoded
			TokenOutDenom: string(args[1]),
		}},
		TokenIn: &osmo.Coin{
			Denom:  AssetMappingReverse[funds[0].Denom][chain],
			Amount: funds[0].Amount.String(),
		},
		TokenOutMinAmount: string(args[2]),
	}
	marshaled, err := proto.Marshal(&msgSwap)
	if err != nil {
		return nil, errors.Wrap(err, "marshal swap msg")
	}

	rendered := new(bytes.Buffer)
	if err := CosmosOp1Tmpl.Execute(rendered, cosmosPayload{
		Vault: vault,
		Args: []string{
			hex.EncodeToString(marshaled),
			string(args[0]),
			string(args[2]), // FIXME: return only the minimum amount
		},
	}); err != nil {
		return nil, errors.Wrap(err, "execute op1 template")
	}

	renderedBz := rendered.Bytes()
	rendered.Reset()
	if err := json.Compact(rendered, renderedBz); err != nil {
		return nil, errors.Wrap(err, "compact json")
	}

	return rendered.Bytes(), nil
}
