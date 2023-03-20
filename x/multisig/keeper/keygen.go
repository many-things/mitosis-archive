package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/types"
)

type KeygenKeeper interface {
	RegisterKeygenEvent(ctx sdk.Context, chainId string, keygen *types.Keygen) error

	QueryKeygen(ctx sdk.Context, chainId string, id uint64) (*types.Keygen, error)
	QueryKeygenList(ctx sdk.Context, chainId string, pageReq *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error)
}

func (k keeper) RegisterKeygenEvent(ctx sdk.Context, chainId string, keygen *types.Keygen) error {
	// TODO: implement me!
	return nil
}

func (k keeper) QueryKeygen(ctx sdk.Context, chainId string, id uint64) (*types.Keygen, error) {
	// TODO: implement me
	panic("Implement me!")
}

func (k keeper) QueryKeygenList(ctx sdk.Context, chainId string, pageReq *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error) {
	// TODO: implement me
	panic("Implement me!")
}
