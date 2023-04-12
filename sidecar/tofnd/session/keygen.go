package session

import (
	"context"
	"fmt"
	"sync"

	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/types"
	"google.golang.org/grpc"
)

type KeygenSession interface {
	StartSession() error
	CloseSession() error
	BroadcastMsg(msg types.TrafficIn) error
	ConsumeMsg(msg types.TrafficIn) error
	IsRunning() bool
}

type keygenSessionMgr struct {
	sessions map[string]KeygenSession
}

type keygenSession struct {
	config    config.TofNConfig
	msg       types.KeygenInit
	sessions  map[string]*grpc.ClientConn
	isRunning bool
}

var mgrInstance *keygenSessionMgr
var lock = &sync.Mutex{}

func GetInstance() *keygenSessionMgr { //nolint: revive
	if mgrInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		// Is still nil after get Lock
		if mgrInstance == nil {
			mgrInstance = &keygenSessionMgr{}
		}
	}

	return mgrInstance
}

func (m *keygenSessionMgr) CreateSession(cfg config.TofNConfig, msg types.KeygenInit) KeygenSession {
	m.sessions[msg.NewKeyUid] = &keygenSession{
		config:   cfg,
		msg:      msg,
		sessions: map[string]*grpc.ClientConn{},
	}

	return m.sessions[msg.NewKeyUid]
}

func (m *keygenSessionMgr) GetSession(key string) (KeygenSession, bool) {
	sess, ok := m.sessions[key]
	if !ok {
		return nil, false
	}

	return sess, true
}

func (m *keygenSessionMgr) Consume(msg types.ShareKeygenRequest) error {
	sess, ok := m.GetSession(msg.NewKeyUid)
	if !ok {
		return fmt.Errorf("key_uid not found: %s", msg.NewKeyUid)
	}

	if !sess.IsRunning() {
		return fmt.Errorf("key_uid not initialized: %s", msg.NewKeyUid)
	}

	return sess.ConsumeMsg(*msg.Traffic)
}

func (m *keygenSessionMgr) ReleaseSession(key string) error {
	if sess, ok := m.GetSession(key); ok {
		if sess.IsRunning() {
			err := sess.CloseSession()
			if err != nil {
				return err
			}
		}
		delete(m.sessions, key)
	}

	return nil
}

func (s *keygenSession) StartSession() error {
	// convert node list to HashMap
	nodeMap := map[string]string{}
	for _, v := range s.config.Nodes {
		nodeMap[v.Validator] = v.Host
	}

	for _, p := range s.msg.PartyUids {
		if p == s.config.Validator {
			continue
		}

		dialURL := nodeMap[p]
		conn, err := grpc.Dial(dialURL, grpc.WithInsecure()) // nolint: staticcheck
		if err != nil {
			return err
		}

		s.sessions[p] = conn
	}

	s.isRunning = true

	return nil
}

func (s *keygenSession) CloseSession() error {
	// TODO: apply mutex
	for _, p := range s.sessions {
		err := p.Close()

		if err != nil {
			return err
		}
	}

	s.isRunning = false
	return nil
}

func (s *keygenSession) BroadcastMsg(msg types.TrafficIn) error {
	for _, v := range s.sessions {
		serv := types.NewSidecarClient(v)
		_, err := serv.ShareKeygenTraffic(context.Background(), &types.ShareKeygenRequest{
			NewKeyUid: s.msg.GetNewKeyUid(),
			Traffic:   &msg,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *keygenSession) ConsumeMsg(_ types.TrafficIn) error {
	panic("implement me")
}

func (s *keygenSession) IsRunning() bool {
	return s.isRunning
}
