package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"
	multisigtypes "github.com/many-things/mitosis/x/multisig/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// MultisigKeeper defined the expected interface needed to initiate signing sequence
type MultisigKeeper interface {
	RegisterSignEvent(ctx sdk.Context, chainID string, sign *exported.Sign) (uint64, error)

	QueryKeygenResult(ctx sdk.Context, chainID string, keyID uint64) (*multisigtypes.KeygenResult, error)
	QueryKeygen(ctx sdk.Context, chainID string, id uint64) (*multisigtypes.Keygen, error)
	QueryKeygenList(ctx sdk.Context, chainID string, page *query.PageRequest) ([]mitosistype.KV[uint64, *multisigtypes.Keygen], *query.PageResponse, error)
}

type EventKeeper interface {
	QueryChains(ctx sdk.Context, pageReq *query.PageRequest) ([]mitosistype.KV[string, byte], *query.PageResponse, error)
}
