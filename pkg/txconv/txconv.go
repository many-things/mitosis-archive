package txconv

import (
	sdkerrutils "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/pkg/txconv/cosmos"
	txconvtypes "github.com/many-things/mitosis/pkg/txconv/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

var Converter ConverterI = &converter{}

type ConverterI interface {
	RegisterEvmChain(chainID, chainName string) error
	RegisterCosmosChain(chainID, chainName string, encoder client.TxConfig) error

	Convert(signer txconvtypes.Signer, chainID string, id uint64, args ...[]byte) ([]byte, []byte, error)
}

type converter struct {
	chainReg []mitotypes.KV[string, txconvtypes.ChainInfo]
}

func (c *converter) findChain(chainID string) *mitotypes.KV[string, txconvtypes.ChainInfo] {
	return mitotypes.FindKV(
		c.chainReg,
		func(k string, _ txconvtypes.ChainInfo, _ int) bool { return k == chainID },
	)
}

func (c *converter) RegisterEvmChain(chainID, chainName string) error {
	if chain := c.findChain(chainID); chain != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrConflict, "chain already registered")
	}

	c.chainReg = append(c.chainReg, makeEvmInfo(chainID, chainName))
	return nil
}

func (c *converter) RegisterCosmosChain(chainID, chainName string, encoder client.TxConfig) error {
	if chain := c.findChain(chainID); chain != nil {
		return sdkerrutils.Wrap(sdkerrors.ErrConflict, "chain already registered")
	}

	c.chainReg = append(c.chainReg, cosmos.MakeChainInfo(chainID, chainName, encoder))
	return nil
}

// Convert returns the full bytes of tx, and it's hash
func (c *converter) Convert(signer txconvtypes.Signer, chainID string, id uint64, args ...[]byte) ([]byte, []byte, error) {
	chain := c.findChain(chainID)
	if chain == nil {
		return nil, nil, sdkerrutils.Wrap(sdkerrors.ErrNotFound, "supported chain")
	}

	if signer.Type() != chain.Value.Type {
		return nil, nil, sdkerrutils.Wrap(sdkerrors.ErrPanic, "signer type does not match chain type")
	}

	return chain.Value.TxConv(signer, id, args...)
}
