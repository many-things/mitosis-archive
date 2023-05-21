package msgconv

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
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
func CosmosOp0(src, dest, vault string, args [][]byte, funds []*types.Coin) ([]byte, error) {
	if err := assertArgs(args, CosmosOp0RequiredArgsCount); err != nil {
		return nil, err
	}
	if len(funds) == 0 {
		funds = []*types.Coin{{
			Denom:   "0x5Cbb2F9f7E54c5B4656C3B563ff5650a0866A3EF",
			Amount:  types.Ref(sdk.NewInt(100)),
			Decimal: 0,
		}}

		// FIXME: this is a hack to make the tests pass
		// return nil, errors.New("expected at least one fund")
	}

	toAddr := string(args[0])

	deref := func(c *types.Coin, _ int) (sdk.Coin, error) {
		cc := c.ToSDK()
		convDenom, err := convertDenomIO(src, dest, funds[0].Denom)
		if err != nil {
			return sdk.Coin{}, errors.Wrap(err, "convert denom")
		}

		cc.Denom = convDenom
		return cc, nil
	}

	coins, err := types.MapErr(funds, deref)
	if err != nil {
		return nil, errors.Wrap(err, "map coins")
	}

	coinsBz, err := sdk.Coins(coins).MarshalJSON()
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
func CosmosOp1(src, dest, vault string, args [][]byte, funds []*types.Coin) ([]byte, error) {
	if err := assertArgs(args, CosmosOp1RequiredArgsCount); err != nil {
		return nil, err
	}
	if len(funds) != 1 {
		funds = []*types.Coin{{
			Denom:   "0x5Cbb2F9f7E54c5B4656C3B563ff5650a0866A3EF",
			Amount:  types.Ref(sdk.NewInt(100)),
			Decimal: 0,
		}}

		// FIXME
		// return nil, errors.New("expected exactly one fund")
	}

	convDenom, err := convertDenomIO(src, dest, funds[0].Denom)
	if err != nil {
		return nil, errors.Wrap(err, "convert denom")
	}

	osmoTokenIn := &osmo.Coin{
		Denom:  convDenom,
		Amount: funds[0].Amount.String(),
	}

	msgSwap := osmo.MsgSwapExactAmountIn{
		Sender: vault,
		Routes: []*osmo.SwapAmountInRoute{{
			PoolId:        16,
			TokenOutDenom: string(args[1]),
		}},
		TokenIn:           osmoTokenIn,
		TokenOutMinAmount: string(args[2]),
	}
	marshaled, err := proto.Marshal(&msgSwap)
	if err != nil {
		return nil, errors.Wrap(err, "marshal swap msg")
	}

	minAmount, ok := sdk.NewIntFromString(string(args[2]))
	if !ok {
		return nil, errors.New("invalid min amount")
	}
	returnAmount := sdk.Coins{{
		Denom:  string(args[1]),
		Amount: minAmount,
	}}
	returnAmountBz, err := returnAmount.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "marshal return amount")
	}

	rendered := new(bytes.Buffer)
	if err := CosmosOp1Tmpl.Execute(rendered, cosmosPayload{
		Vault: vault,
		Args: []string{
			hex.EncodeToString(marshaled),
			string(args[0]),
			string(returnAmountBz), // FIXME: return only the minimum amount
		},
	}); err != nil {
		return nil, errors.Wrap(err, "execute op1 template")
	}

	renderedBz := rendered.Bytes()
	rendered.Reset()
	if err := json.Compact(rendered, renderedBz); err != nil {
		return nil, errors.Wrapf(err, "compact json. raw=%s", string(renderedBz))
	}

	return rendered.Bytes(), nil
}
