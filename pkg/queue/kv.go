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
	if bz := k.root.Get(kvKeyFirstItem); bz == nil {
		return 0
	} else {
		return sdk.BigEndianToUint64(bz)
	}
}

func (k kvq[T]) setFirstItem(firstItem uint64) {
	k.root.Set(
		kvKeyFirstItem,
		sdk.Uint64ToBigEndian(firstItem),
	)
}

func (k kvq[T]) getLastItem() uint64 {
	if bz := k.root.Get(kvKeyLastItem); bz == nil {
		return 0
	} else {
		return sdk.BigEndianToUint64(bz)
	}
}

func (k kvq[T]) setLastItem(lastItem uint64) {
	k.root.Set(
		kvKeyLastItem,
		sdk.Uint64ToBigEndian(lastItem),
	)
}

func (k kvq[T]) Size() uint64 {
	lastItem := k.getLastItem()
	firstItem := k.getFirstItem()

	return lastItem - firstItem
}

func (k kvq[T]) Pick(i uint64) (T, error) {
	m := k.constructor()

	if i < k.getFirstItem() && k.getLastItem() < i {
		return m, errors.New("index out of range")
	}

	bz := k.items.Get(sdk.Uint64ToBigEndian(i))
	if bz == nil {
		return m, errors.Errorf("queue not found in index %d", i)
	}

	if err := m.Unmarshal(bz); err != nil {
		return m, errors.Wrap(err, "unmarshal")
	}
	return m, nil
}

func (k kvq[T]) Range(amount *uint64, f func(T, uint64) error) error {
	lastItem := k.getLastItem()
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

func (k kvq[T]) Consume(amount uint64) ([]T, error) {
	lastItem := k.getLastItem()
	firstItem := k.getFirstItem()
	if lastItem == firstItem {
		return nil, errors.New("empty queue")
	}

	var (
		ms       []T
		want     = min(lastItem-firstItem, amount)
		queryReq = &query.PageRequest{
			Key:     sdk.Uint64ToBigEndian(firstItem),
			Limit:   want,
			Reverse: false,
		}
	)

	_, err := query.Paginate(
		k.items, queryReq,
		func(_ []byte, value []byte) error {
			m := k.constructor()
			if err := m.Unmarshal(value); err != nil {
				return err
			}
			ms = append(ms, m)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	k.setFirstItem(firstItem + uint64(len(ms)))
	return ms, nil
}

func (k kvq[T]) MsgConstructor() func() T {
	return k.constructor
}
