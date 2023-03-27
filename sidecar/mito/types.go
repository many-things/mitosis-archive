package mito

import (
	"context"
	"fmt"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/sync/errgroup"
	"time"
)

type Job func(ctx context.Context) error

type EventMgr interface {
	EventTxMgr
	EventListenMgr
}

type eventMgr struct {
	cfg      config.TendermintConfig
	wallet   tendermint.Wallet
	eventBus *tendermint.TmEventBus
	jobs     []Job
	errGroup *errgroup.Group
	eventCtx context.Context
}

func NewEventMgr(ctx context.Context, cfg config.TendermintConfig, logger log.Logger) (EventMgr, error) {
	dialUrl := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	// TODO: interfaceRegistry
	wallet, err := tendermint.NewWallet(cfg.PrivKey, cfg.Prefix, cfg.ChainId, dialUrl, nil)
	if err != nil {
		return nil, err
	}

	// TODO: implement block getter
	fetcher, err := sdkClient.NewClientFromNode(dialUrl)

	listener := tendermint.NewBlockListener(ctx, fetcher, time.Second*5)
	pubSub := tendermint.NewPubSub[tendermint.TmEvent]()
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
