package msgconv

import "github.com/pkg/errors"

var (
	AssetMapping = map[string]map[string]string{
		"usdc": {
			"osmo-test-5": "factory/osmo109ns4u04l44kqdkvp876hukd3hxz8zzm7809el/uusdc",
			"evm-5":       "0x5Cbb2F9f7E54c5B4656C3B563ff5650a0866A3EF",
		},
	}
	AssetMappingReverse = make(map[string]map[string]string)
)

func init() {
	for asset, perChain := range AssetMapping {
		for chain, addr := range perChain {
			if AssetMappingReverse[addr] == nil {
				AssetMappingReverse[addr] = make(map[string]string)
			}
			AssetMappingReverse[addr][chain] = asset
		}
	}
}

func convertDenomIO(src, dest, in string) (string, error) {
	perChain, ok := AssetMappingReverse[in]
	if !ok {
		return "", errors.Errorf("unknown asset %s", in)
	}

	asset, ok := perChain[src]
	if !ok {
		return "", errors.Errorf("unknown asset %s for chain %s", in, src)
	}

	return AssetMapping[asset][dest], nil
}
