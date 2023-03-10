package queue

import "github.com/gogo/protobuf/proto"

type Message interface {
	proto.Marshaler
	proto.Unmarshaler
}

type Queue[T Message] interface {
	Produce(msg T) error
	Consume(amount int) ([]T, error)
}
