package queue

type memq[T Message] struct {
	store [][]byte
}

func NewMemoryQueue[T Message]() Queue[T] {
	return &memq[T]{}
}

func (k *memq[T]) Size() uint64 {
	return uint64(len(k.store))
}

func (k *memq[T]) Produce(msgs []T) error {
	for _, msg := range msgs {
		bz, err := msg.Marshal()
		if err != nil {
			return err
		}

		k.store = append(k.store, bz)
	}

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

func (k *memq[T]) Consume(amount uint64, f func([]byte) (T, error)) ([]T, error) {
	l := min(uint64(len(k.store)), amount)

	arr := k.store[:l]
	k.store = k.store[l:]

	ms := make([]T, len(arr))
	for i, bz := range arr {
		m, err := f(bz)
		if err != nil {
			return nil, err
		}

		ms[i] = m
	}

	return ms, nil
}
