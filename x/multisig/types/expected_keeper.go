package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper interface {
	GetParams(ctx sdk.Context) Params
	SetParams(ctx sdk.Context, params Params)
	Logger(ctx sdk.Context) log.Logger

	KeygenKeeper
	PubKeyKeeper
	SignKeeper
	SignatureKeeper
}

type KeygenKeeper interface {
	RegisterKeygenEvent(ctx sdk.Context, chainId string, keygen *Keygen) (uint64, error)
	RemoveKeygenEvent(ctx sdk.Context, chainId string, id uint64) error
	UpdateKeygenStatus(ctx sdk.Context, chainId string, id uint64, newStatus Keygen_Status) (*Keygen, error)

	QueryKeygen(ctx sdk.Context, chainId string, id uint64) (*Keygen, error)
	QueryKeygenList(ctx sdk.Context, chainId string, page *query.PageRequest) ([]mitosistype.KV[uint64, *Keygen], *query.PageResponse, error)
}

type PubKeyKeeper interface {
	RegisterPubKey(ctx sdk.Context, chainId string, pubKey *PubKey) error
	RemovePubKey(ctx sdk.Context, chainId string, keyId uint64, participant sdk.ValAddress) error

	QueryPubKey(ctx sdk.Context, chainId string, keyId uint64, participant sdk.ValAddress) (*PubKey, error)
	QueryPubKeyList(ctx sdk.Context, chainId string, keyId uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *PubKey], *query.PageResponse, error)
}

type SignKeeper interface {
	RegisterSignEvent(ctx sdk.Context, chainId string, sign *Sign) (uint64, error)
	RemoveSignEvent(ctx sdk.Context, chainId string, id uint64) error
	UpdateSignStatus(ctx sdk.Context, chainId string, id uint64, newStatus Sign_Status) (*Sign, error)

	QuerySign(ctx sdk.Context, chainId string, id uint64) (*Sign, error)
	QuerySignList(ctx sdk.Context, chainId string, page *query.PageRequest) ([]mitosistype.KV[uint64, *Sign], *query.PageResponse, error)
}

type SignatureKeeper interface {
	RegisterSignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress, signature Signature) error
	RemoveSignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress) error

	QuerySignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress) (Signature, error)
	QuerySignatureList(ctx sdk.Context, chainId string, sigId uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, Signature], *query.PageResponse, error)
}
