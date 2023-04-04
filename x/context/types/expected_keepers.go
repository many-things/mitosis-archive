package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
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
	RegisterSignEvent(ctx sdk.Context, chainID string, sign *multisigtypes.Sign) (uint64, error)
}
