package msgconv

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestCosmosOp0(t *testing.T) {
	rendered, err := CosmosOp0("vault", []byte("vaultvault"), []byte("1000000000stake"))
	require.Nil(t, err)

	log.Println(string(rendered))
}
