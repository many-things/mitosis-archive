package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

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

	GetLastTotalPower(ctx sdk.Context) sdkmath.Int

	GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64)

	IterateLastValidatorPowers(ctx sdk.Context, handler func(operator sdk.ValAddress, power int64) (stop bool))
}

// ContextKeeper defined the expected context keeper used for push confirmed event to operation queue
type ContextKeeper interface {
	InitOperation(ctx sdk.Context, chain string, poll *Poll) (uint64, error)

	FinishOperation(ctx sdk.Context, id uint64, poll *Poll) error
}
