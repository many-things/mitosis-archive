package tendermint

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotype "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostype "github.com/cosmos/cosmos-sdk/types"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	cosmostx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/go-bip39"
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

type RawTx struct {
	Mode    string
	TxBytes []byte
}

func NewWalletWithMnemonic(mnemonic string, chainPrefix string, chainID string, dialUrl string) (*Wallet, error) {
	privBytes, err := bip39.MnemonicToByteArray(mnemonic)
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
	fromAddress, err := libs.ConvertPubKeyToBech32Address(w.privateKey.PubKey(), w.ChainPrefix)
	if err != nil {
		return nil, err
	}

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

func (w *Wallet) BroadCastRawTx(rawTxByte []byte) error {
	rawTxBody := RawTx{
		Mode:    "BROADCAST_MODE_SYNC",
		TxBytes: rawTxByte,
	}

	// TODO: change LCD to gRPC
	postBodyBytes, _ := json.Marshal(rawTxBody)
	resp, err := http.Post(w.DialURL+"/cosmos/tx/v1beta1/txs", "application/json", bytes.NewBuffer(postBodyBytes))
	if err != nil || resp.StatusCode != 200 {
		return err // TODO:
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// TODO: Return Response
	return nil
}

func (w *Wallet) BroadcastMsg(msg cosmostype.Msg) error {
	rawTx, err := w.CreateSignedRawTx(msg)
	if err != nil {
		return nil
	}

	err = w.BroadCastRawTx(rawTx)
	return err
}

func IsMnemonic(mnemonic_or_privkey string) bool {
	return bip39.IsMnemonicValid(mnemonic_or_privkey)
}
