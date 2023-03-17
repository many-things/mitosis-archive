package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

type PollKeeper interface {
	// SubmitPoll handles [server.MsgSubmit]
	SubmitPoll(ctx sdk.Context, val sdk.ValAddress, chain string, events []*types.Event) ([]mitotypes.KV[uint64, []byte], error)

	// VotePoll handles [server.MsgVote]
	VotePoll(ctx sdk.Context, val sdk.ValAddress, chain string, ids []uint64) error

	// QueryPoll handles [server.QueryPoll]
	QueryPoll(ctx sdk.Context, chain string, id uint64) (*types.Poll, error)

	// QueryPolls handles [server.QueryPolls]
	QueryPolls(ctx sdk.Context, chain string, pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error)
}

func (k keeper) SubmitPoll(
	ctx sdk.Context,
	val sdk.ValAddress,
	chain string,
	events []*types.Event,
) ([]mitotypes.KV[uint64, []byte], error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return nil, err
	}

	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))

	set := make([]mitotypes.KV[uint64, []byte], len(events))
	for i, event := range events {
		checksum, err := event.Hash()
		if err != nil {
			return nil, err
		}

		var pollId uint64
		// check duplication by query poll by event checksum
		loaded, err := pollRepo.LoadByHash(checksum)
		if err != nil {
			return nil, err
		}
		if loaded == nil {
			// create new one
			poll := types.Poll{
				// don't need to fill now
				//Id:       0,
				//OpId:     0,

				Chain:    chain,
				Proposer: val,
				Status:   types.Poll_StatusPending,
				Tally: &types.Tally{

					TotalPower: mitotypes.Ref(sdk.NewUint(0)),
					Ack:        mitotypes.Ref(sdk.NewUint(0)),
					Nak:        mitotypes.Ref(sdk.NewUint(0)),
				},
				Payload: event,
			}
			pollId, err = pollRepo.Save(poll)
			if err != nil {
				return nil, err
			}
		} else {
			pollId = loaded.Id
		}

		set[i] = mitotypes.NewKV(pollId, checksum)
	}

	return set, nil
}

func (k keeper) VotePoll(ctx sdk.Context, val sdk.ValAddress, chain string, ids []uint64) error {
	// TODO
	return nil
}

func (k keeper) QueryPoll(ctx sdk.Context, chain string, id uint64) (*types.Poll, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return nil, err
	}
	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))

	poll, err := pollRepo.Load(id)
	if err != nil {
		return nil, err
	}

	return poll, nil
}

func (k keeper) QueryPolls(ctx sdk.Context, chain string, pageReq *query.PageRequest) ([]mitotypes.KV[uint64, *types.Poll], *query.PageResponse, error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return nil, nil, err
	}
	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))

	polls, pageResp, err := pollRepo.Paginate(pageReq)
	if err != nil {
		return nil, nil, err
	}

	return polls, pageResp, nil
}
