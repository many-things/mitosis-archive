package tendermint

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotype "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostype "github.com/cosmos/cosmos-sdk/types"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	cosmostx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/many-things/mitosis/sidecar/tendermint/libs"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

type Wallet struct {
	privateKey  cryptotype.PrivKey
	ChainPrefix string
	ChainID     string
	DialURL     string
}

type AccountInfo struct {
	SequenceNumber uint64
	AccountNumber  uint64
}

func NewWallet(privateKey string, chainPrefix string, chainID string, dialUrl string) (*Wallet, error) {
	privBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		privateKey:  &secp256k1.PrivKey{Key: privBytes},
		ChainPrefix: chainPrefix,
		ChainID:     chainID,
		DialURL:     dialUrl,
	}, nil
}

func (w *Wallet) createTxConfig() client.TxConfig {
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	return cosmostx.NewTxConfig(codec, cosmostx.DefaultSignModes)
}

func (w *Wallet) GetAccountInfo() (*AccountInfo, error) {
	fromAddress := libs.ConvertUncompressedSecp256k1ToBech32(w.privateKey.PubKey().Bytes(), w.ChainPrefix)

	response, err := http.Get(w.DialURL + "/cosmos/auth/v1beta1/accounts/" + fromAddress)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{
		SequenceNumber: gjson.GetBytes(data, "account.sequence").Uint(),
		AccountNumber:  gjson.GetBytes(data, "account.account_number").Uint(),
	}, nil
}

func (w *Wallet) CreateSignedRawTx(msg cosmostype.Msg) ([]byte, error) {
	txConfig := w.createTxConfig()
	txBuilder := txConfig.NewTxBuilder()

	txBuilder.SetMsgs(msg)
	txBuilder.SetGasLimit(100000)

	accountInfo, err := w.GetAccountInfo()
	if err != nil {
		return nil, err
	}

	signerData := authsigning.SignerData{
		ChainID:       w.ChainID,
		AccountNumber: accountInfo.AccountNumber,
		Sequence:      accountInfo.SequenceNumber,
	}
	signatureData := txsigning.SingleSignatureData{
		SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	signature := txsigning.SignatureV2{
		PubKey:   w.privateKey.PubKey(),
		Data:     &signatureData,
		Sequence: accountInfo.SequenceNumber,
	}

	if err := txBuilder.SetSignatures(signature); err != nil {
		return nil, err
	}

	// Generate bytes to be signed
	bytesToSign, err := txConfig.SignModeHandler().GetSignBytes(txsigning.SignMode_SIGN_MODE_DIRECT, signerData, txBuilder.GetTx())
	if err != nil {
		return nil, err
	}
	signedBytes, err := w.privateKey.Sign(bytesToSign)
	if err != nil {
		return nil, err
	}

	signatureData = txsigning.SingleSignatureData{
		SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
		Signature: signedBytes,
	}
	signature = txsigning.SignatureV2{
		PubKey:   w.privateKey.PubKey(),
		Data:     &signatureData,
		Sequence: accountInfo.SequenceNumber,
	}

	txBuilder.SetSignatures(signature)

	// Is `txConfig.TxJSONDecoder()` required?
	return txConfig.TxEncoder()(txBuilder.GetTx())
}
