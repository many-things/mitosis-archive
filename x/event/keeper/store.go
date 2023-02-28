package keeper

import (
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	PrefixVoteIncomingEvent = "vie"
	PrefixVoteOutgoingEvent = "vio"
)

var _ StoreProvider = Keeper{}

type StoreProvider interface {
	// GetIncomingEventsStore returns prefixed store with root of events
	GetIncomingEventsStore(ctx sdk.Context) store.KVStore

	// GetOutgoingEventsStore returns prefixed store with root of events
	GetOutgoingEventsStore(ctx sdk.Context) store.KVStore
}

func (k Keeper) GetIncomingEventsStore(ctx sdk.Context) store.KVStore {
	return prefix.NewStore(
		ctx.KVStore(k.storeKey),
		joinBytes(
			[]byte(":"),
			[]byte(PrefixVoteIncomingEvent),
			sdk.Uint64ToBigEndian(uint64(ctx.BlockHeight())),
		),
	)
}

func (k Keeper) GetOutgoingEventsStore(ctx sdk.Context) store.KVStore {
	return prefix.NewStore(
		ctx.KVStore(k.storeKey),
		joinBytes(
			[]byte(":"),
			[]byte(PrefixVoteOutgoingEvent),
			sdk.Uint64ToBigEndian(uint64(ctx.BlockHeight())),
		),
	)
}
