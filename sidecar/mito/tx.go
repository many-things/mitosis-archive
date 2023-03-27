package mito

import sdk "github.com/cosmos/cosmos-sdk/types"

type EventTxMgr interface {
	BroadcastTx(msg sdk.Msg) error
}

func (m eventMgr) BroadcastTx(msg sdk.Msg) error {
	return m.wallet.BroadcastMsg(msg)
}
