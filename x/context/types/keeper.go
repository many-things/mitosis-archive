package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	evttypes "github.com/many-things/mitosis/x/event/types"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params

	SetParams(ctx sdk.Context, params Params)
}

type SignerKeeper interface {
	SetReadyToSigner(ctx sdk.Context, chain string) error

	RegisterCosmosSigner(ctx sdk.Context, chain string, pubKey []byte, accountNumber uint64) error

	RegisterEVMSigner(ctx sdk.Context, chain string, pubKey []byte) error
}

type OperationKeeper interface {
	InitOperation(ctx sdk.Context, chain string, poll *evttypes.Poll) (uint64, error)

	StartSignOperation(ctx sdk.Context, id, sigID uint64) error

	FinishSignOperation(ctx sdk.Context, id uint64, signature []byte) error

	FinishOperation(ctx sdk.Context, id uint64, poll *evttypes.Poll) error

	QueryOperation(ctx sdk.Context, id uint64) (*Operation, error)

	QueryOperations(ctx sdk.Context, pageReq *query.PageRequest) ([]*Operation, *query.PageResponse, error)

	QueryOperationsByStatus(ctx sdk.Context, status Operation_Status, pageReq *query.PageRequest) ([]*Operation, *query.PageResponse, error)

	QueryOperationByHash(ctx sdk.Context, chain string, hash []byte) (*Operation, error)
}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context) (genesis *GenesisState, err error)

	ImportGenesis(ctx sdk.Context, genesis *GenesisState) error
}
