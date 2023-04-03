package txconv

import (
	txconvtypes "github.com/many-things/mitosis/pkg/txconv/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
)

func makeEvmInfo(chainID, chainName string) mitotypes.KV[string, txconvtypes.ChainInfo] {
	chainInfo := txconvtypes.ChainInfo{
		Type:      txconvtypes.ChainTypeEvm,
		ChainID:   chainID,
		ChainName: chainName,

		TxConv: func(s txconvtypes.Signer, id uint64, args ...[]byte) ([]byte, []byte, error) {
			signer := s.(txconvtypes.EvmSigner)
			_ = signer

			// TODO: implement me

			return nil, nil, nil
		},
	}

	return mitotypes.NewKV(chainID, chainInfo)
}
