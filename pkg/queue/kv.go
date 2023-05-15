package queue

import (
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
)

var (
	kvKeyFirstItem = []byte{0x01}
	kvKeyLastItem  = []byte{0x02}
	kvPrefixItems  = []byte{0x03}
)

type kvq[T Message] struct {
	root        store.KVStore
	items       store.KVStore
	constructor func() T
}

func NewKVQueue[T Message](root store.KVStore, constructor func() T) Queue[T] {
	items := prefix.NewStore(root, kvPrefixItems)

	return kvq[T]{root, items, constructor}
}

func (k kvq[T]) getFirstItem() uint64 {
	if bz := k.root.Get(kvKeyFirstItem); bz != nil {
		return sdk.BigEndianToUint64(bz)
	}

	return 0
}

func (k kvq[T]) setFirstItem(firstItem uint64) {
	k.root.Set(
		kvKeyFirstItem,
		sdk.Uint64ToBigEndian(firstItem),
	)
}

func (k kvq[T]) getLastItem() uint64 {
	bz := k.root.Get(kvKeyLastItem)
	if bz != nil {
		return sdk.BigEndianToUint64(bz)
	}

	return 0
}

func (k kvq[T]) setLastItem(lastItem uint64) {
	k.root.Set(
		kvKeyLastItem,
		sdk.Uint64ToBigEndian(lastItem),
	)
}

// Size returns the number of items in the queue.
func (k kvq[T]) Size() uint64 {
	lastItem := k.getLastItem()
	firstItem := k.getFirstItem()

	return lastItem - firstItem
}

// LastIndex returns the last item's index of the queue.
func (k kvq[T]) LastIndex() uint64 {
	return k.getLastItem()
}

// Get returns the item of specific id
func (k kvq[T]) Get(i uint64) (T, error) {
	m := k.constructor()

	// TODO: query historical polls
	//if i < k.getFirstItem() || k.getLastItem() <= i {
	//	return m, errors.New("index out of range")
	//}

	bz := k.items.Get(sdk.Uint64ToBigEndian(i))
	if bz == nil {
		return m, errors.Errorf("item not found in index %d", i)
	}

	if err := m.Unmarshal(bz); err != nil {
		return m, errors.Wrap(err, "unmarshal")
	}
	return m, nil
}

// Range iterates over the queue and calls the callback for each item.
func (k kvq[T]) Range(amount *uint64, f func(T, uint64) error) error {
	lastItem := k.getLastItem() // nolint: ifshort
	firstItem := k.getFirstItem()

	if lastItem == firstItem {
		return errors.New("empty queue")
	}

	limit := uint64(query.MaxLimit)
	if amount != nil {
		limit = *amount
	}

	var (
		want     = min(lastItem-firstItem, limit)
		queryReq = &query.PageRequest{
			Key:     sdk.Uint64ToBigEndian(firstItem),
			Limit:   want,
			Reverse: false,
		}
	)

	_, err := query.Paginate(
		k.items,
		queryReq,
		func(key []byte, value []byte) error {
			m := k.constructor()
			if err := m.Unmarshal(value); err != nil {
				return err
			}
			return f(m, sdk.BigEndianToUint64(key))
		},
	)
	return err
}

// Paginate iterates over the queue and calls the callback for each item.
func (k kvq[T]) Paginate(req *query.PageRequest, f func(T, uint64) error) (*query.PageResponse, error) {
	lastItem := k.getLastItem()
	firstItem := k.getFirstItem()
	if lastItem == firstItem {
		return nil, errors.New("empty queue")
	}

	if req.Key == nil {
		if req.Reverse {
			req.Key = sdk.Uint64ToBigEndian(lastItem)
		} else {
			req.Key = sdk.Uint64ToBigEndian(firstItem)
		}
	} else {
		i := sdk.BigEndianToUint64(req.Key)
		if i < firstItem || lastItem <= i {
			return nil, errors.New("key out of range")
		}
	}

	if req.Limit == 0 {
		req.Limit = lastItem - firstItem
	} else {
		req.Limit = min(lastItem-firstItem, req.Limit)
	}

	resp, err := query.Paginate(
		k.items,
		req,
		func(key []byte, value []byte) error {
			m := k.constructor()
			if err := m.Unmarshal(value); err != nil {
				return err
			}
			return f(m, sdk.BigEndianToUint64(key))
		},
	)
	if err != nil {
		return nil, err
	}

	next := sdk.BigEndianToUint64(resp.NextKey)
	if next < firstItem || lastItem <= next {
		resp.NextKey = nil
	}

	return resp, nil
}

// MsgConstructor returns the constructor of the message type.
func (k kvq[T]) MsgConstructor() func() T {
	return k.constructor
}

// Produce pushes the given messages to the queue.
func (k kvq[T]) Produce(msgs ...T) ([]uint64, error) {
	lastItem := k.getLastItem()
	for i, msg := range msgs {
		bz, err := msg.Marshal()
		if err != nil {
			return nil, err
		}

		k.items.Set(sdk.Uint64ToBigEndian(lastItem+uint64(i)), bz)
	}
	k.setLastItem(lastItem + uint64(len(msgs))) // -> last item

	return mitotypes.Map(
		make([]byte, len(msgs)),
		func(_ byte, i int) uint64 { return lastItem + uint64(i) },
	), nil
}

// Update updates the item of specific id.
func (k kvq[T]) Update(i uint64, msg T) error {
	if i < k.getFirstItem() && k.getLastItem() < i {
		return errors.New("index out of range")
	}

	bz, err := msg.Marshal()
	if err != nil {
		return errors.Wrap(err, "marshal")
	}

	k.items.Set(sdk.Uint64ToBigEndian(i), bz)
	return nil
}

// Consume pops the given amount of items from the queue.
func (k kvq[T]) Consume(amount uint64) ([]mitotypes.KV[uint64, T], error) {
	lastItem := k.getLastItem() // nolint: ifshort
	firstItem := k.getFirstItem()
	if lastItem == firstItem {
		return nil, errors.New("empty queue")
	}

	var (
		ms       []mitotypes.KV[uint64, T]
		want     = min(lastItem-firstItem, amount)
		queryReq = &query.PageRequest{
			Key:     sdk.Uint64ToBigEndian(firstItem),
			Limit:   want,
			Reverse: false,
		}
	)

	_, err := query.Paginate(
		k.items, queryReq,
		func(key []byte, value []byte) error {
			m := k.constructor()
			if err := m.Unmarshal(value); err != nil {
				return err
			}
			ms = append(ms, mitotypes.NewKV(sdk.BigEndianToUint64(key), m))
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	k.setFirstItem(firstItem + uint64(len(ms)))
	return ms, nil
}

// ConsumeUntil pops the items from the queue until the given condition is met.
func (k kvq[T]) ConsumeUntil(f func(T, uint64) (bool, error)) ([]mitotypes.KV[uint64, T], error) {
	var (
		ms        []mitotypes.KV[uint64, T]
		firstItem = k.getFirstItem()
		lastItem  = k.getLastItem()
	)

	iter := k.items.Iterator(
		sdk.Uint64ToBigEndian(firstItem),
		sdk.Uint64ToBigEndian(lastItem),
	)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		key, value := iter.Key(), iter.Value()
		i := sdk.BigEndianToUint64(key)

		m := k.constructor()
		if err := m.Unmarshal(value); err != nil {
			return nil, err
		}
		ms = append(ms, mitotypes.NewKV(i, m))

		if ok, err := f(m, i); err != nil {
			return nil, err
		} else if !ok {
			break
		}

		if err := iter.Error(); err != nil {
			return nil, err
		}
	}

	k.setFirstItem(firstItem + uint64(len(ms)))
	return ms, nil
}

// ImportGenesis imports the queue's genesis state.
func (k kvq[T]) ImportGenesis(g GenesisState[T]) error {
	k.setFirstItem(g.FirstIndex)
	k.setLastItem(g.LastIndex)

	for _, item := range g.Items {
		i, m := item.Key, item.Value

		bz, err := m.Marshal()
		if err != nil {
			return errors.Wrap(err, "marshal")
		}

		k.items.Set(sdk.Uint64ToBigEndian(i), bz)
	}

	return nil
}

// ExportGenesis exports the queue's genesis state.
func (k kvq[T]) ExportGenesis() (GenesisState[T], error) {
	g := GenesisState[T]{
		FirstIndex: k.getFirstItem(),
		LastIndex:  k.getLastItem(),
	}

	_, err := query.Paginate(
		k.items,
		&query.PageRequest{Limit: query.MaxLimit},
		func(key []byte, value []byte) error {
			m := k.constructor()
			if err := m.Unmarshal(value); err != nil {
				return err
			}

			g.Items = append(
				g.Items,
				mitotypes.NewKV(sdk.BigEndianToUint64(key), m),
			)
			return nil
		},
	)
	if err != nil {
		return GenesisState[T]{}, err
	}

	return g, nil
}
