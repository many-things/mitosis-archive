package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/many-things/mitosis/x/multisig/server"
	"github.com/tendermint/tendermint/libs/os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/many-things/mitosis/x/multisig/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		StartKeygenCmd(),
		SubmitPubKeyCmd(),
		SubmitSignatureCmd(),
	)

	return cmd
}

func StartKeygenCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start-keygen [payload]",
		Short: "Start Keygen",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgStartKeygen)
			clientCtx.Codec.MustUnmarshalJSON(os.MustReadFile(args[0]), msg)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

func SubmitPubKeyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "submit-pubkey [payload]",
		Short: "Submit PubKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgSubmitPubkey)
			clientCtx.Codec.MustUnmarshalJSON(os.MustReadFile(args[0]), msg)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

func SubmitSignatureCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "submit-signature [payload]",
		Short: "Submit Signature",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgSubmitSignature)
			clientCtx.Codec.MustUnmarshalJSON(os.MustReadFile(args[0]), msg)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
