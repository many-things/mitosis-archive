package cosmos

import (
	sdkerrutils "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	txconvtypes "github.com/many-things/mitosis/pkg/txconv/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

var (
	DefaultOsmosisFee = sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(250)))
	DefaultOsmosisOps = []mitotypes.KV[uint64, Operation]{
		mitotypes.NewKV(uint64(0), Operation{
			Msgs: func(args ...[]byte) ([]sdk.Msg, error) {
				if len(args) < 3 {
					return nil, sdkerrutils.Wrap(sdkerrors.ErrPanic, "invalid args")
				}

				amount, err := sdk.ParseCoinsNormalized(string(args[2]))
				if err != nil {
					return nil, err
				}

				return []sdk.Msg{
					&banktypes.MsgSend{
						FromAddress: string(args[0]),
						ToAddress:   string(args[1]),
						Amount:      amount,
					},
				}, nil
			},
			Gas: 0,
		}),
	}
)

type Operation struct {
	Msgs func(args ...[]byte) ([]sdk.Msg, error)
	Gas  uint64
}

type ChainInfoConfig struct {
	fee sdk.Coins
	ops []mitotypes.KV[uint64, Operation] // [opID, gas]
}

type ChainInfoOption func(*ChainInfoConfig)

func WithFee(fee sdk.Coins) ChainInfoOption {
	return func(c *ChainInfoConfig) {
		c.fee = fee
	}
}

func AddOperation(id uint64, op Operation) ChainInfoOption {
	return func(c *ChainInfoConfig) {
		if op := mitotypes.FindKV(c.ops, func(k uint64, _ Operation, _ int) bool { return k == id }); op != nil {
			panic("operation already exists")
		}

		c.ops = append(c.ops, mitotypes.NewKV(id, op))
	}
}

func MakeChainInfo(chainID, chainName string, encoder client.TxConfig, opts ...ChainInfoOption) mitotypes.KV[string, txconvtypes.ChainInfo] {
	config := ChainInfoConfig{
		fee: DefaultOsmosisFee,
		ops: DefaultOsmosisOps,
	}
	for _, opt := range opts {
		opt(&config)
	}

	chainInfo := txconvtypes.ChainInfo{
		Type:      txconvtypes.ChainTypeCosmos,
		ChainID:   chainID,
		ChainName: chainName,

		TxConv: convert(chainID, encoder, config),
	}

	return mitotypes.NewKV(chainID, chainInfo)
}
