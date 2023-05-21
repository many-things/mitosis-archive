package msgconv

import (
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestToMsgsKeccak(t *testing.T) {
	txPayload, txBytesToSign, err := ToMsgs(
		"evm-5", "vault", 0,
		mitotypes.Join([]byte("op0")),
		[]*mitotypes.Coin{},
	)
	require.Nil(t, err)

	log.Println(len(txPayload), string(txPayload))
	assert.Equal(t, len(txBytesToSign), 32)
}

func TestToMsgsSha(t *testing.T) {
	txPayload, txBytesToSign, err := ToMsgs(
		"osmo-test-5", "vault", 0,
		mitotypes.Join([]byte("op0")),
		[]*mitotypes.Coin{},
	)
	require.Nil(t, err)

	log.Println(len(txPayload), string(txPayload))
	assert.Equal(t, len(txBytesToSign), 32)
}
