package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultEpochInterval = 10
	DefaultPollThreshold = 0.5
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
		PollThreshold: DefaultPollThreshold,
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
	pollThreshold := p.GetPollThreshold()
	if pollThreshold < 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "poll threshold must be positive")
	} else if pollThreshold > 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "poll threshold must be less than 1")
	}

	return nil
}

// String implements the Stringer interface.
func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
