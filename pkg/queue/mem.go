package queue

import "sync"

type memq[T Message] struct {
	mux   *sync.RWMutex
	store [][]byte
}

func NewMemoryQueue[T Message]() Queue[T] {
	return &memq[T]{
		mux:   new(sync.RWMutex),
		store: [][]byte{},
	}
}

func (k *memq[T]) Size() uint64 {
	return uint64(len(k.store))
}

func (k *memq[T]) Produce(msgs ...T) error {
	bzs := make([][]byte, len(msgs))
	for i, msg := range msgs {
		bz, err := msg.Marshal()
		if err != nil {
			return err
		}
		bzs[i] = bz
	}

	k.mux.Lock()
	k.store = append(k.store, bzs...)
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

func (k *memq[T]) Consume(amount uint64, conv func([]byte) (T, error)) ([]T, error) {
	l := min(uint64(len(k.store)), amount)

	k.mux.Lock()
	bzs := k.store[:l]
	k.store = k.store[l:]
	k.mux.Unlock()

	ms := make([]T, len(bzs))
	for i, bz := range bzs {
		m, err := conv(bz)
		if err != nil {
			return nil, err
		}

		ms[i] = m
	}

	return ms, nil
}
