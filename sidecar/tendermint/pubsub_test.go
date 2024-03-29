package tendermint

import (
	"context"
	"sync"
	"testing"

	"github.com/smallnest/chanx"
	"gotest.tools/assert"
)

func mockPubSub() pubSub[int] {
	ctx, ctxCancel := context.WithCancel(context.Background())
	bufferCapacity := 1000

	return pubSub[int]{
		context:        ctx,
		contextCancel:  ctxCancel,
		subscribeMutex: &sync.Mutex{},
		subscriptions:  []*subscriber[int]{},
		buffer:         chanx.NewUnboundedChan[int](bufferCapacity),
		bufferCapacity: bufferCapacity,
		done:           make(chan struct{}),
		once:           &sync.Once{},
	}
}

func Test_Publish(t *testing.T) {
	pubsub := mockPubSub()

	// Insert 10 Different Data
	bufferData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range bufferData {
		err := pubsub.Publish(value)
		assert.NilError(t, err)
	}
	pubsub.Close()

	i := 0
	for elem := range pubsub.buffer.Out {
		assert.Equal(t, bufferData[i], elem)
		i++
	}

	pubsub.Run()
	err := pubsub.Publish(1)
	assert.Error(t, err, "[Pubsub] Not Started")
}

func Test_Subscribe(t *testing.T) {
	// Add Successfully on Subscription
	pubsub := mockPubSub()

	oddFilter := pubsub.Subscribe(func(item int) bool {
		return item%2 == 1
	})
	evenFilter := pubsub.Subscribe(func(item int) bool {
		return item%2 == 0
	})

	assert.Equal(t, len(pubsub.subscriptions), 2)

	bufferData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range bufferData {
		err := pubsub.Publish(value)
		assert.NilError(t, err)
	}
	pubsub.Close() // Do not wait
	pubsub.Run()

	oddExpectData := []int{1, 3, 5, 7, 9}
	i := 0
	for elem := range oddFilter {
		assert.Equal(t, oddExpectData[i], elem)
		i++
	}
	assert.Equal(t, i, 5) // All event consumed successfully

	evenExpectData := []int{2, 4, 6, 8, 10}
	i = 0
	for elem := range evenFilter {
		assert.Equal(t, evenExpectData[i], elem)
		i++
	}
	assert.Equal(t, i, 5) // All event consumed successfully
}
