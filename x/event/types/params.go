package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmoserrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultEpochInterval = 10
	DefaultPollThreshold = "0.5"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		EpochInterval: DefaultEpochInterval, // 10 blocks
		PollThreshold: mitotypes.Ref(sdk.MustNewDecFromStr(DefaultPollThreshold)),
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p *Params) Validate() error {
	if p.PollThreshold.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(cosmoserrors.ErrInvalidType, "poll threshold must be positive")
	} else if p.PollThreshold.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(cosmoserrors.ErrInvalidType, "poll threshold must be less than 1")
	}

	return nil
}

// String implements the Stringer interface.
func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
