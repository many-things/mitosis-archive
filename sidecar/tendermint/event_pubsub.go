package tendermint

import abci "github.com/tendermint/tendermint/abci/types"

type TmEvent struct {
	BlockHeight int64
	abci.Event
}
