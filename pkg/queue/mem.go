package queue

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
	"sync"
)

type memq[T Message] struct {
	mux         *sync.RWMutex
	lastIdx     uint64
	store       [][]byte
	constructor func() T
}

func NewMemoryQueue[T Message](constructor func() T) Queue[T] {
	return &memq[T]{
		mux:         new(sync.RWMutex),
		store:       [][]byte{},
		constructor: constructor,
	}
}

func (k *memq[T]) Size() uint64 {
	k.mux.RLock()
	defer k.mux.RUnlock()

	return uint64(len(k.store))
}

func (k *memq[T]) Get(i uint64) (T, error) {
	m := k.constructor()

	size := k.Size()
	if i < k.lastIdx-size || k.lastIdx <= i {
		return m, errors.New("index out of range")
	}

	k.mux.Lock()
	bz := k.store[i+size-k.lastIdx]
	k.mux.Unlock()

	if err := m.Unmarshal(bz); err != nil {
		return m, errors.Wrap(err, "unmarshal")
	}

	return m, nil
}

func (k *memq[T]) Range(amount *uint64, f func(T, uint64) error) error {
	limit := uint64(query.MaxLimit)
	if amount != nil {
		limit = *amount
	}

	size := k.Size()
	l := min(size, limit)

	k.mux.Lock()
	bzs := k.store[:l]
	bi := k.lastIdx
	k.mux.Unlock()

	for i, bz := range bzs {
		m := k.constructor()
		if err := m.Unmarshal(bz); err != nil {
			return errors.Wrap(err, "unmarshal")
		}
		if err := f(m, uint64(i)+bi-size); err != nil {
			return err
		}
	}
	return nil
}

func (k *memq[T]) Paginate(req *query.PageRequest, f func(T, uint64) error) (*query.PageResponse, error) {
	panic("unimplmented")
}

func (k *memq[T]) Produce(msgs ...T) ([]uint64, error) {
	bzs := make([][]byte, len(msgs))
	for i, msg := range msgs {
		bz, err := msg.Marshal()
		if err != nil {
			return nil, err
		}
		bzs[i] = bz
	}

	k.mux.Lock()
	k.store = append(k.store, bzs...)

	idx := k.lastIdx
	k.lastIdx += uint64(len(msgs))
	k.mux.Unlock()

	return mitotypes.Map(
		make([]byte, len(msgs)),
		func(_ byte, i int) uint64 { return idx + uint64(i) },
	), nil
}

func (k *memq[T]) Update(i uint64, msg T) error {
	size := k.Size()
	if i < k.lastIdx-size && k.lastIdx < i {
		return errors.New("index out of range")
	}

	bz, err := msg.Marshal()
	if err != nil {
		return errors.Wrap(err, "marshal")
	}

	k.mux.Lock()
	k.store[i+size-k.lastIdx] = bz
	k.mux.Unlock()

	return nil
}

func (k *memq[T]) unmarshal(arr [][]byte) ([]T, error) {
	ms := make([]T, len(arr))
	for i := range ms {
		ms[i] = *new(T)
		if err := ms[i].Unmarshal(arr[i]); err != nil {
			return nil, err
		}
	}
	return ms, nil
}

func (k *memq[T]) Consume(amount uint64) ([]T, error) {
	l := min(k.Size(), amount)

	k.mux.Lock()
	bzs := k.store[:l]
	k.store = k.store[l:]
	k.mux.Unlock()

	ms := make([]T, len(bzs))
	for i, bz := range bzs {
		m := k.constructor()
		if err := m.Unmarshal(bz); err != nil {
			return nil, err
		}

		ms[i] = m
	}

	return ms, nil
}

func (k *memq[T]) MsgConstructor() func() T {
	return k.constructor
}
