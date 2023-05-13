package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/many-things/mitosis/x/event/server"
	"github.com/tendermint/tendermint/libs/os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/many-things/mitosis/x/event/types"
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

	cmd.AddCommand(SubmitEventCmd())
	cmd.AddCommand(RegisterProxyCmd())
	cmd.AddCommand(RegisterChainCmd())

	return cmd
}

func SubmitEventCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "submit-event [event payload]",
		Short:   "Submit Event",
		Aliases: []string{"submit", "sm", "s"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgSubmitEvent)
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

func RegisterProxyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register-proxy [event payload]",
		Short:   "Register Proxy",
		Aliases: []string{"reg-proxy", "rg"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgRegisterProxy)
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

func RegisterChainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register-chain [event payload]",
		Short:   "Register Chain",
		Aliases: []string{"reg-chain", "rc"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(server.MsgRegisterChain)
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
