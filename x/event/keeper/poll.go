package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

var _ PollKeeper = Keeper{}

type PollKeeper interface {
	// SubmitPoll handles [server.MsgSubmit]
	SubmitPoll(ctx sdk.Context, sender sdk.AccAddress, chain string, events []*types.Event) (uint64, error)

	// VotePoll handles [server.MsgVote]
	VotePoll(ctx sdk.Context, sedner sdk.AccAddress, chain string, ids []uint64) error

	// QueryPoll handles [server.QueryPoll]
	QueryPoll(ctx sdk.Context, chain string, id uint64) (*types.Poll, error)

	// QueryPolls handles [server.QueryPolls]
	QueryPolls(ctx sdk.Context, chain string, pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error)
}

func (k Keeper) SubmitPoll(ctx sdk.Context, sender sdk.AccAddress, chain string, events []*types.Event) (uint64, error) {
	// TODO
	return 0, nil
}

func (k Keeper) VotePoll(ctx sdk.Context, sedner sdk.AccAddress, chain string, ids []uint64) error {
	// TODO
	return nil
}

func (k Keeper) QueryPoll(ctx sdk.Context, chain string, id uint64) (*types.Poll, error) {
	defaultChain := byte(0x01) // TODO: make chain registry
	pollRepo := state.NewKVPollRepo(k.cdc, defaultChain, ctx.KVStore(k.storeKey))

	poll, err := pollRepo.Load(id)
	if err != nil {
		return nil, err
	}

	return &poll, nil
}

func (k Keeper) QueryPolls(ctx sdk.Context, chain string, pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error) {
	defaultChain := byte(0x01) // TODO: make chain registry
	pollRepo := state.NewKVPollRepo(k.cdc, defaultChain, ctx.KVStore(k.storeKey))

	polls, pageResp, err := pollRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return polls, pageResp, nil
}
