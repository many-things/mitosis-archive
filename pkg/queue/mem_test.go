package queue

import "testing"

func setup(t *testing.T) Queue[Message] {
	return NewMemoryQueue[Message]()
}
