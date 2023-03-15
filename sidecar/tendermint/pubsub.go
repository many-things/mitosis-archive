package tendermint

import (
	"context"
	"errors"
	"sync"

	"github.com/smallnest/chanx"
)

type subscriber[T any] struct {
	buffer *chanx.UnboundedChan[T]
	filter func(T) bool
}

// PubSub is Generic Pub-Sub implementation for Type T
type PubSub[T any] interface {
	Publish(T) error
	Subscribe(filter func(T) bool) <-chan T
	Close()
	Done() <-chan struct{}
}

type pubSub[T any] struct {
	context        context.Context
	contextCancel  context.CancelFunc
	subscribeMutex *sync.Mutex
	subscriptions  []*subscriber[T]
	buffer         *chanx.UnboundedChan[T]
	bufferCapacity int
	done           chan struct{}
	once           *sync.Once
}

func NewPubSub[T any]() PubSub[T] {
	ctx, ctxCancel := context.WithCancel(context.Background())
	bufferCapacity := 1000

	return &pubSub[T]{
		context:        ctx,
		contextCancel:  ctxCancel,
		bufferCapacity: bufferCapacity,
		subscriptions:  []*subscriber[T]{},
		buffer:         chanx.NewUnboundedChan[T](bufferCapacity),
		subscribeMutex: &sync.Mutex{},
		done:           make(chan struct{}),
		once:           &sync.Once{},
	}
}

func (p *pubSub[T]) Publish(item T) error {
	select {
	case <-p.context.Done():
		return errors.New("[Pubsub] Not Started")
	default:
		p.buffer.In <- item
		return nil
	}
}

func (p *pubSub[T]) Subscribe(filter func(T) bool) <-chan T {
	p.subscribeMutex.Lock()
	defer p.subscribeMutex.Unlock()

	select {
	case <-p.context.Done():
		ch := make(chan T)
		close(ch)
		return ch
	default:
		subscription := &subscriber[T]{
			filter: filter,
			buffer: chanx.NewUnboundedChan[T](p.bufferCapacity),
		}

		p.subscriptions = append(p.subscriptions, subscription)
		return subscription.buffer.Out
	}
}

func (p *pubSub[T]) Close() {
	p.once.Do(func() {
		p.contextCancel()
		close(p.buffer.In)
	})
}

func (p *pubSub[T]) Done() <-chan struct{} {
	return p.done
}

// run iterate whole subscriptions every published T and send to each subscription's channel
func (p *pubSub[T]) run() {
	for item := range p.buffer.Out {
		p.subscribeMutex.Lock()
		for _, sub := range p.subscriptions {
			if sub.filter(item) {
				sub.buffer.In <- item
			}
		}
		p.subscribeMutex.Unlock()
	}

	p.subscribeMutex.Lock()
	for _, sub := range p.subscriptions {
		close(sub.buffer.In)
	}
	p.subscribeMutex.Unlock()

	close(p.done)
}
