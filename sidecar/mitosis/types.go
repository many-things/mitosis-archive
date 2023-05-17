package mitosis

import (
	"context"
	"fmt"
	"time"

	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/sync/errgroup"
)

type Job func(ctx context.Context) error

type EventMgr interface {
	EventTxMgr
	EventListenMgr
}

type eventMgr struct {
	cfg      config.TmConfig
	wallet   tendermint.Wallet
	eventBus *tendermint.TmEventBus
	jobs     []Job
	errGroup *errgroup.Group
	eventCtx context.Context
}

func NewEventMgr(ctx context.Context, cfg config.TmConfig, logger log.Logger) (EventMgr, error) {
	dialURL := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	// TODO: interfaceRegistry
	wallet, err := tendermint.NewWallet(cfg.PrivKey, cfg.Prefix, cfg.ChainID, dialURL, nil)
	if err != nil {
		return nil, err
	}

	// TODO: implement block getter
	fetcher, err := sdkClient.NewClientFromNode(dialURL)
	if err != nil {
		return nil, err
	}

	listener := tendermint.NewBlockListener(ctx, fetcher, time.Second*5)
	pubSub := tendermint.NewPubSub[*tendermint.TmEvent]()
	eventBus := tendermint.NewTmEventBus(listener, pubSub, logger)

	errGroup, eventCtx := errgroup.WithContext(ctx)

	return &eventMgr{
		cfg:      cfg,
		wallet:   wallet,
		eventBus: eventBus,
		jobs:     []Job{},
		errGroup: errGroup,
		eventCtx: eventCtx,
	}, nil
}

func NewWalletFromConfig(cfg config.TmConfig) (tendermint.Wallet, error) {
	dialURL := fmt.Sprintf("%s:%d", cfg.Host, 9090)
	return tendermint.NewWallet(cfg.PrivKey, cfg.Prefix, cfg.ChainID, dialURL, nil)
}
