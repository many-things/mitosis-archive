package keeper

import "github.com/many-things/mitosis/x/context/server"

var _ server.QueryServer = Keeper{}
