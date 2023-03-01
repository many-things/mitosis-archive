package store

import (
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/utils"
)

const (
	PrefixVoteIncomingEvent = "vie"
	PrefixVoteOutgoingEvent = "vio"
	PrefixValidatorProxy    = "vp"
)

func IncomingEventStore(ctx sdk.Context, key storetypes.StoreKey) store.KVStore {
	return prefix.NewStore(
		ctx.KVStore(key),
		utils.JoinBytes(
			[]byte(":"),
			[]byte(PrefixVoteIncomingEvent),
			sdk.Uint64ToBigEndian(uint64(ctx.BlockHeight())),
		),
	)
}

func OutgoingEventStore(ctx sdk.Context, key storetypes.StoreKey) store.KVStore {
	return prefix.NewStore(
		ctx.KVStore(key),
		utils.JoinBytes(
			[]byte(":"),
			[]byte(PrefixVoteOutgoingEvent),
			sdk.Uint64ToBigEndian(uint64(ctx.BlockHeight())),
		),
	)
}

func ValidatorProxyStore(ctx sdk.Context, key storetypes.StoreKey) store.KVStore {
	return prefix.NewStore(
		ctx.KVStore(key),
		utils.JoinBytes([]byte(":"), []byte(PrefixValidatorProxy)),
	)
}
