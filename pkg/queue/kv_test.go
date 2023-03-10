package queue

import "testing"

func setupMemQueue(t *testing.T) Queue[Message] {
	return NewMemoryQueue[Message]()
}

//func setupKVQueue(t *testing.T) Queue[Message] {
//	return NewKVQueue[Message]()
//}
