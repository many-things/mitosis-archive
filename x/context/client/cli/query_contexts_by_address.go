package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdContextsByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contexts-by-address",
		Short: "Query ContextsByAddress",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryContextsByAddressRequest{}

			res, err := queryClient.ContextsByAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
