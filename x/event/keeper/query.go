package keeper

import (
	"github.com/many-things/mitosis/x/event/types"
)

var _ types.QueryServer = Keeper{}
