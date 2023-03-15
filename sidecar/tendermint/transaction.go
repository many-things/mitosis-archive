package tendermint

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
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
)

type Wallet interface {
	GetAddress() (string, error)
	GetAccountInfo() (*AccountInfo, error)
	CreateSignedRawTx(msg cosmostype.Msg, accountInfo AccountInfo) ([]byte, error)
	BroadcastRawTx(rawTxByte []byte) error
	BroadcastMsg(msg cosmostype.Msg) error
}

type wallet struct {
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

type MnemonicDeriveOption struct {
	BIP39Passphrase string
	HDPath          string
}

type MnemonicDeriveOptionHandler = func(option *MnemonicDeriveOption)

func WithBIP39Passphrase(passphrase string) MnemonicDeriveOptionHandler {
	return func(option *MnemonicDeriveOption) {
		option.BIP39Passphrase = passphrase
	}
}

func WithHDPath(hdPath string) MnemonicDeriveOptionHandler {
	return func(option *MnemonicDeriveOption) {
		option.HDPath = hdPath
	}
}

func NewWalletWithMnemonic(mnemonic string, chainPrefix string, chainID string, dialUrl string, options ...MnemonicDeriveOptionHandler) (Wallet, error) {
	deriveFn := hd.Secp256k1.Derive()
	option := &MnemonicDeriveOption{
		BIP39Passphrase: "",
		HDPath:          hd.CreateHDPath(cosmostype.CoinType, 0, 0).String(),
	}

	for _, o := range options {
		o(option)
	}
	privBytes, err := deriveFn(mnemonic, option.BIP39Passphrase, option.HDPath)

	if err != nil {
		return nil, err
	}

	return &wallet{
		privateKey:  &secp256k1.PrivKey{Key: privBytes},
		ChainPrefix: chainPrefix,
		ChainID:     chainID,
		DialURL:     dialUrl,
	}, nil
}

func NewWallet(privateKey string, chainPrefix string, chainID string, dialUrl string) (Wallet, error) {
	privBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	return &wallet{
		privateKey:  &secp256k1.PrivKey{Key: privBytes},
		ChainPrefix: chainPrefix,
		ChainID:     chainID,
		DialURL:     dialUrl,
	}, nil
}

func (w wallet) createTxConfig() client.TxConfig {
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	return cosmostx.NewTxConfig(codec, cosmostx.DefaultSignModes)
}

func (w wallet) GetAddress() (string, error) {
	return libs.ConvertPubKeyToBech32Address(w.privateKey.PubKey(), w.ChainPrefix)
}

func (w wallet) GetAccountInfo() (*AccountInfo, error) {
	fromAddress, err := w.GetAddress()
	if err != nil {
		return nil, err
	}

	response, err := libs.JsonGet(w.DialURL + "/cosmos/auth/v1beta1/accounts/" + fromAddress)
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

func (w wallet) CreateSignedRawTx(msg cosmostype.Msg, accountInfo AccountInfo) ([]byte, error) {
	txConfig := w.createTxConfig()
	txBuilder := txConfig.NewTxBuilder()

	txBuilder.SetMsgs(msg)
	txBuilder.SetGasLimit(100000)

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

func (w wallet) BroadcastRawTx(rawTxByte []byte) error {
	rawTxBody := RawTx{
		Mode:    "BROADCAST_MODE_SYNC",
		TxBytes: rawTxByte,
	}
	resp, err := libs.JsonPost(w.DialURL+"/cosmos/tx/v1beta1/txs", rawTxBody)

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

func (w wallet) BroadcastMsg(msg cosmostype.Msg) error {
	accountInfo, err := w.GetAccountInfo()
	if err != nil {
		return err
	}

	rawTx, err := w.CreateSignedRawTx(msg, *accountInfo)
	if err != nil {
		return nil
	}

	err = w.BroadcastRawTx(rawTx)
	return err
}

func IsMnemonic(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}
