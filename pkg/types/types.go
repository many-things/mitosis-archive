package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (c Coin) ToSDK() sdk.Coin {
	return sdk.Coin{
		Denom:  c.Denom,
		Amount: sdk.NewIntFromBigInt(c.Amount.BigInt()),
	}
}
