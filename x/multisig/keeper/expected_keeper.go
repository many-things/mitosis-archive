package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/types"
)

type KeygenKeeper interface {
	RegisterKeygenEvent(ctx sdk.Context, chainId string, keygen *types.Keygen) (uint64, error)

	QueryKeygen(ctx sdk.Context, chainId string, id uint64) (*types.Keygen, error)
	QueryKeygenList(ctx sdk.Context, chainId string, pageReq *query.PageRequest) ([]mitosistype.KV[uint64, *types.Keygen], *query.PageResponse, error)
}

type PubKeyKeeper interface {
	RegisterPubKey(ctx sdk.Context, chainId string, pubKey *types.PubKey) error
	RemovePubKey(ctx sdk.Context, chainId string, keyId uint64, participant sdk.ValAddress) error

	QueryPubKey(ctx sdk.Context, chainId string, keyId uint64, participant sdk.ValAddress) (*types.PubKey, error)
	QueryPubKeyList(ctx sdk.Context, chainId string, keyId uint64, page *query.PageRequest) (*mitosistype.KV[sdk.ValAddress, *types.PubKey], *query.PageResponse, error)
}
