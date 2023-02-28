package keeper

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
)

type EventStoreIncoming struct{ store.KVStore }

func newIncomingEventStore(store store.KVStore) EventStoreIncoming {
	return EventStoreIncoming{store}
}

func (s EventStoreIncoming) buildKey(event *types.IncomingEvent) []byte {
	return joinBytes(
		[]byte(":"),
		unwrap1(hex.DecodeString, event.GetTxHash()),
		sdk.Uint64ToBigEndian(event.GetEventIndex()),
	)
}

func (s EventStoreIncoming) Store(events []*types.IncomingEvent) error {
	for _, evt := range events {
		// ensure txHash already validated before execute this logic
		key := s.buildKey(evt)
		if !s.KVStore.Has(key) {
			s.KVStore.Set(key, unwrap(evt.Marshal))
		}
	}
	return nil
}

func (s EventStoreIncoming) Get(txHash string, evtIndex uint64) (*types.IncomingEvent, error) {
	key := s.buildKey(&types.IncomingEvent{TxHash: txHash, EventIndex: evtIndex})
	bz := s.KVStore.Get(key)

	event := new(types.IncomingEvent)
	if err := event.Unmarshal(bz); err != nil {
		return nil, err
	}
	return event, nil
}

func (s EventStoreIncoming) List(txHash string) ([]*types.IncomingEvent, error) {
	iter := prefix.
		NewStore(s.KVStore, unwrap1(hex.DecodeString, txHash)).
		Iterator(nil, nil)

	var events []*types.IncomingEvent
	for ; iter.Valid(); iter.Next() {
		bz := iter.Value()

		event := new(types.IncomingEvent)
		if err := event.Unmarshal(bz); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := iter.Error(); err != nil {
		return nil, err
	}

	return events, nil
}

type EventStoreOutgoing struct{ store.KVStore }

func newOutgoingEventStore(store store.KVStore) EventStoreOutgoing {
	return EventStoreOutgoing{store}
}

func (s EventStoreOutgoing) buildKey(event *types.OutgoingEvent) []byte {
	return joinBytes(
		[]byte(":"),
		unwrap1(hex.DecodeString, event.GetTxHash()),
	)
}

func (s EventStoreOutgoing) Store(events []*types.OutgoingEvent) error {
	for _, evt := range events {
		key := s.buildKey(evt)
		if !s.KVStore.Has(key) {
			s.KVStore.Set(key, unwrap(evt.Marshal))
		}
	}
	return nil
}

func (s EventStoreOutgoing) Get(txHash string) (*types.OutgoingEvent, error) {
	key := s.buildKey(&types.OutgoingEvent{TxHash: txHash})
	bz := s.KVStore.Get(key)

	event := new(types.OutgoingEvent)
	if err := event.Unmarshal(bz); err != nil {
		return nil, err
	}
	return event, nil
}

func (s EventStoreOutgoing) List() ([]*types.OutgoingEvent, error) {
	iter := s.KVStore.Iterator(nil, nil)

	var events []*types.OutgoingEvent
	for ; iter.Valid(); iter.Next() {
		bz := iter.Value()

		event := new(types.OutgoingEvent)
		if err := event.Unmarshal(bz); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := iter.Error(); err != nil {
		return nil, nil
	}

	return events, nil
}
