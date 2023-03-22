package queue

import (
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	"testing"
)

func setupKVQueue[T Message](t *testing.T, constructor func() T) Queue[T] {
	storeKey := sdk.NewKVStoreKey("kv-test")
	memStoreKey := storetypes.NewMemoryStoreKey("mem-test")

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return NewKVQueue[T](ctx.KVStore(storeKey), constructor)
}

func TestKVQueue(t *testing.T) {
	q := setupKVQueue[Message](t, ConstructTestMessage)

	testQueue(t, q)
}

func TestKVQueue_Paginate(t *testing.T) {
	var (
		err error
		q   = setupKVQueue[Message](t, ConstructTestMessage)
		tms = MakeTestMessages(12)
	)

	_, err = q.Produce(tms...)
	require.NoError(t, err)

	_, err = q.Consume(2)
	require.NoError(t, err)

	{ // Test happy case
		for _, c := range []bool{true, false} {
			resp, err := q.Paginate(
				&query.PageRequest{Reverse: c},
				func(m Message, i uint64) error {
					require.Less(t, i, uint64(len(tms)))
					require.GreaterOrEqual(t, i, uint64(2))
					return nil
				},
			)
			require.NoError(t, err)
			require.Nil(t, resp.NextKey)
		}
	}

	{ // Test overflowed key
		cs := []struct {
			key     uint64
			reverse bool
		}{
			{key: 1, reverse: false},
			{key: 13, reverse: true},
		}
		for _, c := range cs {
			_, err = q.Paginate(
				&query.PageRequest{Key: sdk.Uint64ToBigEndian(c.key), Reverse: c.reverse},
				func(m Message, i uint64) error { return nil },
			)
			require.Error(t, err, "key out of range")
		}
	}

	{ // Test next key
		cs := []struct {
			key      uint64
			limit    uint64
			reverse  bool
			expected uint64
		}{
			{key: 5, limit: 3, reverse: false, expected: 8},
			{key: 8, limit: 3, reverse: true, expected: 5},
		}
		for _, c := range cs {
			resp, err := q.Paginate(
				&query.PageRequest{Key: sdk.Uint64ToBigEndian(c.key), Limit: c.limit, Reverse: c.reverse},
				func(m Message, i uint64) error { return nil },
			)
			require.NoError(t, err)
			require.Equal(t, sdk.Uint64ToBigEndian(c.expected), resp.NextKey)
		}
	}
}

func TestKVQueue_Genesis(t *testing.T) {
	eq := setupKVQueue(t, ConstructTestMessage)
	nq := setupKVQueue(t, ConstructTestMessage)

	tms := MakeTestMessages(10)
	_, err := eq.Produce(tms...)
	require.NoError(t, err)

	eg, err := eq.ExportGenesis()
	require.NoError(t, err)

	require.NoError(t, nq.ImportGenesis(eg))

	ng, err := nq.ExportGenesis()
	require.NoError(t, err)

	require.Equal(t, eg, ng)
}
