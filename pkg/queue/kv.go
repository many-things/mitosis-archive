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

func (k kvq[T]) Produce(msg T) error {
	lastItem := k.getLastItem()
	bz, err := msg.Marshal()
	if err != nil {
		return err
	}

	k.items.Set(sdk.Uint64ToBigEndian(lastItem), bz)
	k.setLastItem(lastItem + 1) // -> last item
	return nil
}

func (k kvq[T]) Consume(amount int) ([]T, error) {
	lastItem := k.getLastItem()
	firstItem := k.getFirstItem()
	if lastItem-firstItem == 0 {
		return nil, errors.New("empty queue")
	}

	iter := k.items.Iterator(sdk.Uint64ToBigEndian(firstItem), sdk.Uint64ToBigEndian(lastItem))
	defer iter.Close()

	ms := make([]T, min(lastItem-firstItem, uint64(amount)))
	for i := 0; iter.Valid(); iter.Next() {
		if err := ms[i].Unmarshal(iter.Value()); err != nil {
			return nil, err
		}
		i++
	}

	k.setFirstItem(firstItem + uint64(len(ms)))
	return ms, nil
}
