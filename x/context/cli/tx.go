package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/many-things/mitosis/x/context/server"
	"github.com/tendermint/tendermint/libs/os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/many-things/mitosis/x/context/types"
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
		SignerReadyCmd(),
		RegisterCosmosSignerCmd(),
		RegisterEVMSignerCmd(),
	)

	return cmd
}

func SignerReadyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signer-ready [event payload]",
		Short: "notifies signer is ready",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgSignerReady)
			clientCtx.Codec.MustUnmarshalJSON(os.MustReadFile(args[0]), msg)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
		SilenceUsage: true,
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func RegisterCosmosSignerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-cosmos-signer [event payload]",
		Short: "Registers Cosmos Signer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgRegisterCosmosSigner)
			clientCtx.Codec.MustUnmarshalJSON(os.MustReadFile(args[0]), msg)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
		SilenceUsage: true,
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func RegisterEVMSignerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-evm-signer [event payload]",
		Short: "Register EVM Signer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgRegisterEVMSigner)
			clientCtx.Codec.MustUnmarshalJSON(os.MustReadFile(args[0]), msg)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
		SilenceUsage: true,
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
