package cli

import (
	"github.com/many-things/mitosis/x/context/server"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdContextByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context-by-address",
		Short: "Query ContextByAddress",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := server.NewQueryClient(clientCtx)

			params := &server.QueryContextByAddressRequest{}

			res, err := queryClient.ContextByAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
