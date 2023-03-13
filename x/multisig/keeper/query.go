package keeper

import "github.com/many-things/mitosis/x/multisig/server"

var _ server.QueryServer = Keeper{}
