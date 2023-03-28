package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/tendermint/tendermint/libs/log"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params
	SetParams(ctx sdk.Context, params Params)
	Logger(ctx sdk.Context) log.Logger
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type KeygenKeeper interface {
	RegisterKeygenEvent(ctx sdk.Context, chainID string, keygen *Keygen) (uint64, error)
	RemoveKeygenEvent(ctx sdk.Context, chainID string, id uint64) error
	UpdateKeygenStatus(ctx sdk.Context, chainID string, id uint64, newStatus Keygen_Status) (*Keygen, error)

	QueryKeygen(ctx sdk.Context, chainID string, id uint64) (*Keygen, error)
	QueryKeygenList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *Keygen], *query.PageResponse, error)
}

type PubKeyKeeper interface {
	RegisterPubKey(ctx sdk.Context, chainID string, pubKey *PubKey) error
	RemovePubKey(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress) error

	QueryPubKey(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress) (*PubKey, error)
	QueryPubKeyList(ctx sdk.Context, chainID string, keyID uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *PubKey], *query.PageResponse, error)
}

type SignKeeper interface {
	RegisterSignEvent(ctx sdk.Context, chainID string, sign *Sign) (uint64, error)
	RemoveSignEvent(ctx sdk.Context, chainID string, id uint64) error
	UpdateSignStatus(ctx sdk.Context, chainID string, id uint64, newStatus Sign_Status) (*Sign, error)

	QuerySign(ctx sdk.Context, chainID string, id uint64) (*Sign, error)
	QuerySignList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *Sign], *query.PageResponse, error)
}

type SignatureKeeper interface {
	RegisterSignature(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress, signature Signature) error
	RemoveSignature(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress) error

	QuerySignature(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress) (Signature, error)
	QuerySignatureList(ctx sdk.Context, chainID string, sigID uint64, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, Signature], *query.PageResponse, error)
}
