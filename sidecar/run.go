package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	golog "log"
	"os"
	"os/signal"
	"syscall"

	sdkerrors "cosmossdk.io/errors"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/utils"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/mitosis"
	"github.com/many-things/mitosis/sidecar/storage"
	"github.com/many-things/mitosis/sidecar/tendermint"
	mitotmclient "github.com/many-things/mitosis/sidecar/tendermint/client"
	"github.com/many-things/mitosis/sidecar/tofnd"
	"github.com/many-things/mitosis/sidecar/tofnd/session"
	"github.com/many-things/mitosis/sidecar/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	multisigserver "github.com/many-things/mitosis/x/multisig/server"
	multisigtypes "github.com/many-things/mitosis/x/multisig/types"
	"github.com/tendermint/tendermint/libs/log"
	tmclient "github.com/tendermint/tendermint/rpc/client"
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

func createKeygenHandler(cfg config.SidecarConfig, log log.Logger) func(msg *multisigtypes.Keygen) error {
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

		log.Info(fmt.Sprintf("Start signInit: %X", keygenInit))

		session := mgr.CreateSession(cfg, keygenInit)
		if err := session.StartSession(); err != nil {
			return err
		}

		return nil
	}
}

func createTssSignHandler(cfg config.SidecarConfig, log log.Logger, ctx context.Context) func(msg *multisigtypes.EventSigningStart) error {
	return func(msg *multisigtypes.EventSigningStart) error {
		mgr := session.GetSignMgrInstance(ctx)

		fmt.Println("get createSignHandler", msg)

		signUID := fmt.Sprintf("%s-%d", msg.Chain, msg.SigId)
		var partyUIDs []string
		for _, v := range msg.Participants {
			partyUIDs = append(partyUIDs, v.String())
		}

		msgToSign := sha256.Sum256(msg.MessageToSign)
		signInit := types.SignInit{
			NewSigUid:     signUID,
			KeyUid:        msg.KeyId,
			PartyUids:     partyUIDs,
			MessageToSign: msgToSign[:],
		}

		log.Info(fmt.Sprintf("Start signInit: %X", signInit))

		session := mgr.CreateSession(cfg, signInit)
		if err := session.StartSession(); err != nil {
			return err
		}

		return nil
	}
}

func createSignHandler(cfg config.SidecarConfig, storage storage.Storage, mitoWallet tendermint.Wallet, log log.Logger) func(msg *multisigtypes.EventSigningStart) error {
	return func(msg *multisigtypes.EventSigningStart) error {
		log.Info("signHandler: new sign event observed")
		privKeyBytes, err := storage.GetKey(msg.GetKeyId())
		if err != nil {
			log.Error("signHandler: key not found: %x", err)
			return err
		}

		privKey := secp256k1.PrivKey{Key: privKeyBytes}
		signature, err := privKey.Sign(msg.MessageToSign)
		if err != nil {
			log.Error("signHandler: wrong signature: %x", err)
			return err
		}

		sender, _ := mitoWallet.GetAddress()
		valAddr, err := sdk.ValAddressFromBech32(cfg.TofNConfig.Validator)

		if err != nil {
			log.Error("signHandler: wrong Validator: %x", err)
			return err
		}

		mitoMsg := &multisigserver.MsgSubmitSignature{
			Module:      "sidecar",
			SigID:       exported.SigID(fmt.Sprintf("%s-%d", msg.Chain, msg.SigId)),
			Participant: valAddr,
			Signature:   signature,
			Sender:      sdk.MustAccAddressFromBech32(sender),
		}

		go func() error {
			if err := mitoWallet.BroadcastMsg(mitoMsg); err != nil {
				log.Error("signHandler: fail broadcast: %x", err)
				return err
			}

			return nil
		}()

		return nil
	}
}

func main() {
	homeEnvDir, _ := os.LookupEnv("HOME")
	homeDir := flag.String("home", homeEnvDir+"/.sidecar", "setting for home")
	flag.Parse()

	cfg, err := config.GetConfigFromFile(*homeDir + "/config.yaml")
	if err != nil {
		golog.Fatal(err)
		return
	}
	cfg.Home = *homeDir

	ctx, cancel := context.WithCancel(context.Background())
	eGroup, ctx := errgroup.WithContext(ctx)
	logger := log.NewTMLogger(os.Stdout)

	store := storage.GetStorage(&cfg)
	sdk.GetConfig().SetBech32PrefixForAccount(cfg.MitoConfig.Prefix, "")
	sdk.GetConfig().SetBech32PrefixForValidator(cfg.MitoConfig.ValidatorPrefix, "")

	golog.Println("Set Robust Tendermint Client")
	robustClient := mitotmclient.NewRobustTmClient(func() (tmclient.Client, error) {
		mitoDialURL := fmt.Sprintf("http://%s:%d", cfg.MitoConfig.Host, cfg.MitoConfig.Port)
		golog.Println(mitoDialURL)
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

	golog.Println("Set Tendermint BlockListener")
	listener := tendermint.NewBlockListener(ctx, robustClient, time.Second*5)
	pubSub := tendermint.NewPubSub[*tendermint.TmEvent]()
	eventBus := tendermint.NewTmEventBus(listener, pubSub, logger)

	keygenEventRecv := eventBus.Subscribe(tendermint.Filter[*multisigtypes.Keygen]())
	signEventRecv := eventBus.Subscribe(tendermint.Filter[*multisigtypes.EventSigningStart]())

	wallet, err := mitosis.NewWalletFromConfig(cfg.MitoConfig)
	if err != nil {
		golog.Fatal("wallet generate", err)
		return
	}

	jobs := []mitosis.Job{
		mitosis.CreateTypedJob(keygenEventRecv, createKeygenHandler(cfg, logger), cancel, logger),
		mitosis.CreateTypedJob(signEventRecv, createSignHandler(cfg, store, wallet, logger), cancel, logger),
	}

	utils.ForEach(jobs, func(j mitosis.Job) {
		eGroup.Go(func() error { return j(ctx) })
	})

	go func() {
		<-eventBus.ListenEvents(ctx)
	}()

	go func() {
		golog.Println("Register Tendermint Jobs")
		if err := eGroup.Wait(); err != nil {
			logger.Error(err.Error())
		}
	}()

	//lis, err := net.Listen("tcp", ":9999")
	//if err != nil {
	//	logger.Error(err.Error())
	//	return
	//}
	//
	//grpcServer := grpc.NewServer()
	//go func() {
	//	golog.Println("Run sidecar server")
	//	types.RegisterSidecarServer(grpcServer, &tofnd.TrafficServer{})
	//	if err := grpcServer.Serve(lis); err != nil {
	//		logger.Error(err.Error())
	//	}
	//}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	golog.Println("Sidecar started successfully")
	<-exit

	//grpcServer.GracefulStop()
	_, timeout := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	defer timeout()

	fmt.Println("Gracefully shutdown")
}
