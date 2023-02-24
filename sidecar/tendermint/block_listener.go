package tendermint

import (
	"context"
	tmHTTP "github.com/tendermint/tendermint/rpc/client/http"
	"time"
)

type blockListener struct {
	client         *tmHTTP.HTTP
	listenInterval time.Duration
}

func NewBlockListener(client *tmHTTP.HTTP, interval time.Duration) *blockListener {
	return &blockListener{
		client:         client,
		listenInterval: interval,
	}
}

func (b *blockListener) GetLatestBlockHeight(ctx context.Context) (*int64, error) {
	blockInfo, err := b.client.BlockchainInfo(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	return &blockInfo.LastHeight, nil
}

func (b *blockListener) GetBlockHeight(ctx context.Context) (<-chan int64, <-chan error) {
	blockHeightChan := make(chan int64)
	errChan := make(chan error, 1)

	var keepAlive context.Context
	var keepAliveCancel context.CancelFunc = func() {}
	go func() {
		defer close(blockHeightChan)
		defer func() { keepAliveCancel() }()

		keepAlive, keepAliveCancel = context.WithTimeout(context.Background(), b.listenInterval)
		var blockHeight *int64
		var err error
		for {
			select {
			case <-keepAlive.Done():
				blockHeight, err = b.GetLatestBlockHeight(context.Background())
				if err != nil {
					errChan <- err
					return
				}
			case <-ctx.Done():
				return
			}

			select {
			case blockHeightChan <- *blockHeight:
				break
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockHeightChan, errChan
}
