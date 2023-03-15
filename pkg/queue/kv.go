package queue

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	kvKeyFirstItem = []byte{0x01}
	kvKeyLastItem  = []byte{0x02}
	kvPrefixItems  = []byte{0x03}
)

type kvq[T Message] struct {
	root  store.KVStore
	items store.KVStore
}

func NewKVQueue[T Message](root store.KVStore) Queue[T] {
	items := prefix.NewStore(root, kvPrefixItems)

	return kvq[T]{root, items}
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

func (k kvq[T]) Produce(msgs []T) error {
	lastItem := k.getLastItem()
	for i, msg := range msgs {
		bz, err := msg.Marshal()
		if err != nil {
			return err
		}

		k.items.Set(sdk.Uint64ToBigEndian(lastItem+uint64(i)), bz)
	}
	k.setLastItem(lastItem + uint64(len(msgs))) // -> last item
	return nil
}

func (k kvq[T]) Consume(amount uint64, conv func([]byte) (T, error)) ([]T, error) {
	lastItem := k.getLastItem()
	firstItem := k.getFirstItem()
	if lastItem == firstItem {
		return nil, errors.New("empty queue")
	}

	iter := k.items.Iterator(sdk.Uint64ToBigEndian(firstItem), sdk.Uint64ToBigEndian(lastItem))
	defer iter.Close()

	var (
		ms   []T
		want = min(lastItem-firstItem, amount)
	)
	for ; iter.Valid(); iter.Next() {
		m, err := conv(iter.Value())
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
		if uint64(len(ms)) >= want {
			break
		}
	}

	k.setFirstItem(firstItem + uint64(len(ms)))
	return ms, nil
}
