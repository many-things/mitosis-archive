package tendermint

import (
	"context"
	golog "log"
	"sync"
	"time"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type BlockFetcher interface {
	BlockchainInfo(
		ctx context.Context,
		minHeight,
		maxHeight int64,
	) (*coretypes.ResultBlockchainInfo, error)

	BlockResults(
		ctx context.Context,
		height *int64,
	) (*coretypes.ResultBlockResults, error)
}

type BlockListener interface {
	GetLatestBlockHeight() (int64, error)
	GetBlockHeight() (<-chan int64, <-chan error)
	NewBlockWatcher() (<-chan int64, <-chan error)
	GetBlockResult(blockHeight *int64) (*coretypes.ResultBlockResults, error)
	Close()
}

type blockListener struct {
	ctx            context.Context
	ctxCancel      context.CancelFunc
	client         BlockFetcher
	listenInterval time.Duration
	once           *sync.Once
}

func NewBlockListener(ctx context.Context, client BlockFetcher, interval time.Duration) BlockListener {
	ctx, ctxCancel := context.WithCancel(ctx)
	return &blockListener{
		client:         client,
		listenInterval: interval,
		ctx:            ctx,
		ctxCancel:      ctxCancel,
		once:           &sync.Once{},
	}
}

// GetLatestBlockHeight Get The Latest Block Height from Client
func (b *blockListener) GetLatestBlockHeight() (int64, error) {
	blockInfo, err := b.client.BlockchainInfo(b.ctx, 0, 0)
	if err != nil {
		return 0, err
	}

	return blockInfo.LastHeight, nil
}

func (b *blockListener) GetBlockResult(blockHeight *int64) (*coretypes.ResultBlockResults, error) {
	return b.client.BlockResults(b.ctx, blockHeight)
}

// GetBlockHeight Returns Channel that send the Latest Block Height every listenInterval
func (b *blockListener) GetBlockHeight() (<-chan int64, <-chan error) {
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
				blockHeight, err = b.GetLatestBlockHeight()

				// TODO: remove logs
				golog.Printf("Watch blockHeight: %d\n", blockHeight)
				if err != nil {
					errChan <- err
					return
				}
				keepAlive, keepAliveCancel = context.WithTimeout(context.Background(), b.listenInterval)
			case <-b.ctx.Done():
				return
			}

			select {
			// TODO: change into default
			case blockHeightChan <- blockHeight:
				break // nolint: revive
			case <-b.ctx.Done():
				return
			}
		}
	}()

	return blockHeightChan, errChan
}

// NewBlockWatcher Returns Channel that send New Block Height
func (b *blockListener) NewBlockWatcher() (<-chan int64, <-chan error) {
	newBlockHeightChan := make(chan int64, 1)
	errChan := make(chan error, 1)

	blockHeightChan, watchErrChan := b.GetBlockHeight()

	go func() {
		defer close(newBlockHeightChan)
		latestBlockHeight, err := b.GetLatestBlockHeight()
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
			case <-b.ctx.Done():
				return
			}

			// Processing current Block is more important than receive new block
			for processedBlockHeight < latestBlockHeight {
				select {
				case newBlockHeightChan <- processedBlockHeight + 1:
					processedBlockHeight++
				case <-b.ctx.Done():
					return
				}
			}
		}
	}()

	return newBlockHeightChan, errChan
}

func (b *blockListener) Close() {
	b.once.Do(func() {
		b.ctxCancel()
	})
}
