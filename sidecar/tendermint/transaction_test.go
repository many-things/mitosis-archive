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
	mockedRawTx := []byte("Hello World!")

	wallet, err := NewWallet("", "", "", "https://test.com")
	assert.NilError(t, err)

	expectedBody := RawTx{
		Mode:    "BROADCAST_MODE_SYNC",
		TxBytes: mockedRawTx,
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

	// Call
	wallet.BroadcastRawTx(mockedRawTx)
}

func Test_GetAccountInfo(t *testing.T) {
	type MockedAccountResponse struct {
		Sequence      uint64 `json:"sequence"`
		AccountNumber uint64 `json:"account_number"`
	}

	type MockedResponse struct {
		Account MockedAccountResponse `json:"account"`
	}

	response := MockedResponse{
		Account: MockedAccountResponse{
			Sequence:      1,
			AccountNumber: 2,
		},
	}

	privKey := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	wallet, err := NewWallet(privKey, "mito", "", "https://test.com")
	assert.NilError(t, err)

	libs.Client = &mocks.MockClient{}
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, req.URL.String(), "https://test.com/cosmos/auth/v1beta1/accounts/mito17h6ufy9kmpkc8ldzzsltl26y8agm604ae6ea2r")

		jsonResponse, err := json.Marshal(response)
		assert.NilError(t, err)

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(jsonResponse)),
		}, nil
	}

	accountInfo, err := wallet.GetAccountInfo()
	assert.NilError(t, err)

	assert.Equal(t, accountInfo.SequenceNumber, uint64(1))
	assert.Equal(t, accountInfo.AccountNumber, uint64(2))
}

func Test_IsMnemonic(t *testing.T) {
	validMnemonic := "burst visa embark foam office album waste autumn remove tourist moment tail camp trumpet blue grunt catalog metal metal simple school item cotton apart"

	result := IsMnemonic(validMnemonic)
	assert.Equal(t, result, true)

	invalidMnemonic := "burst visa embark foam office album waste autumn remove tourist moment tail camp trumpet blue grunt catalog metal metal simple school item cotton"

	result = IsMnemonic(invalidMnemonic)
	assert.Equal(t, result, false)
}
