package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtype "github.com/cosmos/cosmos-sdk/x/auth/types"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/tendermint/tendermint/libs/log"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params
	SetParams(ctx sdk.Context, params Params)
	Logger(ctx sdk.Context) log.Logger
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtype.AccountI
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
	DeletePubKey(ctx sdk.Context, chainID string, keyID uint64) error
	AddParticipantPubKey(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress, publicKey exported.PublicKey) error
	RemoveParticipantPubKey(ctx sdk.Context, chainID string, keyID uint64, participant sdk.ValAddress) error
	HasPubKey(ctx sdk.Context, chainID string, keyID uint64) bool

	QueryPubKey(ctx sdk.Context, chainID string, keyID uint64) (*PubKey, error)
	QueryPubKeyList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[sdk.ValAddress, *PubKey], *query.PageResponse, error)
}

type SignKeeper interface {
	RegisterSignEvent(ctx sdk.Context, chainID string, sign *exported.Sign) (uint64, error)
	RemoveSignEvent(ctx sdk.Context, chainID string, id uint64) error
	UpdateSignStatus(ctx sdk.Context, chainID string, id uint64, newStatus exported.Sign_Status) (*exported.Sign, error)

	QuerySign(ctx sdk.Context, chainID string, id uint64) (*exported.Sign, error)
	QuerySignList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.Sign], *query.PageResponse, error)
}

type SignatureKeeper interface {
	RegisterSignature(ctx sdk.Context, chainID string, signSignature *exported.SignSignature) error
	RemoveSignature(ctx sdk.Context, chainID string, sigID uint64) error
	HasSignature(ctx sdk.Context, chainID string, sigID uint64) bool
	AddParticipantSignature(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress, signature exported.Signature) error
	RemoveParticipantSignature(ctx sdk.Context, chainID string, sigID uint64, participant sdk.ValAddress) error

	QuerySignature(ctx sdk.Context, chainID string, sigID uint64) (*exported.SignSignature, error)
	QuerySignatureList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *exported.SignSignature], *query.PageResponse, error)
}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context, chains []byte) (*GenesisState, error)
	ImportGenesis(ctx sdk.Context, genState *GenesisState) error
}

type ChainKeeper interface {
	QueryChains(ctx sdk.Context, pageReq *query.PageRequest) ([]mitosistype.KV[string, byte], *query.PageResponse, error)
}
