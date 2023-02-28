package tendermint

import (
	"context"
	"github.com/smallnest/chanx"
	"gotest.tools/assert"
	"sync"
	"testing"
)

func MockPubSub() pubSub[int] {
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
	pubsub := MockPubSub()

	// Insert 10 Different Data
	bufferData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range bufferData {
		pubsub.Publish(value)
	}
	close(pubsub.buffer.In)

	i := 0
	for elem := range pubsub.buffer.Out {
		assert.Equal(t, bufferData[i], elem)
		i++
	}
}
