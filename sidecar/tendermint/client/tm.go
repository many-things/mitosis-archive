package client

import (
	"context"
	"strings"

	"github.com/tendermint/tendermint/rpc/client"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type RobustTmClient struct {
	client  client.Client
	healthy chan bool
	factory func() (client.Client, error)
}

func (r *RobustTmClient) BlockchainInfo(
	ctx context.Context,
	minHeight,
	maxHeight int64,
) (*coretypes.ResultBlockchainInfo, error) {
	var result *coretypes.ResultBlockchainInfo
	var err error

	err = r.execute(func() error {
		result, err = r.client.BlockchainInfo(ctx, minHeight, maxHeight)
		return err
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RobustTmClient) BlockResults(ctx context.Context, blockHeight *int64) (*coretypes.ResultBlockResults, error) {
	var result *coretypes.ResultBlockResults
	var err error

	err = r.execute(func() error {
		result, err = r.client.BlockResults(ctx, blockHeight)
		return err
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewRobustTmClient return a new RobustTmClient
func NewRobustTmClient(factory func() (client.Client, error)) *RobustTmClient {
	healthy := make(chan bool, 1)
	healthy <- false

	return &RobustTmClient{factory: factory, healthy: healthy}
}

// resetClient reset client of RobustTmClient instance.
func (r *RobustTmClient) resetClient() error {
	healthy := <-r.healthy

	// If healthy, not necessary to reset
	if !healthy {
		cli, err := r.factory()
		if err != nil {
			r.healthy <- false
			return err
		}
		r.client = cli
	}

	r.healthy <- true
	return nil
}

// execute runs given function with insure client conn
func (r *RobustTmClient) execute(f func() error) error {
	err := r.resetClient()
	if err != nil {
		return err
	}

	err = f()

	if err != nil && strings.Contains(err.Error(), "post failed") {
		healthy := <-r.healthy
		if healthy { // force reset client
			_ = r.client.Stop()
		}
		return err
	}

	return err // if err is nil, doesnt matter
}

// Stop stops client
func (r *RobustTmClient) Stop() error {
	return r.execute(func() error { return r.client.Stop() })
}
