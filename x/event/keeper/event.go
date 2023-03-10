package keeper

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/keeper/store"
	"github.com/many-things/mitosis/x/event/types"
)

type EventKeeper interface {
	StoreEvent(ctx sdk.Context, incoming []*types.IncomingEvent, outgoing []*types.OutgoingEvent) error

	GetIncomingEvent(ctx sdk.Context, txHash string, evtIndex uint64) (*types.IncomingEvent, error)
	ListIncomingEvent(ctx sdk.Context) ([]*types.IncomingEvent, error)

	GetOutgoingEvent(ctx sdk.Context, txHash string) (*types.OutgoingEvent, error)
	ListOutgoingEvent(ctx sdk.Context) ([]*types.OutgoingEvent, error)
}

type eventKeeper struct {
	storeKey storetypes.StoreKey
}

func newEventKeeper(storeKey storetypes.StoreKey) EventKeeper {
	return eventKeeper{storeKey}
}

func (k eventKeeper) StoreEvent(ctx sdk.Context, incoming []*types.IncomingEvent, outgoing []*types.OutgoingEvent) error {
	inStore := store.NewIncomingEventRepo(ctx, k.storeKey)
	if err := inStore.Store(incoming); err != nil {
		return err
	}

	outStore := store.NewOutgoingEventRepo(ctx, k.storeKey)
	if err := outStore.Store(outgoing); err != nil {
		return err
	}

	return nil
}

func (k eventKeeper) GetIncomingEvent(ctx sdk.Context, txHash string, evtIndex uint64) (*types.IncomingEvent, error) {
	inStore := store.NewIncomingEventRepo(ctx, k.storeKey)

	return inStore.Get(txHash, evtIndex)
}

func (k eventKeeper) ListIncomingEvent(ctx sdk.Context) ([]*types.IncomingEvent, error) {
	inStore := store.NewIncomingEventRepo(ctx, k.storeKey)

	return inStore.List()
}

func (k eventKeeper) GetOutgoingEvent(ctx sdk.Context, txHash string) (*types.OutgoingEvent, error) {
	outStore := store.NewOutgoingEventRepo(ctx, k.storeKey)

	return outStore.Get(txHash)
}

func (k eventKeeper) ListOutgoingEvent(ctx sdk.Context) ([]*types.OutgoingEvent, error) {
	outStore := store.NewOutgoingEventRepo(ctx, k.storeKey)

	return outStore.List()
}
