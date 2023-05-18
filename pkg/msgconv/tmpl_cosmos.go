package msgconv

import (
	"bytes"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"text/template"
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
// 1 - amounts to send (formatted like `1uosmo,2uatom`)
func CosmosOp0(vault string, args ...[]byte) ([]byte, error) {
	toAddr := string(args[0])
	amount := string(args[1])

	coins, err := sdk.ParseCoinsNormalized(amount)
	if err != nil {
		return nil, errors.Wrap(err, "parse coins")
	}

	coinsBz, err := coins.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "marshal coins to json")
	}

	rendered := bytes.NewBuffer([]byte{})
	err = CosmosOp0Tmpl.Execute(rendered, cosmosPayload{
		Vault: vault,
		Args: []string{
			toAddr,
			string(coinsBz),
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

const CosmosOp1Tmpl = `[
	{
		"stargate": {
			"type_url": "/osmosis.poolmanager.v1beta1.MsgSwapExactAmountIn",
			"value": ""
		}
	}
]`

func CosmosOp1(_ string, _ ...[]byte) ([]byte, error) {
	// TODO: define me
	return nil, nil
}
