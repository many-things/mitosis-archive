package queue

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gogo/protobuf/proto"
)

type Message interface {
	proto.Marshaler
	proto.Unmarshaler
}

type Queue[T Message] interface {
	// ============= immutable =============

	// Size returns the number of items in the queue.
	Size() uint64

	// Pick returns the item of specific id
	Get(uint64) (T, error)

	// Range iterates over the queue and calls the callback for each item.
	Range(*uint64, func(T, uint64) error) error

	// Paginate iterates over the queue and calls the callback for each item.
	Paginate(*query.PageRequest, func(T, uint64) error) (*query.PageResponse, error)

	// MsgConstructor returns the constructor of the message type.
	MsgConstructor() func() T

	// ============= mutable =============

	// Produce pushes the given messages to the queue.
	Produce(...T) ([]uint64, error)

	// Update updates the item of specific id.
	Update(uint64, T) error

	// Consume pops the given amount of items from the queue.
	Consume(uint64) ([]T, error)
}
