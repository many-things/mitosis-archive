package sidecar

import (
	"context"
	"fmt"
	golog "log"
	"os"

	sdkerrors "cosmossdk.io/errors"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/utils"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/mitosis"
	"github.com/many-things/mitosis/sidecar/storage"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/many-things/mitosis/sidecar/tofnd"
	"github.com/many-things/mitosis/sidecar/types"
	multisigtypes "github.com/many-things/mitosis/x/multisig/types"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"time"
)

func connectGrpc(host string, port int, timeout time.Duration, logger log.Logger) (*grpc.ClientConn, error) {
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	logger.Info(fmt.Sprintf("initial connection to tofnd server: %s", serverAddr))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return grpc.DialContext(ctx, serverAddr, grpc.WithInsecure(), grpc.WithBlock()) // nolint: staticcheck
}

func createTofNManager(cliCtx sdkclient.Context, config config.SidecarConfig, logger log.Logger, valAddr sdk.ValAddress) *tofnd.Manager {
	conn, err := connectGrpc(config.TofNConfig.Host, config.TofNConfig.Port, config.TofNConfig.DialTimeout, logger)
	if err != nil {
		panic(sdkerrors.Wrapf(err, "failed to create multisig manager"))
	}
	logger.Debug("successful connection to tofnd gRPC server")

	return tofnd.NewManager(types.NewMultisigClient(conn), cliCtx, valAddr, logger, config.TofNConfig.DialTimeout)
}

func createKeygenHandler(store storage.Storage, sigCli types.MultisigClient, _ tendermint.Wallet, logger log.Logger) func(msg *multisigtypes.Keygen) error {
	return func(msg *multisigtypes.Keygen) error {
		if !utils.Any(msg.Participants, store.IsTarget) {
			return nil // Just not targeted.
		}

		keyUID := fmt.Sprintf("%s-%d", msg.Chain, msg.KeyID)

		// TODO: propagate match ctx
		_, err := sigCli.Keygen(context.Background(), &types.KeygenRequest{
			KeyUid:   keyUID,
			PartyUid: store.GetValidator().String(),
		})

		if err != nil {
			logger.Error(err.Error())
			return err
		}

		return nil
	}
}

func createSignHandler(store storage.Storage, sigCli types.MultisigClient, _ tendermint.Wallet, logger log.Logger) func(msg *multisigtypes.Sign) error {
	return func(msg *multisigtypes.Sign) error {
		if !utils.Any(msg.Participants, store.IsTarget) {
			return nil
		}

		pubKey, err := store.GetKey(msg.KeyID)

		if err != nil {
			logger.Error(err.Error())
			return err
		}

		_, err = sigCli.Sign(context.Background(), &types.SignRequest{
			KeyUid:    msg.KeyID,
			MsgToSign: msg.MessageToSign,
			PartyUid:  store.GetValidator().String(),
			PubKey:    []byte(pubKey),
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		// TODO: handle resp
		return nil
	}
}

func run() {
	cfg := config.DefaultSidecarConfig()
	ctx, cancel := context.WithCancel(context.Background())
	eGroup, ctx := errgroup.WithContext(ctx)
	logger := log.NewTMLogger(os.Stdout)

	wallet, err := mitosis.NewWalletFromConfig(cfg.MitoConfig)
	if err != nil {
		panic(err)
	}
	// TODO: make these Rpc robust
	sigDialURL := fmt.Sprintf("%s:%d", cfg.TofNConfig.Host, cfg.TofNConfig.Port)
	sigRPC, err := grpc.Dial(sigDialURL)

	if err != nil {
		panic(fmt.Errorf("cannot dial to tofn network: %w", err))
	}
	sigCli := types.NewMultisigClient(sigRPC)
	store := storage.GetStorage(&cfg)

	// TODO: implement block getter
	mitoDialURL := fmt.Sprintf("%s:%d", cfg.MitoConfig.Host, cfg.MitoConfig.Port)
	fetcher, err := sdkclient.NewClientFromNode(mitoDialURL)
	if err != nil {
		golog.Fatal(err)
	}

	listener := tendermint.NewBlockListener(ctx, fetcher, time.Second*5)
	pubSub := tendermint.NewPubSub[tendermint.TmEvent]()
	eventBus := tendermint.NewTmEventBus(listener, pubSub, logger)

	keygenEventRecv := eventBus.Subscribe(tendermint.Filter[*multisigtypes.Keygen]())
	signEventRecv := eventBus.Subscribe(tendermint.Filter[*multisigtypes.Sign]())

	jobs := []mitosis.Job{
		mitosis.CreateTypedJob(keygenEventRecv, createKeygenHandler(store, sigCli, wallet, logger), cancel, logger),
		mitosis.CreateTypedJob(signEventRecv, createSignHandler(store, sigCli, wallet, logger), cancel, logger),
	}

	utils.ForEach(jobs, func(j mitosis.Job) {
		eGroup.Go(func() error { return j(ctx) })
	})

	if err := eGroup.Wait(); err != nil {
		logger.Error(err.Error())
	}
}
