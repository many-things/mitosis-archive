package keeper

import (
	"github.com/many-things/mitosis/x/multisig/types"
)

var _ types.QueryServer = Keeper{}
