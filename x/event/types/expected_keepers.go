package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type BaseKeeper interface {
	//GetIncomingEvent(ctx sdk.Context, txHash string, evtIndex uint64) (*IncomingEvent, error)
	//ListIncomingEvent(ctx sdk.Context) ([]*IncomingEvent, error)
	//
	//GetOutgoingEvent(ctx sdk.Context, txHash string) (*OutgoingEvent, error)
	//ListOutgoingEvent(ctx sdk.Context) ([]*OutgoingEvent, error)
}

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
