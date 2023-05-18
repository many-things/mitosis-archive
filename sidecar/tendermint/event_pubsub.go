package tendermint

import (
	"context"

	"github.com/gogo/protobuf/proto"
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
	pubSub   PubSub[*TmEvent]
}

func NewTmEventBus(listener BlockListener, pubSub PubSub[*TmEvent], logger log.Logger) *TmEventBus {
	return &TmEventBus{
		listener: listener,
		pubSub:   pubSub,
		logger:   logger,
	}
}

// ListenBlock listen The Latest Block Height and get BlockResult from tendermint client
func (tb *TmEventBus) ListenBlock(ctx context.Context) (<-chan coretypes.ResultBlockResults, <-chan error) {
	blockResultChan := make(chan coretypes.ResultBlockResults)
	errChan := make(chan error, 1)

	go func() {
		newBlockHeightChan, heightErrChan := tb.listener.NewBlockWatcher()

		for {
			select {
			case blockHeight := <-newBlockHeightChan:
				block, err := tb.listener.GetBlockResult(&blockHeight)
				if err != nil {
					errChan <- err
					return
				}

				//if len(block.TxsResults) > 0 {
				//	fmt.Println("========= TxResults =========")
				//	for _, item := range block.TxsResults {
				//		fmt.Println(item)
				//	}
				//	fmt.Println("========= TxResults =========")
				//}

				//if len(block.BeginBlockEvents) > 0 {
				//	fmt.Println("========= BeginBlockEvents =========")
				//	for _, item := range block.BeginBlockEvents {
				//		fmt.Println(item)
				//	}
				//	fmt.Println("========= BeginBlockEvents =========")
				//}

				//if len(block.EndBlockEvents) > 0 {
				//	fmt.Println("========= EndBlockEvents =========")
				//	for _, item := range block.EndBlockEvents {
				//		fmt.Println(item)
				//	}
				//	fmt.Println("========= EndBlockEvents =========")
				//}

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

// publish iterate blockEvent and send to pubSub.Publish
func (tb *TmEventBus) publish(block *coretypes.ResultBlockResults) error {
	//blockEvents := append(block.BeginBlockEvents, block.EndBlockEvents...) // nolint: gocritic
	//for _, event := range block.BeginBlockEvents {
	//	err := tb.pubSub.Publish(&TmEvent{
	//		BlockHeight: block.Height,
	//		Event:       event,
	//	})
	//
	//	if err != nil {
	//		return err
	//	}
	//}

	for _, event := range block.EndBlockEvents {
		err := tb.pubSub.Publish(&TmEvent{
			BlockHeight: block.Height,
			Event:       event,
		})

		if err != nil {
			return err
		}
	}

	for _, txRes := range block.TxsResults {
		for _, event := range txRes.Events {
			err := tb.pubSub.Publish(&TmEvent{
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

	tb.pubSub.Run()

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
				tb.listener.Close()

				return
			}
		}
	}()

	return errChan
}

func (tb *TmEventBus) Subscribe(filter func(*TmEvent) bool) <-chan *TmEvent {
	return tb.pubSub.Subscribe(filter)
}

func Filter[T proto.Message]() func(e *TmEvent) bool {
	return func(e *TmEvent) bool {
		if e == nil {
			return false
		}

		return e.Event.Type == proto.MessageName(*new(T)) // nolint: gocritic
	}
}
