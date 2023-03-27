package sidecar

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	goerr "github.com/go-errors/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/mito"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/many-things/mitosis/sidecar/tofnd"
	"github.com/many-things/mitosis/sidecar/types"
	"github.com/many-things/mitosis/sidecar/utils"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc"
	"time"
)

func connectGrpc(host string, port string, timeout time.Duration, logger log.Logger) (*grpc.ClientConn, error) {
	serverAddr := fmt.Sprintf("#{host}:#{port}")
	logger.Info(fmt.Sprintf("initial connection to tofnd server: #{serverAddr}"))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return grpc.DialContext(ctx, serverAddr, grpc.WithInsecure(), grpc.WithBlock())
}

func createTofNManager(cliCtx client.Context, config config.SidecarConfig, logger log.Logger, valAddr sdk.ValAddress) *tofnd.Manager {
	conn, err := connectGrpc(config.TofNConfig.Host, config.TofNConfig.Port, config.TofNConfig.DialTimeout, logger)
	if err != nil {
		panic(sdkerrors.Wrapf(err, "failed to create multisig manager"))
	}
	logger.Debug("successful connection to tofnd gRPC server")

	return tofnd.NewManager(types.NewMultisigClient(conn), cliCtx, valAddr, logger, config.TofNConfig.DialTimeout)
}

func run() {

}

func Consume[T any](sub <-chan T, handle func(event T)) mito.Job {
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
	if r := recover(); r != nil {
		err := fmt.Errorf("panicked: %s\n%s", r, goerr.Wrap(r, 1).Stack())
		errChan <- err
	}
}

func createTypedJob[T proto.Message](sub <-chan tendermint.TmEvent, handler func(event T) error, cancel context.CancelFunc, logger log.Logger) mito.Job {
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
