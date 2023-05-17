package session

import (
	"context"
	"fmt"
	"log"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/sidecar/config"
	"github.com/many-things/mitosis/sidecar/mitosis"
	"github.com/many-things/mitosis/sidecar/tendermint"
	"github.com/many-things/mitosis/sidecar/types"
	multisigserver "github.com/many-things/mitosis/x/multisig/server"
	"google.golang.org/grpc"
)

type SignSession interface {
	StartSession() error
	CloseSession() error
	BroadcastMsg(msg types.TrafficIn) error
	ConsumeMsg(msg types.TrafficIn) error
	IsRunning() bool
}

type signSessionMgr struct {
	ctx      context.Context
	sessions map[string]SignSession
}

type signSession struct {
	ctx       context.Context
	config    config.TofNConfig
	msg       types.SignInit
	sessions  map[string]*grpc.ClientConn
	stream    *GG20StreamSession
	wallet    tendermint.Wallet
	isRunning bool
}

var signMgrInstance *signSessionMgr
var signLock = &sync.Mutex{}

func GetSignMgrInstance(ctx context.Context) *signSessionMgr { //nolint: revive
	if signMgrInstance == nil {
		signLock.Lock()
		defer signLock.Unlock()

		// Is still nil after get Lock
		if signMgrInstance == nil {
			signMgrInstance = &signSessionMgr{
				ctx:      ctx,
				sessions: map[string]SignSession{},
			}
		}
	}

	return signMgrInstance
}

func (m *signSessionMgr) CreateSession(cfg config.SidecarConfig, msg types.SignInit) SignSession {
	wallet, err := mitosis.NewWalletFromConfig(cfg.MitoConfig)
	if err != nil {
		return nil
	}

	m.sessions[msg.NewSigUid] = &signSession{
		ctx:      m.ctx,
		config:   cfg.TofNConfig,
		msg:      msg,
		sessions: map[string]*grpc.ClientConn{},
		stream:   nil,
		wallet:   wallet,
	}

	return m.sessions[msg.NewSigUid]
}

func (m *signSessionMgr) GetSession(key string) (SignSession, bool) {
	sess, ok := m.sessions[key]
	if !ok {
		return nil, false
	}

	return sess, true
}

func (m *signSessionMgr) Consume(msg types.ShareKeygenRequest) error {
	sess, ok := m.GetSession(msg.NewKeyUid)
	if !ok {
		return fmt.Errorf("key_uid not found: %s", msg.NewKeyUid)
	}

	if !sess.IsRunning() {
		return fmt.Errorf("key_uid not initialized: %s", msg.NewKeyUid)
	}

	return sess.ConsumeMsg(*msg.Traffic)
}

func (m *signSessionMgr) ReleaseSession(key string) error {
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

func (s *signSession) StartSession() error {
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

	s.spawnReceiver()

	return nil
}

func (s *signSession) spawnReceiver() error {
	// TODO: handle go routine errs

	go func() error {
		defer func() { s.isRunning = false }()
		dialURL := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
		conn, err := grpc.Dial(dialURL, grpc.WithInsecure())
		if err != nil {
			return err
		}

		cli := types.NewGG20Client(conn)
		stream, err := cli.Sign(s.ctx)
		if err != nil {
			return err
		}

		fmt.Println("SignInit", s.msg)

		s.stream = &GG20StreamSession{
			conn:   conn,
			stream: stream,
		}
		s.isRunning = true

		if err := stream.Send(&types.MessageIn{Data: &types.MessageIn_SignInit{SignInit: &s.msg}}); err != nil {
			fmt.Println("stream.Send", err)
			return err
		}

		for {
			res, err := s.stream.stream.Recv()
			if err != nil {
				log.Fatal(err)
				return err
			}

			switch v := res.GetData().(type) {
			case *types.MessageOut_Traffic:
				fmt.Println(v)
				bMsg := types.TrafficIn{
					FromPartyUid: s.config.Validator,
					Payload:      v.Traffic.Payload,
					IsBroadcast:  v.Traffic.IsBroadcast,
				}
				if err := s.BroadcastMsg(bMsg); err != nil {
					log.Fatal(err)
				}
			case *types.MessageOut_SignResult_:
				switch k := v.SignResult.GetSignResultData().(type) {
				case *types.MessageOut_SignResult_Signature:
					addr, err := s.wallet.GetAddress()
					if err != nil {
						log.Fatal(err)
						return err
					}
					accAddress, err := sdk.AccAddressFromBech32(addr)
					if err != nil {
						log.Fatal(err)
						return err
					}

					msg := &multisigserver.MsgSubmitSignature{
						Module:      "sidecar",
						Participant: sdk.ValAddress(s.config.Validator),
						Signature:   k.Signature,
						Sender:      accAddress,
					}

					log.Println("SignResult", msg)
					err = s.wallet.BroadcastMsg(msg)
					if err != nil {
						// TODO: handle more well
						log.Fatal(err)
					}

					s.stream.stream.CloseSend()
					s.stream.conn.Close()
					return nil
				case *types.MessageOut_SignResult_Criminals:
					// TODO: handle jail function
					log.Fatal(fmt.Errorf("criminal"))
					return nil
				}
			case *types.MessageOut_NeedRecover:
				fmt.Println(v)
				return nil
			}
		}
	}()

	return nil
}

func (s *signSession) CloseSession() error {
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

func (s *signSession) BroadcastMsg(msg types.TrafficIn) error {
	// to self
	s.stream.stream.Send(&types.MessageIn{&types.MessageIn_Traffic{Traffic: &msg}})

	for _, v := range s.sessions {
		serv := types.NewSidecarClient(v)
		_, err := serv.ShareSignTraffic(context.Background(), &types.ShareSignRequest{
			NewSigUid: s.msg.GetNewSigUid(),
			Traffic:   &msg,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *signSession) ConsumeMsg(msg types.TrafficIn) error {
	if s.IsRunning() {
		return s.stream.stream.Send(&types.MessageIn{Data: &types.MessageIn_Traffic{Traffic: &msg}})
	}

	return nil
}

func (s *signSession) IsRunning() bool {
	return s.isRunning
}
