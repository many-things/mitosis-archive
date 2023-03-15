package queue

import (
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	"testing"
)

func setupKVQueue[T Message](t *testing.T) Queue[T] {
	storeKey := sdk.NewKVStoreKey("kv-test")
	memStoreKey := storetypes.NewMemoryStoreKey("mem-test")

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return NewKVQueue[T](ctx.KVStore(storeKey))
}

func TestKVQueue(t *testing.T) {
	testQueue(t, setupKVQueue[Message](t))
}
