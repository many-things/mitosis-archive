package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/event/keeper/state"
	"github.com/many-things/mitosis/x/event/types"
)

// determinism
var _ types.PollKeeper = keeper{}

func (k keeper) FilterNewPolls(ctx sdk.Context, chain string, polls []*types.Poll) ([]*types.Poll, []mitotypes.KV[uint64, []byte], error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return nil, nil, err
	}

	var (
		repo   = state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))
		news   []*types.Poll
		exists []mitotypes.KV[uint64, []byte]
	)
	for _, poll := range polls {
		checksum, err := poll.GetPayload().Hash()
		if err != nil {
			return nil, nil, err
		}

		loaded, err := repo.LoadByHash(checksum)
		if err != nil {
			return nil, nil, err
		}
		if loaded == nil {
			news = append(news, loaded)
		} else {
			exists = append(exists, mitotypes.NewKV(loaded.GetId(), checksum))
		}
	}

	return news, exists, nil
}

func (k keeper) SubmitPolls(
	ctx sdk.Context,
	chain string,
	polls []*types.Poll,
	totalPower sdk.Int,
	valPower sdk.Int,
) ([]mitotypes.KV[uint64, []byte], error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return nil, err
	}

	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))

	set := make([]mitotypes.KV[uint64, []byte], len(polls))
	for i, poll := range polls {
		// initialize poll
		poll.Status = types.Poll_StatusPending
		poll.Tally = &types.Tally{
			TotalPower: &totalPower,
			Confirmed:  &valPower,
		}

		pollId, err := pollRepo.Create(*poll)
		if err != nil {
			return nil, err
		}

		checksum, err := poll.GetPayload().Hash()
		if err != nil {
			return nil, err
		}

		set[i] = mitotypes.NewKV(pollId, checksum)
	}

	return set, nil
}

func (k keeper) VotePolls(ctx sdk.Context, chain string, votes []uint64, valPower sdk.Int) error {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return err
	}

	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))
	for _, id := range votes {
		loaded, err := pollRepo.Load(id)
		if err != nil {
			return err
		}
		if loaded == nil {
			return errors.ErrKeyNotFound
		}

		loaded.Tally.Confirmed = mitotypes.Ref(loaded.Tally.Confirmed.Add(valPower))

		if err := pollRepo.Save(*loaded); err != nil {
			return err
		}
	}

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
