package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
)

var _ EventKeeper = Keeper{}

type EventKeeper interface {
	StoreEvent(ctx sdk.Context, msg *types.MsgVoteEvent) error

	GetIncomingEvent(ctx sdk.Context, txHash string, evtIndex uint64) (*types.IncomingEvent, error)
	ListIncomingEvent(ctx sdk.Context, txHash string) ([]*types.IncomingEvent, error)

	GetOutgoingEvent(ctx sdk.Context, txHash string) (*types.OutgoingEvent, error)
	ListOutgoingEvent(ctx sdk.Context) ([]*types.OutgoingEvent, error)
}

func (k Keeper) StoreEvent(ctx sdk.Context, msg *types.MsgVoteEvent) error {
	inStore := newIncomingEventStore(k.GetIncomingEventsStore(ctx))
	if err := inStore.Store(msg.GetIncoming()); err != nil {
		return err
	}

	outStore := newOutgoingEventStore(k.GetOutgoingEventsStore(ctx))
	if err := outStore.Store(msg.GetOutgoing()); err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetIncomingEvent(ctx sdk.Context, txHash string, evtIndex uint64) (*types.IncomingEvent, error) {
	inStore := newIncomingEventStore(k.GetIncomingEventsStore(ctx))
	return inStore.Get(txHash, evtIndex)
}

func (k Keeper) ListIncomingEvent(ctx sdk.Context, txHash string) ([]*types.IncomingEvent, error) {
	inStore := newIncomingEventStore(k.GetIncomingEventsStore(ctx))
	return inStore.List(txHash)
}

func (k Keeper) GetOutgoingEvent(ctx sdk.Context, txHash string) (*types.OutgoingEvent, error) {
	outStore := newOutgoingEventStore(k.GetOutgoingEventsStore(ctx))
	return outStore.Get(txHash)
}

func (k Keeper) ListOutgoingEvent(ctx sdk.Context) ([]*types.OutgoingEvent, error) {
	outStore := newOutgoingEventStore(k.GetOutgoingEventsStore(ctx))
	return outStore.List()
}
