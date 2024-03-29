package keeper

import (
	"fmt"

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
		return nil, nil, fmt.Errorf("chain cannot found: %w", err)
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
			return nil, nil, fmt.Errorf("load by hash: %w", err)
		}
		if loaded == nil {
			news = append(news, poll)
		} else {
			exists = append(exists, mitotypes.NewKV(loaded.GetId(), checksum))
		}
	}

	return news, exists, nil
}

func (k keeper) SubmitPolls(
	ctx sdk.Context,
	chain string,
	val sdk.ValAddress,
	polls []*types.Poll,
) ([]mitotypes.KV[uint64, []byte], error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return nil, fmt.Errorf("chain: %w", err)
	}

	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))
	epoch, err := snapshotRepo.LatestEpoch()
	if err != nil {
		return nil, fmt.Errorf("snapshot: %w", err)
	}

	power, err := snapshotRepo.PowerOf(epoch.GetEpoch(), val)
	if err != nil {
		return nil, fmt.Errorf("powerof: %w", err)
	}

	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))

	set := make([]mitotypes.KV[uint64, []byte], len(polls))
	for i, poll := range polls {
		// initialize poll
		poll.Epoch = epoch.GetEpoch()
		poll.Status = types.Poll_StatusPending
		poll.Tally = &types.Tally{
			TotalPower: epoch.TotalPower,
			Confirmed:  mitotypes.Ref(sdk.NewInt(power)),
		}

		pollID, err := pollRepo.Create(*poll)
		if err != nil {
			return nil, fmt.Errorf("create poll: %w", err)
		}
		pollRepo.SetVoted(pollID, val)

		checksum, err := poll.GetPayload().Hash()
		if err != nil {
			return nil, fmt.Errorf("get payload hash: %w", err)
		}

		set[i] = mitotypes.NewKV(pollID, checksum)
	}

	return set, nil
}

func (k keeper) VotePolls(ctx sdk.Context, chain string, val sdk.ValAddress, votes []uint64) error {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return fmt.Errorf("chain: %w", err)
	}

	snapshotRepo := state.NewKVSnapshotRepo(k.cdc, ctx.KVStore(k.storeKey))

	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))
	for _, id := range votes {
		if pollRepo.IsVoted(id, val) {
			continue
		}

		loaded, err := pollRepo.Load(id)
		if err != nil {
			return fmt.Errorf("load poll: %w", err)
		}
		if loaded == nil {
			return fmt.Errorf("load poll: %w", errors.ErrKeyNotFound)
		}

		power, err := snapshotRepo.PowerOf(loaded.GetEpoch(), val)
		if err != nil {
			return fmt.Errorf("get power of: %w", err)
		}

		loaded.Tally.Confirmed = mitotypes.Ref(loaded.Tally.Confirmed.Add(sdk.NewInt(power)))

		if err := pollRepo.Save(*loaded); err != nil {
			return fmt.Errorf("saving poll: %w", err)
		}
		pollRepo.SetVoted(id, val)
	}

	return nil
}

func (k keeper) FlushPolls(ctx sdk.Context, chain string, threshold sdk.Dec) ([]mitotypes.KV[uint64, *types.Poll], error) {
	chainRepo := state.NewKVChainRepo(k.cdc, ctx.KVStore(k.storeKey))
	chainPrefix, err := chainRepo.Load(chain)
	if err != nil {
		return nil, err
	}

	pollRepo := state.NewKVPollRepo(k.cdc, chainPrefix, ctx.KVStore(k.storeKey))
	flushed, err := pollRepo.Flush(threshold)
	if err != nil {
		return nil, err
	}

	return flushed, nil
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
