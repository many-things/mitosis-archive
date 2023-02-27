package keeper

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/many-things/mitosis/x/event/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

const (
	PrefixVoteIncomingEvent = "vie"
	PrefixVoteOutgoingEvent = "vio"
)

func must[expected any](e expected, err error) expected {
	if err != nil {
		panic(err.Error())
	}
	return e
}

func unwrap[expected any](f func() (expected, error)) expected {
	return must(f())
}
func unwrap1[expected any, arg1 any](f func(arg1) (expected, error), a1 arg1) expected {
	return must(f(a1))
}
func unwrap2[expected any, arg1 any, arg2 any](f func(arg1, arg2) (expected, error), a1 arg1, a2 arg2) expected {
	return must(f(a1, a2))
}

func (k Keeper) handleIncomingEvent(ctx sdk.Context, evts []*types.IncomingEvent) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(PrefixVoteIncomingEvent))
	height := uint64(ctx.BlockHeight())

	for _, evt := range evts {
		// ensure txHash already validated before execute this logic
		key := bytes.Join([][]byte{
			sdk.Uint64ToBigEndian(height),
			unwrap1(hex.DecodeString, evt.GetTxHash()),
			sdk.Uint64ToBigEndian(evt.EventIndex),
		}, []byte(":"))

		store.Set(key, unwrap(evt.Marshal))
	}

	return nil
}

func (k Keeper) handleOutgoingEvent(ctx sdk.Context, evts []*types.OutgoingEvent) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(PrefixVoteOutgoingEvent))
	height := uint64(ctx.BlockHeight())

	for _, evt := range evts {
		// ensure txHash already validated before execute this logic
		key := bytes.Join([][]byte{
			sdk.Uint64ToBigEndian(height),
			unwrap1(hex.DecodeString, evt.GetTxHash()),
		}, []byte(":"))

		store.Set(key, unwrap(evt.Marshal))
	}

	return nil
}

func (k Keeper) VoteEvent(ctx sdk.Context, msg *types.MsgVoteEvent) error {
	if err := k.handleIncomingEvent(ctx, msg.GetIncoming()); err != nil {
		return err
	}

	if err := k.handleOutgoingEvent(ctx, msg.GetOutgoing()); err != nil {
		return err
	}

	return nil
}
