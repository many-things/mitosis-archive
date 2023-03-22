package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (t *Tally) Passed(threshold sdk.Dec) bool {
	return sdk.NewDecFromInt(*t.Confirmed).QuoInt(*t.TotalPower).GT(threshold)
}
