package tofnd

import (
	"fmt"

	"github.com/many-things/mitosis/sidecar/tofnd/session"
	"github.com/many-things/mitosis/sidecar/types"
)

type TrafficServer struct {
	types.SidecarServer
}

func (s *TrafficServer) ShareKeygenRequest(msg types.ShareKeygenRequest) (*types.ShareKeygenResponse, error) {
	mgr := session.GetKeygenMgrInstance()

	session, ok := mgr.GetSession(msg.NewKeyUid)
	if !ok {
		return nil, fmt.Errorf("cannot find keygen: %s", msg.NewKeyUid)
	}

	if !session.IsRunning() {
		return nil, fmt.Errorf("keygen session not started: %s", msg.NewKeyUid)
	}

	if err := session.ConsumeMsg(*msg.Traffic); err != nil {
		return nil, err
	}
	return &types.ShareKeygenResponse{}, nil
}

func (s *TrafficServer) ShareSignTraffic(msg types.ShareSignRequest) (*types.ShareSignResponse, error) {
	mgr := session.GetSignMgrInstance()

	session, ok := mgr.GetSession(msg.NewSigUid)
	if !ok {
		return nil, fmt.Errorf("cannot find sign: %s", msg.NewSigUid)
	}

	if !session.IsRunning() {
		return nil, fmt.Errorf("sign session not started: %s", msg.NewSigUid)
	}

	if err := session.ConsumeMsg(*msg.Traffic); err != nil {
		return nil, err
	}

	return &types.ShareSignResponse{}, nil
}
