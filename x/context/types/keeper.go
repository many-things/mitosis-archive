package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	evttypes "github.com/many-things/mitosis/x/event/types"
	"github.com/many-things/mitosis/x/multisig/exported"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params

	SetParams(ctx sdk.Context, params Params)
}

type VaultKeeper interface {
	RegisterVault(ctx sdk.Context, chain string, vault string) error

	ClearVault(ctx sdk.Context, chain string) error

	QueryVault(ctx sdk.Context, chain string) (string, error)

	QueryVaults(ctx sdk.Context, page *query.PageRequest) ([]mitotypes.KV[string, string], *query.PageResponse, error)
}

type OperationKeeper interface {
	InitOperation(ctx sdk.Context, chain string, poll *evttypes.Poll) (uint64, error)

	StartSignOperation(ctx sdk.Context, id, sigID uint64, pubkey exported.PublicKey) error

	FinishSignOperation(ctx sdk.Context, id uint64) error

	FinishOperation(ctx sdk.Context, id uint64, poll *evttypes.Poll) error

	QueryOperation(ctx sdk.Context, id uint64) (*Operation, error)

	QueryOperations(ctx sdk.Context, pageReq *query.PageRequest) ([]*Operation, *query.PageResponse, error)

	QueryOperationsByStatus(ctx sdk.Context, status Operation_Status, pageReq *query.PageRequest) ([]*Operation, *query.PageResponse, error)

	QueryOperationByHash(ctx sdk.Context, chain string, hash []byte) (*Operation, error)
}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context, chains []string) (genesis *GenesisState, err error)

	ImportGenesis(ctx sdk.Context, genesis *GenesisState) error
}
