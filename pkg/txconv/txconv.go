package txconv

import (
	sdkerrutils "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/pkg/txconv/cosmos"
	txconvtypes "github.com/many-things/mitosis/pkg/txconv/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

var Chains = []mitotypes.KV[string, txconvtypes.ChainInfo]{
	cosmos.MakeChainInfo("osmosis-1", "osmosis-mainnet"),
	cosmos.MakeChainInfo("osmo-test-4", "osmosis-testnet"),

	makeEvmInfo("evm-1", "eth-mainnet"),
	makeEvmInfo("evm-5", "eth-testnet-goerli"),
}

// Convert returns the full bytes of tx, and it's hash
func Convert[T txconvtypes.Signer](signer T, chainID string, id uint64, args ...[]byte) ([]byte, []byte, error) {
	chain := mitotypes.FindKV(
		Chains,
		func(k string, _ txconvtypes.ChainInfo, _ int) bool { return k == chainID },
	)
	if chain == nil {
		return nil, nil, sdkerrutils.Wrap(sdkerrors.ErrNotFound, "supported chain")
	}

	if signer.Type() != chain.Value.Type {
		return nil, nil, sdkerrutils.Wrap(sdkerrors.ErrPanic, "signer type does not match chain type")
	}

	return chain.Value.TxConv(signer, id, args...)
}
