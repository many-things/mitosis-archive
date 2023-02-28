package tendermint

import (
	"bytes"
	"encoding/json"
	"github.com/many-things/mitosis/sidecar/tendermint/libs"
	"github.com/many-things/mitosis/sidecar/tendermint/libs/mocks"
	"gotest.tools/assert"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_BroadCastRawTx(t *testing.T) {
	mocked_raw_tx := []byte("Hello World!")

	wallet, err := NewWallet("", "", "", "https://test.com")
	assert.NilError(t, err)

	expectedBody := RawTx{
		Mode:    "BROADCAST_MODE_SYNC",
		TxBytes: mocked_raw_tx,
	}
	jsonBytes, err := json.Marshal(expectedBody)
	assert.NilError(t, err)

	// Mock HTTP Request
	libs.Client = &mocks.MockClient{}
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		reqBody, err := io.ReadAll(req.Body)
		assert.NilError(t, err)

		assert.Equal(t, req.URL.String(), "https://test.com/cosmos/tx/v1beta1/txs")
		assert.DeepEqual(t, reqBody, jsonBytes)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
		}, nil
	}

	wallet.BroadCastRawTx(mocked_raw_tx)
}

func Test_IsMnemonic(t *testing.T) {
	validMnemonic := "burst visa embark foam office album waste autumn remove tourist moment tail camp trumpet blue grunt catalog metal metal simple school item cotton apart"

	result := IsMnemonic(validMnemonic)
	assert.Equal(t, result, true)

	invalidMnemonic := "burst visa embark foam office album waste autumn remove tourist moment tail camp trumpet blue grunt catalog metal metal simple school item cotton"

	result = IsMnemonic(invalidMnemonic)
	assert.Equal(t, result, false)
}
