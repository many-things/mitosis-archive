package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper interface {
	GetParams(ctx sdk.Context) types.Params
	SetParams(ctx sdk.Context, params types.Params)
	Logger(ctx sdk.Context) log.Logger
}

type KeygenKeeper interface {
	RegisterKeygenEvent(ctx sdk.Context, chainId string, keygen *types.Keygen) (uint64, error)
	RemoveKeygenEvent(ctx sdk.Context, chainId string, id uint64) error
	UpdateKeygenStatus(ctx sdk.Context, chainId string, id uint64, newStatus types.Keygen_Status) (*types.Keygen, error)

	QueryKeygen(ctx sdk.Context, chainId string, id uint64) (*types.Keygen, error)
	QueryKeygenList(ctx sdk.Context, chainId string, page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error)
}

type PubKeyKeeper interface {
	RegisterPubKey(ctx sdk.Context, chainId string, pubKey *types.PubKey) error
	RemovePubKey(ctx sdk.Context, chainId string, keyId uint64, participant sdk.ValAddress) error

	QueryPubKey(ctx sdk.Context, chainId string, keyId uint64, participant sdk.ValAddress) (*types.PubKey, error)
	QueryPubKeyList(ctx sdk.Context, chainId string, keyId uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *types.PubKey], *query.PageResponse, error)
}

type SignKeeper interface {
	RegisterSignEvent(ctx sdk.Context, chainId string, sign *types.Sign) (uint64, error)
	RemoveSignEvent(ctx sdk.Context, chainId string, id uint64) error
	UpdateSignStatus(ctx sdk.Context, chainId string, id uint64, newStatus types.Sign_Status) (*types.Sign, error)

	QuerySign(ctx sdk.Context, chainId string, id uint64) (*types.Sign, error)
	QuerySignList(ctx sdk.Context, chainId string, page *query.PageRequest) ([]mitosistype.KV[uint64, *types.Sign], *query.PageResponse, error)
}

type SignatureKeeper interface {
	RegisterSignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress, signature types.Signature) error
	RemoveSignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress) error

	QuerySignature(ctx sdk.Context, chainId string, sigId uint64, participant sdk.ValAddress) (types.Signature, error)
	QuerySignatureList(ctx sdk.Context, chainId string, sigId uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, types.Signature], *query.PageResponse, error)
}
