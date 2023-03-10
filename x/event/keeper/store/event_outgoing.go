package store

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
)

type OutgoingEventRepo interface {
	Store(events []*types.OutgoingEvent) error
	Get(txHash string) (*types.OutgoingEvent, error)
	List() ([]*types.OutgoingEvent, error)
}

type outgoingEventRepo struct{ store.KVStore }

func NewOutgoingEventRepo(ctx sdk.Context, key storetypes.StoreKey) OutgoingEventRepo {
	return outgoingEventRepo{prefix.NewStore(
		ctx.KVStore(key),
		types.GetOutgoingEventPrefix(uint64(ctx.BlockHeight())),
	)}
}

func (s outgoingEventRepo) buildKey(event *types.OutgoingEvent) []byte {
	txHashBz, err := hex.DecodeString(event.GetTxHash())
	if err != nil {
		panic(err.Error())
	}
	return txHashBz
}

func (s outgoingEventRepo) Store(events []*types.OutgoingEvent) error {
	for _, evt := range events {
		key := s.buildKey(evt)
		if !s.KVStore.Has(key) {
			bz, err := evt.Marshal()
			if err != nil {
				panic(err.Error())
			}
			s.KVStore.Set(key, bz)
		}
	}
	return nil
}

func (s outgoingEventRepo) Get(txHash string) (*types.OutgoingEvent, error) {
	key := s.buildKey(&types.OutgoingEvent{TxHash: txHash})
	bz := s.KVStore.Get(key)

	event := new(types.OutgoingEvent)
	if err := event.Unmarshal(bz); err != nil {
		return nil, err
	}
	return event, nil
}

func (s outgoingEventRepo) List() ([]*types.OutgoingEvent, error) {
	iter := s.KVStore.Iterator(nil, nil)
	defer iter.Close()

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
