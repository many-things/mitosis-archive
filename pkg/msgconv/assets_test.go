package msgconv

import (
	"gotest.tools/assert"
	"log"
	"testing"
)

func TestAssetMappings(t *testing.T) {
	denom := "usdc"
	chain := "osmo-test-5"
	input := AssetMapping[denom][chain]
	assert.Equal(
		t,
		denom,
		AssetMappingReverse[input][chain],
	)

	log.Println(AssetMappingReverse)

	log.Println(AssetMappingReverse["0x5Cbb2F9f7E54c5B4656C3B563ff5650a0866A3EF"]["osmo-test-5"])
}
