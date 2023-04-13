package session

import (
	"github.com/many-things/mitosis/sidecar/types"
	"google.golang.org/grpc"
)

type GG20Stream interface {
	Send(*types.MessageIn) error
	Recv() (*types.MessageOut, error)
	grpc.ClientStream
}

type GG20StreamSession struct {
	conn   *grpc.ClientConn
	stream GG20Stream
}
