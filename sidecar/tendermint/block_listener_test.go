package tendermint

import (
	"context"
	"sync"
	"testing"
	"time"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"gotest.tools/assert"
)

type mockHTTP struct {
	Height int64
	lock   sync.Mutex
}

func (m *mockHTTP) BlockchainInfo(_ context.Context, _ int64, _ int64) (*coretypes.ResultBlockchainInfo, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	return &coretypes.ResultBlockchainInfo{
		LastHeight: m.Height,
		BlockMetas: nil,
	}, nil
}

func (m *mockHTTP) ChangeHeight(newHeight int64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Height = newHeight
}

func Test_GetLatestBlockHeight(t *testing.T) {
	mHTTP := mockHTTP{
		Height: 10,
	}
	blockListener := NewBlockListener(context.Background(), &mHTTP, time.Millisecond*500)

	height, err := blockListener.GetLatestBlockHeight()
	assert.NilError(t, err)
	assert.Equal(t, height, int64(10))
}

func Test_GetBlockHeight(t *testing.T) {
	mHTTP := mockHTTP{
		Height: 10,
	}
	blockListener := NewBlockListener(context.Background(), &mHTTP, time.Second)

	heightChan, _ := blockListener.GetBlockHeight()
	var i int64
	for i = 10; i < 13; i++ {
		mHTTP.ChangeHeight(i)
		elem := <-heightChan
		assert.Equal(t, i, elem)
	}
	blockListener.Close()
}

func Test_GetNewBlock(t *testing.T) {
	mHTTP := mockHTTP{
		Height: 10,
	}
	blockListener := NewBlockListener(context.Background(), &mHTTP, time.Millisecond*500)

	newBlockChan, _ := blockListener.NewBlockWatcher()
	time.Sleep(time.Second)
	mHTTP.ChangeHeight(20)

	var i int64
	for i = 11; i <= 20; i++ {
		elem := <-newBlockChan
		assert.Equal(t, elem, i)
	}
}
