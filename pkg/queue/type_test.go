package queue

import (
	"encoding/json"
	"fmt"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var _ Message = &TestMessage{}

type TestMessage struct {
	Data string `json:"data"`
}

func (t *TestMessage) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TestMessage) Unmarshal(bytes []byte) error {
	return json.Unmarshal(bytes, t)
}

func MakeTestMessages(size int) []Message {
	ts := make([]Message, size)
	for i := 0; i < len(ts); i++ {
		ts[i] = &TestMessage{Data: fmt.Sprintf("t%d", i)}
	}
	return ts
}

func testQueue(t *testing.T, q Queue[Message]) {
	ts := MakeTestMessages(50)

	// produce
	pushed, err := q.Produce(ts...)
	require.NoError(t, err)
	require.Equal(
		t,
		pushed,
		mitotypes.Map(ts, func(_ Message, i int) uint64 { return uint64(i) }),
	)
	require.Equal(t, uint64(len(ts)), q.Size())

	// consume
	var (
		consumeSize uint64 = 25
		msgs        []Message
	)

	// consume half
	msgs, err = q.Consume(consumeSize)
	require.NoError(t, err)
	require.Equal(t, consumeSize, uint64(len(msgs)))
	require.Equal(t, uint64(len(ts)-len(msgs)), q.Size())
	require.Equal(t, "t0", msgs[0].(*TestMessage).Data)

	// produce half
	pushed, err = q.Produce(ts[:consumeSize]...)
	require.NoError(t, err)
	require.Equal(
		t,
		pushed,
		mitotypes.Map(ts[:consumeSize], func(_ Message, i int) uint64 { return uint64(len(ts) + i) }),
	)

	// consume half again
	msgs, err = q.Consume(consumeSize)
	require.NoError(t, err)
	require.Equal(t, consumeSize, uint64(len(msgs)))
	require.Equal(t, uint64(len(ts)-len(msgs)), q.Size())
	require.Equal(t, "t25", msgs[0].(*TestMessage).Data)

	// consume rest & check produced item
	msgs, err = q.Consume(consumeSize)
	require.NoError(t, err)
	require.Equal(t, consumeSize, uint64(len(msgs)))
	require.Equal(t, uint64(len(ts)-(len(msgs)*2)), q.Size())
	require.Equal(t, "t0", msgs[0].(*TestMessage).Data)
}
