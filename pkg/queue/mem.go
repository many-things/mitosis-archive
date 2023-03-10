package queue

type memq[T Message] struct {
	store [][]byte
}

func NewMemoryQueue[T Message]() Queue[T] {
	return memq[T]{}
}

func (k memq[T]) Produce(msg T) error {
	bz, err := msg.Marshal()
	if err != nil {
		return err
	}

	k.store = append(k.store, bz)

	return nil
}

func (k memq[T]) unmarshal(arr [][]byte) ([]T, error) {
	ms := make([]T, len(arr))
	for i := range ms {
		if err := ms[i].Unmarshal(arr[i]); err != nil {
			return nil, err
		}
	}
	return ms, nil
}

func (k memq[T]) Consume(amount int) ([]T, error) {
	l := min(len(k.store), amount)
	ms, err := k.unmarshal(k.store[:l])
	if err != nil {
		return nil, err
	}

	k.store = k.store[l:]
	return ms, nil
}
