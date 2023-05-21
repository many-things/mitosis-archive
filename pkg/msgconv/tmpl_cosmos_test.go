package msgconv

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/types"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestCosmosOp0(t *testing.T) {
	rendered, err := CosmosOp0(
		"chain",
		"vault",
		types.Join(
			[]byte("vaultvault"),
		),
		[]*types.Coin{},
	)
	require.Nil(t, err)

	log.Println(string(rendered))
}

func TestCosmosOp1(t *testing.T) {
	rendered, err := CosmosOp1(
		"chain",
		"vault",
		types.Join(
			[]byte("osmo1pe6llrv0y5vz0c9msdg2kndes9eh6jf620hjll"),
			[]byte("uosmo"),
			[]byte("500000"),
		),
		[]*types.Coin{{
			Denom:   "uusdc",
			Amount:  types.Ref(sdk.NewInt(100000)),
			Decimal: 0,
		}},
	)
	require.Nil(t, err)
	require.Contains(t, string(rendered), "0a057661756c74120908101205756f736d6f1a0812063130303030302206353030303030")
}
