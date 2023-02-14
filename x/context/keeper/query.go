package keeper

import (
	"github.com/many-things/mitosis/x/context/types"
)

var _ types.QueryServer = Keeper{}
