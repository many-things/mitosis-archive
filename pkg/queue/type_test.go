package queue

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
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

func ConvTestMessage(bz []byte) (Message, error) {
	var v TestMessage
	if err := v.Unmarshal(bz); err != nil {
		return nil, err
	}
	return &v, nil
}

func testQueue(t *testing.T, q Queue[Message]) {
	ts := MakeTestMessages(50)

	// produce
	assert.NoError(t, q.Produce(ts...))
	assert.Equal(t, uint64(len(ts)), q.Size())

	// consume
	var (
		consumeSize uint64 = 25
		msgs        []Message
		err         error
	)

	// consume half
	msgs, err = q.Consume(consumeSize, ConvTestMessage)
	assert.NoError(t, err)
	assert.Equal(t, consumeSize, uint64(len(msgs)))
	assert.Equal(t, uint64(len(ts)-len(msgs)), q.Size())
	assert.Equal(t, "t0", msgs[0].(*TestMessage).Data)

	// produce half
	assert.NoError(t, q.Produce(ts[:consumeSize]...))

	// consume half again
	msgs, err = q.Consume(consumeSize, ConvTestMessage)
	assert.NoError(t, err)
	assert.Equal(t, consumeSize, uint64(len(msgs)))
	assert.Equal(t, uint64(len(ts)-len(msgs)), q.Size())
	assert.Equal(t, "t25", msgs[0].(*TestMessage).Data)

	// consume rest & check produced item
	msgs, err = q.Consume(consumeSize, ConvTestMessage)
	assert.NoError(t, err)
	assert.Equal(t, consumeSize, uint64(len(msgs)))
	assert.Equal(t, uint64(len(ts)-(len(msgs)*2)), q.Size())
	assert.Equal(t, "t0", msgs[0].(*TestMessage).Data)
}
