package sidecar

import (
	"context"
	"crypto/sha256"
	"fmt"
	golog "log"
	"net"
	"os"

	mitotmclient "github.com/many-things/mitosis/sidecar/tendermint/client"
	tmclient "github.com/tendermint/tendermint/rpc/client"

	sdkerrors "cosmossdk.io/errors"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/utils"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/mitosis"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/many-things/mitosis/sidecar/tofnd"
	"github.com/many-things/mitosis/sidecar/tofnd/session"
	"github.com/many-things/mitosis/sidecar/types"
	multisigexport "github.com/many-things/mitosis/x/multisig/exported"
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

func createKeygenHandler(cfg config.SidecarConfig, _ log.Logger) func(msg *multisigtypes.Keygen) error {
	return func(msg *multisigtypes.Keygen) error {
		mgr := session.GetKeygenMgrInstance()

		keyUID := fmt.Sprintf("%s-%d", msg.Chain, msg.KeyID)
		var partyUIDs []string
		var partySharesCounts []uint32
		for _, v := range msg.Participants {
			partyUIDs = append(partyUIDs, v.Address.String())
			partySharesCounts = append(partySharesCounts, v.Share)
		}
		myPartyIndex := utils.IndexOf(cfg.TofNConfig.Validator, partyUIDs)
		if myPartyIndex < 0 {
			return fmt.Errorf("no available validator in participant")
		}

		keygenInit := types.KeygenInit{
			NewKeyUid:        keyUID,
			PartyUids:        partyUIDs,
			PartyShareCounts: partySharesCounts,
			MyPartyIndex:     uint32(myPartyIndex),
			Threshold:        uint32(msg.Threshold),
		}

		session := mgr.CreateSession(cfg, keygenInit)
		if err := session.StartSession(); err != nil {
			return err
		}

		// if err := storage.SaveKey(keyUID, ""); err != nil {
		//	return err
		// }

		return nil
	}
}

func createSignHandler(cfg config.SidecarConfig, _ log.Logger) func(msg *multisigexport.Sign) error {
	return func(msg *multisigexport.Sign) error {
		mgr := session.GetSignMgrInstance()

		signUID := fmt.Sprintf("%s-%d", msg.Chain, msg.SigID)
		var partyUIDs []string
		for _, v := range msg.Participants {
			partyUIDs = append(partyUIDs, v.String())
		}

		msgToSign := sha256.Sum256(msg.MessageToSign)
		signInit := types.SignInit{
			NewSigUid:     signUID,
			KeyUid:        msg.KeyID,
			PartyUids:     partyUIDs,
			MessageToSign: msgToSign[:],
		}

		session := mgr.CreateSession(cfg, signInit)
		if err := session.StartSession(); err != nil {
			return err
		}

		return nil
	}
}

func run() {
	cfg := config.DefaultSidecarConfig()
	ctx, cancel := context.WithCancel(context.Background())
	eGroup, ctx := errgroup.WithContext(ctx)
	logger := log.NewTMLogger(os.Stdout)

	// TODO: apply Storage on use
	// store := storage.GetStorage(&cfg)

	// TODO: implement block getter

	robustClient := mitotmclient.NewRobustTmClient(func() (tmclient.Client, error) {
		mitoDialURL := fmt.Sprintf("%s:%d", cfg.MitoConfig.Host, cfg.MitoConfig.Port)
		fetcher, err := sdkclient.NewClientFromNode(mitoDialURL)
		if err != nil {
			golog.Fatal(err)
			return nil, err
		}

		err = fetcher.Start()
		if err != nil {
			golog.Fatal(err)
			return nil, err
		}

		return fetcher, err
	})

	listener := tendermint.NewBlockListener(ctx, robustClient, time.Second*5)
	pubSub := tendermint.NewPubSub[tendermint.TmEvent]()
	eventBus := tendermint.NewTmEventBus(listener, pubSub, logger)

	keygenEventRecv := eventBus.Subscribe(tendermint.Filter[*multisigtypes.Keygen]())
	signEventRecv := eventBus.Subscribe(tendermint.Filter[*multisigexport.Sign]())

	jobs := []mitosis.Job{
		mitosis.CreateTypedJob(keygenEventRecv, createKeygenHandler(cfg, logger), cancel, logger),
		mitosis.CreateTypedJob(signEventRecv, createSignHandler(cfg, logger), cancel, logger),
	}

	utils.ForEach(jobs, func(j mitosis.Job) {
		eGroup.Go(func() error { return j(ctx) })
	})

	if err := eGroup.Wait(); err != nil {
		logger.Error(err.Error())
	}

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	grpcServer := grpc.NewServer()
	types.RegisterSidecarServer(grpcServer, &tofnd.TrafficServer{})
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error(err.Error())
	}
}
