package queue

import "github.com/gogo/protobuf/proto"

type Message interface {
	proto.Marshaler
	proto.Unmarshaler
}

type Queue[T Message] interface {
	Size() uint64
	Produce(msgs []T) error
	Consume(amount uint64, conv func([]byte) (T, error)) ([]T, error)
}
