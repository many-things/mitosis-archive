package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

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
	SubmitPolls(ctx sdk.Context, chain string, polls []*Poll, totalPower, valPower sdk.Int) ([]mitotypes.KV[uint64, []byte], error)

	// VotePolls handles [server.MsgVote]
	VotePolls(ctx sdk.Context, chain string, votes []uint64, valPower sdk.Int) error

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

type BaseKeeper interface {
	GetParams(ctx sdk.Context) Params
	//SetParams(ctx sdk.Context, params Params)
	//Logger(ctx sdk.Context) log.Logger
	//
	//ChainKeeper
	//GenesisKeeper
	//PollKeeper
	//ProxyKeeper
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// StakingKeeper defined the expected staking keeper used for retrieve validator info
type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)

	GetLastTotalPower(ctx sdk.Context) sdk.Int

	GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64)

	IterateLastValidatorPowers(ctx sdk.Context, handler func(operator sdk.ValAddress, power int64) (stop bool))
}
