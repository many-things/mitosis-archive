package msgconv

import (
	"gotest.tools/assert"
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
}
