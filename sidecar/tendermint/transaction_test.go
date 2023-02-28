package tendermint

import (
	"gotest.tools/assert"
	"testing"
)

func Test_BroadCastRawTx(t *testing.T) {

}

func Test_IsMnemonic(t *testing.T) {
	validMnemonic := "burst visa embark foam office album waste autumn remove tourist moment tail camp trumpet blue grunt catalog metal metal simple school item cotton apart"

	result := IsMnemonic(validMnemonic)
	assert.Equal(t, result, true)

	invalidMnemonic := "burst visa embark foam office album waste autumn remove tourist moment tail camp trumpet blue grunt catalog metal metal simple school item cotton"

	result = IsMnemonic(invalidMnemonic)
	assert.Equal(t, result, false)
}
