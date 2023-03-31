package client

import (
	"context"
	"strings"

	gorpc "google.golang.org/grpc"
)

type RobustGRPCClient interface {
	gorpc.ClientConnInterface

	resetClient() error
	execute(f func() error) error
	Close() error
}

type robustGRPCClient struct {
	client  gorpc.ClientConn
	healthy chan bool
	factory func() (gorpc.ClientConn, error)
}

func NewRobustGRPCClient(factory func() (gorpc.ClientConn, error)) RobustGRPCClient {
	healthy := make(chan bool, 1)
	healthy <- false

	return &robustGRPCClient{factory: factory, healthy: healthy}
}

func (r *robustGRPCClient) resetClient() error {
	healthy := <-r.healthy

	// If healthy, not necessary to reset
	if !healthy {
		cli, err := r.factory()
		if err != nil {
			r.healthy <- false
			return err
		}

		//TODO: resolve this mutex problem
		r.client = cli //nolint: govet
	}

	r.healthy <- true
	return nil
}

func (r *robustGRPCClient) execute(f func() error) error {
	err := r.resetClient()
	if err != nil {
		return err
	}

	err = f()

	if err != nil && strings.Contains(err.Error(), "post failed") {
		healthy := <-r.healthy
		if healthy { // force reset client
			_ = r.client.Close()
		}
		return err
	}

	return err // if err is nil, doesnt matter
}

func (r *robustGRPCClient) Close() error {
	return r.execute(func() error { return r.client.Close() })
}

func (r *robustGRPCClient) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...gorpc.CallOption) error {
	return r.execute(func() error {
		return r.client.Invoke(ctx, method, args, reply, opts...)
	})
}

func (r *robustGRPCClient) NewStream(ctx context.Context, desc *gorpc.StreamDesc, method string, opts ...gorpc.CallOption) (gorpc.ClientStream, error) {
	var stream gorpc.ClientStream
	var err error

	err = r.execute(func() error {
		stream, err = r.client.NewStream(ctx, desc, method, opts...)
		return err
	})

	if err != nil {
		return nil, err
	}

	return stream, nil
}
