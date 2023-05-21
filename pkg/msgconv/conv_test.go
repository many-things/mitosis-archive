package msgconv

import (
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestToMsgsDecodeAndEncodeABIValue(t *testing.T) {
	encodeABI, err := sigToABI("test(string)")
	require.Nil(t, err)

	arrstr := []string{"hello", "world", "yolo", "jake", "swag"}
	opArgs, err := mitotypes.MapErr(
		arrstr,
		func(str string, _ int) ([]byte, error) {
			return encodeABI.Methods["test"].Inputs.Pack(str)
		},
	)
	require.Nil(t, err)

	txPayload, toByteToSign, err := ToMsgs("evm-5", "osmo-test-5", "vault", 0, opArgs, []*mitotypes.Coin{})
	require.Nil(t, err)

	log.Println(len(txPayload), string(txPayload))
	assert.Equal(t, len(toByteToSign), 32)
}

func TestToMsgsKeccak(t *testing.T) {
	txPayload, txBytesToSign, err := ToMsgs(
		"osmo-test-5", "evm-5", "vault", 0,
		mitotypes.Join([]byte("op0")),
		[]*mitotypes.Coin{},
	)
	require.Nil(t, err)

	log.Println(len(txPayload), string(txPayload))
	assert.Equal(t, len(txBytesToSign), 32)
}

func TestToMsgsSha(t *testing.T) {
	txPayload, txBytesToSign, err := ToMsgs(
		"osmo-test-5", "evm-5", "vault", 0,
		mitotypes.Join([]byte("op0")),
		[]*mitotypes.Coin{},
	)
	require.Nil(t, err)

	log.Println(len(txPayload), string(txPayload))
	assert.Equal(t, len(txBytesToSign), 32)
}
