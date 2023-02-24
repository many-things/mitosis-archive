package tendermint

import (
	"context"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmHttp "github.com/tendermint/tendermint/rpc/client/http"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type TmEvent struct {
	BlockHeight int64
	abci.Event
}

type TmEventBus struct {
	client   *tmHttp.HTTP
	listener BlockListener
	logger   log.Logger
	pubSub   PubSub[TmEvent]
}

func NewTmEventBus(listener BlockListener, pubSub PubSub[TmEvent], logger log.Logger) *TmEventBus {
	return &TmEventBus{
		listener: listener,
		pubSub:   pubSub,
		logger:   logger,
	}
}

func (tb *TmEventBus) ListenBlock(ctx context.Context) (<-chan coretypes.ResultBlockResults, <-chan error) {
	blockResultChan := make(chan coretypes.ResultBlockResults)
	errChan := make(chan error, 1)

	go func() {
		newBlockHeightChan, heightErrChan := tb.listener.NewBlockWatcher(context.Background())

		var blockHeight int64
		for {
			select {
			case blockHeight = <-newBlockHeightChan:
				block, err := tb.client.BlockResults(context.Background(), &blockHeight)
				if err != nil {
					errChan <- err
					return
				}

				blockResultChan <- *block
			case blockHeightErr := <-heightErrChan:
				errChan <- blockHeightErr
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return blockResultChan, errChan
}

func (tb *TmEventBus) publish(block *coretypes.ResultBlockResults) error {
	blockEvents := append(block.BeginBlockEvents, block.EndBlockEvents...)

	for _, event := range blockEvents {
		err := tb.pubSub.Publish(TmEvent{
			BlockHeight: block.Height,
			Event:       event,
		})

		if err != nil {
			return err
		}
	}

	for _, txRes := range block.TxsResults {
		for _, event := range txRes.Events {
			err := tb.pubSub.Publish(TmEvent{
				block.Height,
				event,
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (tb *TmEventBus) ListenEvents(ctx context.Context) <-chan error {
	errChan := make(chan error, 2)

	ctx, ctxCancel := context.WithCancel(ctx)
	blockResultChan, blockErr := tb.ListenBlock(ctx)

	go func() {

		for {
			select {
			case block, ok := <-blockResultChan:
				if !ok {
					ctxCancel()
				} else if err := tb.publish(&block); err != nil {
					errChan <- err
					ctxCancel()
				}
			case err := <-blockErr:
				errChan <- err
				ctxCancel()
			case <-ctx.Done():
				tb.pubSub.Close()

				<-tb.pubSub.Done()
				<-tb.listener.Done()

				close(tb.done)
				return
			}
		}
	}()

	return errChan
}

func (tb *TmEventBus) Subscribe(filter func(TmEvent) bool) <-chan TmEvent {
	return tb.pubSub.Subscribe(filter)
}
