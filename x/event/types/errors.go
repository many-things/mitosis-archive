package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
