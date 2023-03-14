package queue

import (
	"testing"
)

func setupMemQueue[T Message](t *testing.T) Queue[T] {
	return NewMemoryQueue[T]()
}

func TestMemoryQueue(t *testing.T) {
	testQueue(t, setupMemQueue[Message](t))
}
