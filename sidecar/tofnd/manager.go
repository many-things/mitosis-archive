package tofnd

import (
	"context"
	"fmt"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/sidecar/types"
	"github.com/tendermint/tendermint/libs/log"
	"time"
)

type Manager struct {
	client      Client
	ctx         sdkclient.Context
	participant sdk.ValAddress
	logger      log.Logger
	timeout     time.Duration
}

func NewManager(client Client, ctx sdkclient.Context, participant sdk.ValAddress, logger log.Logger, timeout time.Duration) *Manager {
	return &Manager{
		client:      client,
		ctx:         ctx,
		participant: participant,
		logger:      logger,
		timeout:     timeout,
	}
}

func (m Manager) generateKey(KeyUID string) (types.PublicKey, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	res, err := m.client.Keygen(ctx, &types.KeygenRequest{
		KeyUid:   KeyUID,
		PartyUid: m.participant.String(),
	})

	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed generating key")
	}

	switch res.GetKeygenResponse().(type) {
	case *types.KeygenResponse_PubKey:
		return res.GetPubKey(), nil
	case *types.KeygenResponse_Error:
		return nil, fmt.Errorf(res.GetError())
	default:
		panic(fmt.Errorf("unknown TofN keygen response: #{res.GetKeygenResponse()}"))
	}
}

func (m Manager) sign(keyUID string, payloadHash types.Hash, partyUID string, pubKey types.PublicKey) (types.Signature, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	res, err := m.client.Sign(ctx, &types.SignRequest{
		KeyUid:    keyUID,
		MsgToSign: payloadHash,
		PartyUid:  m.participant.String(),
		PubKey:    pubKey,
	})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "fail signing")
	}

	switch res.GetSignResponse().(type) {
	case *types.SignResponse_Signature:
		return res.GetSignature(), nil
	case *types.SignResponse_Error:
		return nil, fmt.Errorf(res.GetError())
	default:
		panic(fmt.Errorf("unknown TofN sign response: #{res.GetSignResponse()}"))
	}
}
