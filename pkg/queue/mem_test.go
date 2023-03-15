package queue

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func setupMemQueue[T Message](t *testing.T) Queue[T] {
	return NewMemoryQueue[T]()
}

func TestQueue(t *testing.T) {
	testQueue(t, setupMemQueue[Message](t))
}

func TestRaceCondition(t *testing.T) {
	var (
		q  = setupMemQueue[Message](t)
		ts = MakeTestMessages(50)
		wg = sync.WaitGroup{}
	)
	// Produce
	for _, m := range ts {
		wg.Add(1)
		go func(m Message) {
			defer wg.Done()
			assert.NoError(t, q.Produce(m))
		}(m)
	}
	wg.Wait()
	assert.Equal(t, uint64(len(ts)), q.Size())

	// Consume
	for range ts {
		wg.Add(1)
		go func() {
			defer wg.Done()

			msgs, err := q.Consume(1, ConvTestMessage)
			assert.NoError(t, err)
			assert.Len(t, msgs, 1)
		}()
	}
	wg.Wait()
	assert.Equal(t, uint64(0), q.Size())
}
