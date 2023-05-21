package msgconv

import (
	"log"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestCosmosOp0(t *testing.T) {
	rendered, err := CosmosOp0(
		"evm-5",
		"osmo-test-5",
		"vault",
		types.Join(
			[]byte("vaultvault"),
		),
		[]*types.Coin{{
			Denom:   "0x5Cbb2F9f7E54c5B4656C3B563ff5650a0866A3EF",
			Amount:  types.Ref(sdk.NewInt(100000)),
			Decimal: 0,
		}},
	)
	require.Nil(t, err)

	conv, err := convertDenomIO("evm-5", "osmo-test-5", "0x5Cbb2F9f7E54c5B4656C3B563ff5650a0866A3EF")
	require.Nil(t, err)
	require.Contains(t, string(rendered), conv)

	log.Println(string(rendered))
}

func TestCosmosOp1(t *testing.T) {
	rendered, err := CosmosOp1(
		"evm-5",
		"osmo-test-5",
		"vault",
		types.Join(
			[]byte("osmo1pe6llrv0y5vz0c9msdg2kndes9eh6jf620hjll"),
			[]byte("uosmo"),
			[]byte("500000"),
		),
		[]*types.Coin{{
			Denom:   "0x5Cbb2F9f7E54c5B4656C3B563ff5650a0866A3EF",
			Amount:  types.Ref(sdk.NewInt(100000)),
			Decimal: 0,
		}},
	)
	require.Nil(t, err)
	require.Contains(t, string(rendered), "0a057661756c74120908101205756f736d6f1a430a39666163746f72792f6f736d6f3130396e73347530346c34346b71646b767038373668756b643368787a387a7a6d37383039656c2f757573646312063130303030302206353030303030")

	// _, err = CosmosOp1(
	//	"evm-5",
	//	"osmo-test-5",
	//	"vault",
	//	types.Join(
	//		[]byte("osmo1pe6llrv0y5vz0c9msdg2kndes9eh6jf620hjll"),
	//		[]byte("uosmo"),
	//		[]byte("500000"),
	//	),
	//	[]*types.Coin{},
	// )
	// require.Error(t, err, "expected exactly one fund")
}
