package tendermint

import (
	"context"
	tmHTTP "github.com/tendermint/tendermint/rpc/client/http"
	"time"
)

type BlockListener struct {
	client         *tmHTTP.HTTP
	listenInterval time.Duration
	done           chan struct{}
}

func NewBlockListener(client *tmHTTP.HTTP, interval time.Duration) *BlockListener {
	return &BlockListener{
		client:         client,
		listenInterval: interval,
	}
}

// GetLatestBlockHeight Get The Latest Block Height from Client
func (b *BlockListener) GetLatestBlockHeight(ctx context.Context) (int64, error) {
	blockInfo, err := b.client.BlockchainInfo(ctx, 0, 0)
	if err != nil {
		return 0, err
	}

	return blockInfo.LastHeight, nil
}

// GetBlockHeight Returns Channel that send the Latest Block Height every listenInterval
func (b *BlockListener) GetBlockHeight(ctx context.Context) (<-chan int64, <-chan error) {
	blockHeightChan := make(chan int64)
	errChan := make(chan error, 1)

	var keepAlive context.Context
	var keepAliveCancel context.CancelFunc = func() {}
	go func() {
		defer close(blockHeightChan)
		defer func() { keepAliveCancel() }()

		keepAlive, keepAliveCancel = context.WithTimeout(context.Background(), b.listenInterval)
		var blockHeight int64
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
			case blockHeightChan <- blockHeight:
				break
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockHeightChan, errChan
}

// NewBlockWatcher Returns Channel that send New Block Height
func (b *BlockListener) NewBlockWatcher(ctx context.Context) (<-chan int64, <-chan error) {
	newBlockHeightChan := make(chan int64, 1)
	errChan := make(chan error, 1)

	blockHeightChan, watchErrChan := b.GetBlockHeight(ctx)

	go func() {
		defer close(newBlockHeightChan)
		latestBlockHeight, err := b.GetLatestBlockHeight(context.Background())
		if err != nil {
			errChan <- err
			return
		}

		processedBlockHeight := latestBlockHeight
		for {
			select {
			case newBlockHeight := <-blockHeightChan:
				if latestBlockHeight >= newBlockHeight {
					continue
				}
				latestBlockHeight = newBlockHeight
			case err = <-watchErrChan:
				errChan <- err
				return
			case <-ctx.Done():
				return
			}

			// Processing current Block is more important than receive new block
			for processedBlockHeight < latestBlockHeight {
				select {
				case newBlockHeightChan <- processedBlockHeight + 1:
					processedBlockHeight++
					break
				case <-ctx.Done():
					close(b.done)
					return
				}
			}
		}
	}()

	return newBlockHeightChan, errChan
}

func (b *BlockListener) Done() <-chan struct{} {
	return b.done
}
