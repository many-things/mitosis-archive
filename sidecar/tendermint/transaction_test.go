package tendermint

import (
	"bytes"
	sdkmath "cosmossdk.io/math"
	"encoding/json"
	cosmostype "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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

func Test_CreateSignedRawTx(t *testing.T) {
	accountInfo := AccountInfo{
		SequenceNumber: 1,
		AccountNumber:  2,
	}

	msg := banktypes.MsgSend{
		FromAddress: "",
		ToAddress:   "",
		Amount: []cosmostype.Coin{
			{Denom: "uosmo", Amount: sdkmath.NewInt(100000)},
		},
	}

	// I want to know good replacement of this..
	expectedResult := []byte{10, 51, 10, 49, 10, 28, 47, 99, 111, 115, 109, 111, 115, 46, 98, 97, 110, 107, 46, 118,
		49, 98, 101, 116, 97, 49, 46, 77, 115, 103, 83, 101, 110, 100, 18, 17, 26, 15, 10, 5, 117, 111, 115, 109, 111,
		18, 6, 49, 48, 48, 48, 48, 48, 18, 88, 10, 80, 10, 70, 10, 31, 47, 99, 111, 115, 109, 111, 115, 46, 99, 114,
		121, 112, 116, 111, 46, 115, 101, 99, 112, 50, 53, 54, 107, 49, 46, 80, 117, 98, 75, 101, 121, 18, 35, 10, 33,
		3, 145, 102, 194, 137, 185, 249, 5, 229, 95, 158, 61, 249, 246, 157, 127, 53, 107, 74, 34, 9, 95, 137, 79, 71,
		21, 113, 74, 164, 181, 102, 6, 175, 18, 4, 10, 2, 8, 1, 24, 1, 18, 4, 16, 160, 141, 6, 26, 64, 197, 130, 132,
		219, 24, 202, 125, 90, 62, 83, 137, 213, 54, 220, 87, 194, 123, 205, 135, 77, 249, 36, 165, 21, 51, 0, 88, 199,
		42, 22, 238, 80, 114, 48, 77, 138, 120, 81, 247, 117, 152, 6, 102, 202, 241, 22, 160, 245, 253, 243, 111, 208,
		146, 193, 163, 210, 232, 247, 109, 40, 83, 211, 60, 5}

	privKey := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	wallet, err := NewWallet(privKey, "mito", "", "https://test.com")
	assert.NilError(t, err)

	result, err := wallet.CreateSignedRawTx(&msg, accountInfo)
	assert.NilError(t, err)

	assert.DeepEqual(t, result, expectedResult)
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
