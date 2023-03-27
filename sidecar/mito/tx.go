package mito

import sdk "github.com/cosmos/cosmos-sdk/types"

type MitoEventTxMgr interface {
	BroadcastTx(msg sdk.Msg) error
}

func (m mitoEventMgr) BroadcastTx(msg sdk.Msg) error {
	return m.wallet.BroadcastMsg(msg)
}
