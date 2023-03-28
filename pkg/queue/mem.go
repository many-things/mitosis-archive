package queue

import (
	"sort"
	"sync"

	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/pkg/errors"
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

// Size returns the number of items in the queue.
func (k *memq[T]) Size() uint64 {
	k.mux.RLock()
	defer k.mux.RUnlock()

	return uint64(len(k.store))
}

// LastIndex returns the last item's index of the queue.
func (k *memq[T]) LastIndex() uint64 {
	return k.lastIdx
}

// Get returns the item of specific id
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

// Range iterates over the queue and calls the callback for each item.
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

// Paginate iterates over the queue and calls the callback for each item.
func (k *memq[T]) Paginate(_ *query.PageRequest, _ func(T, uint64) error) (*query.PageResponse, error) {
	panic("unimplmented")
}

// MsgConstructor returns the constructor of the message type.
func (k *memq[T]) MsgConstructor() func() T {
	return k.constructor
}

// Produce pushes the given messages to the queue.
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

// Update updates the item of specific id.
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
		ms[i] = *new(T) // nolint: gocritic
		if err := ms[i].Unmarshal(arr[i]); err != nil {
			return nil, err
		}
	}
	return ms, nil
}

// Consume pops the given amount of items from the queue.
func (k *memq[T]) Consume(amount uint64) ([]mitotypes.KV[uint64, T], error) {
	l := min(k.Size(), amount)

	k.mux.Lock()
	size := uint64(len(k.store))
	bzs := k.store[:l]
	k.store = k.store[l:]
	k.mux.Unlock()

	ms := make([]mitotypes.KV[uint64, T], len(bzs))
	for ri, bz := range bzs {
		ai := k.lastIdx - size + uint64(ri) // absolute index

		m := k.constructor()
		if err := m.Unmarshal(bz); err != nil {
			return nil, err
		}

		ms[ri] = mitotypes.NewKV(ai, m)
	}

	return ms, nil
}

// ConsumeUntil pops the items from the queue until the given condition is met.
func (k *memq[T]) ConsumeUntil(f func(T, uint64) (bool, error)) ([]mitotypes.KV[uint64, T], error) {
	k.mux.Lock()
	defer k.mux.Unlock()

	var (
		ms   []mitotypes.KV[uint64, T]
		size = uint64(len(k.store))
	)
	for ri, bz := range k.store {
		ai := k.lastIdx - size + uint64(ri) // absolute index

		m := k.constructor()
		if err := m.Unmarshal(bz); err != nil {
			return nil, err
		}
		ms = append(ms, mitotypes.NewKV(ai, m))

		if ok, err := f(m, ai); err != nil {
			return nil, err
		} else if ok {
			break
		}
	}

	k.store = k.store[len(ms):]
	return ms, nil
}

// ImportGenesis imports the queue's genesis state.
func (k *memq[T]) ImportGenesis(g GenesisState[T]) error {
	k.mux.Lock()
	defer k.mux.Unlock()

	k.store = [][]byte{}
	k.lastIdx = g.LastIndex

	sort.Slice(g.Items, func(i, j int) bool {
		return g.Items[i].Key < g.Items[j].Key
	})

	for _, item := range g.Items {
		bz, err := item.Value.Marshal()
		if err != nil {
			return err
		}

		k.store = append(k.store, bz)
	}

	return nil
}

// ExportGenesis exports the queue's genesis state.
func (k *memq[T]) ExportGenesis() (GenesisState[T], error) {
	k.mux.Lock()
	defer k.mux.Unlock()

	g := GenesisState[T]{
		FirstIndex: 0,
		LastIndex:  k.lastIdx,
	}

	for i, bz := range k.store {
		m := k.constructor()
		if err := m.Unmarshal(bz); err != nil {
			return GenesisState[T]{}, err
		}
		g.Items = append(g.Items, mitotypes.NewKV(uint64(i), m))
	}

	return g, nil
}
