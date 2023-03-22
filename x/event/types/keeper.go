package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params

	SetParams(ctx sdk.Context, params Params)
}

type ChainKeeper interface {
	RegisterChain(ctx sdk.Context, chain string) (byte, error)

	UnregisterChain(ctx sdk.Context, chain string) error

	QueryChain(ctx sdk.Context, chain string) (byte, error)

	QueryChains(ctx sdk.Context, pageReq *query.PageRequest) ([]mitotypes.KV[string, byte], *query.PageResponse, error)
}

type GenesisKeeper interface {
	ExportGenesis(ctx sdk.Context) (genesis *GenesisState, err error)
	ImportGenesis(ctx sdk.Context, genesis *GenesisState) error
}

type PollKeeper interface {
	// FilterNewPolls specifies fresh poll and returns new poll and existing polls' index
	FilterNewPolls(ctx sdk.Context, chain string, polls []*Poll) ([]*Poll, []mitotypes.KV[uint64, []byte], error)

	// SubmitPolls handles [server.MsgSubmit] - returns key value set of poll_id and event_hash
	SubmitPolls(ctx sdk.Context, chain string, val sdk.ValAddress, polls []*Poll) ([]mitotypes.KV[uint64, []byte], error)

	// VotePolls handles [server.MsgVote]
	VotePolls(ctx sdk.Context, chain string, val sdk.ValAddress, votes []uint64) error

	// QueryPoll handles [server.QueryPoll]
	QueryPoll(ctx sdk.Context, chain string, id uint64) (*Poll, error)

	// QueryPolls handles [server.QueryPolls]
	QueryPolls(ctx sdk.Context, chain string, pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *Poll], *query.PageResponse, error)
}

type ProxyKeeper interface {
	RegisterProxy(ctx sdk.Context, val sdk.ValAddress, prx sdk.AccAddress) error

	ClearProxy(ctx sdk.Context, val sdk.ValAddress) error

	QueryProxy(ctx sdk.Context, val sdk.ValAddress) (sdk.AccAddress, bool)

	QueryProxyReverse(ctx sdk.Context, prx sdk.AccAddress) (sdk.ValAddress, bool)

	QueryProxies(ctx sdk.Context, pageReq *query.PageRequest) ([]mitotypes.KV[sdk.ValAddress, sdk.AccAddress], *query.PageResponse, error)
}

type SnapshotKeeper interface {
	CreateSnapshot(ctx sdk.Context, total sdk.Int, powers []mitotypes.KV[sdk.ValAddress, int64]) (*EpochInfo, error)

	VotingPowerOf(ctx sdk.Context, epoch *uint64, val sdk.ValAddress) (int64, error)

	LatestSnapshotEpoch(ctx sdk.Context) (*EpochInfo, error)
}
