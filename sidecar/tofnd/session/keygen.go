package session

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/mitosis"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/many-things/mitosis/sidecar/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	multisigserver "github.com/many-things/mitosis/x/multisig/server"
	"github.com/pkg/errors"
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
	stream    *GG20StreamSession
	wallet    tendermint.Wallet
	isRunning bool
}

var mgrInstance *keygenSessionMgr
var lock = &sync.Mutex{}

func GetKeygenMgrInstance() *keygenSessionMgr { //nolint: revive
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
	wallet, err := mitosis.NewWalletFromConfig(cfg.MitoConfig)
	if err != nil {
		return nil
	}

	m.sessions[msg.NewKeyUid] = &keygenSession{
		config:   cfg.TofNConfig,
		msg:      msg,
		sessions: map[string]*grpc.ClientConn{},
		stream:   nil,
		wallet:   wallet,
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

	dialURL := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	conn, err := grpc.Dial(dialURL, grpc.WithInsecure()) // nolint: staticcheck
	if err != nil {
		return err
	}

	cli := types.NewGG20Client(conn)
	stream, err := cli.Keygen(context.Background())
	if err != nil {
		return err
	}

	s.stream = &GG20StreamSession{
		conn:   conn,
		stream: stream,
	}
	s.isRunning = true

	return nil
}

func (s *keygenSession) spawnReceiver() error {
	// TODO: handle go routine errs

	go func() {
		for {
			res, err := s.stream.stream.Recv()
			if errors.Is(err, io.EOF) {
				return
			}

			switch v := res.GetData().(type) {
			case *types.MessageOut_Traffic:
				bMsg := types.TrafficIn{
					FromPartyUid: s.config.Validator,
					Payload:      v.Traffic.Payload,
					IsBroadcast:  v.Traffic.IsBroadcast,
				}
				if err := s.BroadcastMsg(bMsg); err != nil {
					log.Fatal(err)
				}
				return
			case *types.MessageOut_KeygenResult_:
				switch k := v.KeygenResult.GetKeygenResultData().(type) {
				case *types.MessageOut_KeygenResult_Data:
					err := s.wallet.BroadcastMsg(&multisigserver.MsgSubmitPubkey{
						Module:      "sidecar",
						KeyID:       exported.KeyID(s.msg.NewKeyUid),
						Participant: sdk.ValAddress(s.config.Validator),
						PubKey:      k.Data.PubKey,
					})
					if err != nil {
						// TODO: handle more well
						log.Fatal(err)
					}

					return
				case *types.MessageOut_KeygenResult_Criminals:
					// TODO: handle jail function
					log.Fatal(fmt.Errorf("criminal"))
				}
			case *types.MessageOut_NeedRecover:
				fmt.Println(v)
			}
		}
	}()

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

func (s *keygenSession) ConsumeMsg(msg types.TrafficIn) error {
	if s.IsRunning() {
		return s.stream.stream.Send(&types.MessageIn{Data: &types.MessageIn_Traffic{Traffic: &msg}})
	}

	return nil
}

func (s *keygenSession) IsRunning() bool {
	return s.isRunning
}
