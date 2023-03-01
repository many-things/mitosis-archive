package tendermint

import (
	"context"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"gotest.tools/assert"
	"sync"
	"testing"
	"time"
)

type mockHTTP struct {
	Height int64
	lock   sync.Mutex
}

func (m *mockHTTP) BlockchainInfo(ctx context.Context, minHeight int64, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) {
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
	mHttp := mockHTTP{
		Height: 10,
	}
	blockListener := NewBlockListener(context.Background(), &mHttp, time.Millisecond*500)

	height, err := blockListener.GetLatestBlockHeight()
	assert.NilError(t, err)
	assert.Equal(t, height, int64(10))
}

func Test_GetBlockHeight(t *testing.T) {
	mHttp := mockHTTP{
		Height: 10,
	}
	blockListener := NewBlockListener(context.Background(), &mHttp, time.Second)

	heightChan, _ := blockListener.GetBlockHeight()
	var i int64
	for i = 10; i < 13; i++ {
		mHttp.ChangeHeight(i)
		select {
		case elem := <-heightChan:
			assert.Equal(t, i, elem)
		}
	}
	blockListener.Close()
}

func Test_GetNewBlock(t *testing.T) {
	mHttp := mockHTTP{
		Height: 10,
	}
	blockListener := NewBlockListener(context.Background(), &mHttp, time.Millisecond*500)

	newBlockChan, _ := blockListener.NewBlockWatcher()
	time.Sleep(time.Second)
	mHttp.ChangeHeight(20)

	var i int64
	for i = 11; i <= 20; i++ {
		select {
		case elem := <-newBlockChan:
			assert.Equal(t, elem, i)
		}
	}
}
