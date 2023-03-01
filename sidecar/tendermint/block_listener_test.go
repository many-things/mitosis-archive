package tendermint

import (
	"context"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"gotest.tools/assert"
	"testing"
	"time"
)

type mockHTTP struct {
	Height int64
}

func (m mockHTTP) BlockchainInfo(ctx context.Context, minHeight int64, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) {
	return &coretypes.ResultBlockchainInfo{
		LastHeight: m.Height,
		BlockMetas: nil,
	}, nil
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
	blockListener := NewBlockListener(context.Background(), &mHttp, time.Millisecond*500)

	heightChan, _ := blockListener.GetBlockHeight()
	var i int64
	for i = 10; i < 13; i++ {
		mHttp.Height = i
		select {
		case elem := <-heightChan:
			assert.Equal(t, i, elem)
		}
	}
	blockListener.Close()
}
