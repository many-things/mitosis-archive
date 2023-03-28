package mito

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	goerr "github.com/go-errors/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/many-things/mitosis/pkg/utils"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/tendermint/tendermint/libs/log"
)

type EventListenMgr interface {
	AddJob(job Job)
	Run() error
}

func (m *eventMgr) AddJob(job Job) {
	m.jobs = append(m.jobs, job)
}

func (m *eventMgr) Run() error {
	utils.ForEach(m.jobs, func(j Job) {
		m.errGroup.Go(func() error { return j(m.eventCtx) })
	})

	if err := m.errGroup.Wait(); err != nil {
		return err
	}

	return nil
}

func Consume[T any](sub <-chan T, handle func(event T)) Job {
	return func(ctx context.Context) error {
		errs := make(chan error, 1)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-errs:
				return err
			case event, ok := <-sub:
				if !ok {
					return nil
				}
				go func() {
					defer recovery(errs)
					handle(event)
				}()
			}
		}
	}
}

func recovery(errChan chan<- error) {
	if r := recover(); r != nil { // nolint: revive
		err := fmt.Errorf("panicked: %s\n%s", r, goerr.Wrap(r, 1).Stack())
		errChan <- err
	}
}

// CreateTypedJob is make handler for applied events. One subscription channel must be matched one job.
func CreateTypedJob[T proto.Message](sub <-chan tendermint.TmEvent, handler func(event T) error, cancel context.CancelFunc, logger log.Logger) Job {
	return func(ctx context.Context) error {
		handleWithLog := func(e tendermint.TmEvent) {
			event := utils.Must(sdk.ParseTypedEvent(e.Event)).(T)
			err := handler(event)
			if err != nil {
				logger.Error(err.Error()) // KeyVal?
			}
		}
		consume := Consume(sub, handleWithLog)
		err := consume(ctx)
		if err != nil {
			cancel()
			return err
		}

		return nil
	}
}
