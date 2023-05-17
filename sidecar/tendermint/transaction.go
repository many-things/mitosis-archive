package tendermint

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotype "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	accounttypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/go-bip39"
	"github.com/many-things/mitosis/sidecar/tendermint/libs"
	"google.golang.org/grpc"
)

type Wallet interface {
	GetAddress() (string, error)
	GetAccountInfo() (*AccountInfo, error)
	CreateSignedRawTx(msg sdk.Msg, accountInfo AccountInfo) ([]byte, error)
	BroadcastRawTx(rawTxByte []byte) error
	BroadcastMsg(msg sdk.Msg) error
}

type wallet struct {
	privateKey        cryptotype.PrivKey
	ChainPrefix       string
	ChainID           string
	DialURL           string
	InterfaceRegistry codectypes.InterfaceRegistry
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

func NewWalletWithMnemonic(mnemonic string, chainPrefix string, chainID string, dialURL string, interfaceRegistry codectypes.InterfaceRegistry, options ...MnemonicDeriveOptionHandler) (Wallet, error) {
	deriveFn := hd.Secp256k1.Derive()
	option := &MnemonicDeriveOption{
		BIP39Passphrase: "",
		HDPath:          hd.CreateHDPath(sdk.CoinType, 0, 0).String(),
	}

	for _, o := range options {
		o(option)
	}
	privBytes, err := deriveFn(mnemonic, option.BIP39Passphrase, option.HDPath)

	if err != nil {
		return nil, err
	}

	return &wallet{
		privateKey:        &secp256k1.PrivKey{Key: privBytes},
		ChainPrefix:       chainPrefix,
		ChainID:           chainID,
		DialURL:           dialURL,
		InterfaceRegistry: interfaceRegistry,
	}, nil
}

func NewWallet(privateKey string, chainPrefix string, chainID string, dialURL string, interfaceRegistry codectypes.InterfaceRegistry) (Wallet, error) {
	privBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	return &wallet{
		privateKey:        &secp256k1.PrivKey{Key: privBytes},
		ChainPrefix:       chainPrefix,
		ChainID:           chainID,
		DialURL:           dialURL,
		InterfaceRegistry: interfaceRegistry,
	}, nil
}

func (w wallet) createTxConfig() client.TxConfig {
	cdc := codec.NewProtoCodec(w.InterfaceRegistry)
	return authtx.NewTxConfig(cdc, []txsigning.SignMode{txsigning.SignMode_SIGN_MODE_DIRECT})
}

func (w wallet) GetAddress() (string, error) {
	return libs.ConvertPubKeyToBech32Address(w.privateKey.PubKey(), w.ChainPrefix)
}

func (w wallet) Dial() *grpc.ClientConn {
	conn, err := grpc.Dial(w.DialURL, grpc.WithInsecure()) // nolint: staticcheck
	if err != nil {
		panic(err)
	}
	return conn
}

func (w wallet) GetAccountInfo() (*AccountInfo, error) {
	fromAddress, err := w.GetAddress()
	if err != nil {
		return nil, err
	}

	conn := w.Dial()
	defer conn.Close()

	cli := accounttypes.NewQueryClient(conn)
	res, err := cli.Account(context.Background(), &accounttypes.QueryAccountRequest{Address: fromAddress})
	if err != nil {
		return nil, err
	}

	baseAccount := new(accounttypes.BaseAccount)
	if err := baseAccount.Unmarshal(res.Account.Value); err != nil {
		return nil, err
	}

	return &AccountInfo{
		SequenceNumber: baseAccount.GetSequence(),
		AccountNumber:  baseAccount.GetAccountNumber(),
	}, nil
}

func (w wallet) CreateSignedRawTx(msg sdk.Msg, accountInfo AccountInfo) ([]byte, error) {
	txConfig := w.createTxConfig()
	txBuilder := txConfig.NewTxBuilder()
	payer, _ := w.GetAddress()

	if err := txBuilder.SetMsgs(msg); err != nil {
		return nil, err
	}
	txBuilder.SetGasLimit(100000)
	txBuilder.SetFeePayer(sdk.MustAccAddressFromBech32(payer))
	txBuilder.SetFeeAmount(sdk.Coins{{
		Denom:  "umito",
		Amount: sdk.NewInt(1000),
	}})

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

	if err := txBuilder.SetSignatures(signature); err != nil {
		return nil, err
	}
	return txConfig.TxEncoder()(txBuilder.GetTx())
}

func (w wallet) BroadcastRawTx(rawTxByte []byte) error {
	conn := w.Dial()
	defer conn.Close()
	txClient := txtypes.NewServiceClient(conn)
	resp, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{Mode: txtypes.BroadcastMode_BROADCAST_MODE_BLOCK, TxBytes: rawTxByte},
	)
	if err != nil {
		return err
	}

	fmt.Println(resp.String())
	return nil
}

func (w wallet) BroadcastMsg(msg sdk.Msg) error {
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
