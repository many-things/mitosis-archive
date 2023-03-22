package queue

import "github.com/cosmos/cosmos-sdk/codec"

type Message interface {
	codec.ProtoMarshaler
}

type Queue[T Message] interface {
	Size() uint64
	Pick(uint64) (T, error)
	Produce(msgs ...T) ([]uint64, error)
	Consume(amount uint64) ([]T, error)
	MsgConstructor() func() T
}
