package session

import (
	"sync"

	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/types"
)

type KeygenSession interface {
}

type keygenSessionMgr struct {
	sessions map[string]KeygenSession
}

type keygenSession struct {
	Config   config.SidecarConfig
	msg      types.KeygenInit
	sessions map[string]types.SidecarClient
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

func (m *keygenSessionMgr) CreateSession(cfg config.SidecarConfig, msg types.KeygenInit) KeygenSession {
	m.sessions[msg.NewKeyUid] = &keygenSession{
		Config:   cfg,
		msg:      msg,
		sessions: map[string]types.SidecarClient{},
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

func (m *keygenSessionMgr) ReleaseSession(key string) {
	delete(m.sessions, key)
}
