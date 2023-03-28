package queue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func setupMemQueue[T Message](_ *testing.T, constructor func() T) Queue[T] {
	return NewMemoryQueue[T](constructor)
}

func TestQueue(t *testing.T) {
	q := setupMemQueue[Message](t, ConstructTestMessage)

	testQueue(t, q)
}

func TestRaceCondition(t *testing.T) {
	var (
		q  = setupMemQueue[Message](t, ConstructTestMessage)
		ts = MakeTestMessages(50)
		wg = sync.WaitGroup{}
	)
	// Produce
	for _, m := range ts {
		wg.Add(1)
		go func(m Message) {
			defer wg.Done()
			_, err := q.Produce(m)
			assert.NoError(t, err)
		}(m)
	}
	wg.Wait()
	assert.Equal(t, uint64(len(ts)), q.Size())

	// Consume
	for range ts {
		wg.Add(1)
		go func() {
			defer wg.Done()

			msgs, err := q.Consume(1)
			assert.NoError(t, err)
			assert.Len(t, msgs, 1)
		}()
	}
	wg.Wait()
	assert.Equal(t, uint64(0), q.Size())
}

func TestMemQueue_Genesis(t *testing.T) {
	eq := setupMemQueue(t, ConstructTestMessage)
	nq := setupMemQueue(t, ConstructTestMessage)

	tms := MakeTestMessages(10)
	_, err := eq.Produce(tms...)
	require.NoError(t, err)

	eg, err := eq.ExportGenesis()
	require.NoError(t, err)

	require.NoError(t, nq.ImportGenesis(eg))

	ng, err := nq.ExportGenesis()
	require.NoError(t, err)

	require.Equal(t, eg, ng)
}
