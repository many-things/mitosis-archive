package hook

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, keeper types.BaseKeeper) {

}

func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, keeper types.BaseKeeper) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}
