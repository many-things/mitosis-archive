package queue

import "github.com/cosmos/cosmos-sdk/codec"

type Message interface {
	codec.ProtoMarshaler
}

type Queue[T Message] interface {
	Size() uint64
	Produce(msgs ...T) error
	Consume(amount uint64, conv func([]byte) (T, error)) ([]T, error)
}
