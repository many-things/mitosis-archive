package store

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/utils"
	"github.com/many-things/mitosis/x/event/types"
)

type IncomingEventRepo interface {
	Store(events []*types.IncomingEvent) error
	Get(txHash string, evtIndex uint64) (*types.IncomingEvent, error)
	List() ([]*types.IncomingEvent, error)
}

type incomingEventRepo struct{ store.KVStore }

func NewIncomingEventRepo(ctx sdk.Context, key storetypes.StoreKey) IncomingEventRepo {
	return incomingEventRepo{prefix.NewStore(
		ctx.KVStore(key),
		types.GetIncomingEventPrefix(uint64(ctx.BlockHeight())),
	)}
}

func (s incomingEventRepo) buildKey(event *types.IncomingEvent) []byte {
	return utils.JoinBytes(
		[]byte(":"),
		utils.Unwrap1(hex.DecodeString, event.GetTxHash()),
		sdk.Uint64ToBigEndian(event.GetEventIndex()),
	)
}

func (s incomingEventRepo) Store(events []*types.IncomingEvent) error {
	for _, evt := range events {
		// ensure txHash already validated before execute this logic
		key := s.buildKey(evt)
		if !s.KVStore.Has(key) {
			s.KVStore.Set(key, utils.Unwrap(evt.Marshal))
		}
	}
	return nil
}

func (s incomingEventRepo) Get(txHash string, evtIndex uint64) (*types.IncomingEvent, error) {
	key := s.buildKey(&types.IncomingEvent{TxHash: txHash, EventIndex: evtIndex})
	bz := s.KVStore.Get(key)

	event := new(types.IncomingEvent)
	if err := event.Unmarshal(bz); err != nil {
		return nil, err
	}
	return event, nil
}

func (s incomingEventRepo) List() ([]*types.IncomingEvent, error) {
	iter := s.KVStore.Iterator(nil, nil)

	var events []*types.IncomingEvent
	for ; iter.Valid(); iter.Next() {
		event := new(types.IncomingEvent)
		utils.Must(nil, event.Unmarshal(iter.Value()))
		events = append(events, event)
	}
	if err := iter.Error(); err != nil {
		return nil, err
	}

	return events, nil
}
