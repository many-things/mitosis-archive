package queue

import "github.com/gogo/protobuf/proto"

type Message interface {
	proto.Marshaler
	proto.Unmarshaler
}

type Queue[T Message] interface {
	Size() uint64
	Pick(uint64) (T, error)
	Range(*uint64, func(T, uint64) error) error
	Produce(msgs ...T) ([]uint64, error)
	Consume(amount uint64) ([]T, error)
	MsgConstructor() func() T
}
