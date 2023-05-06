package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mritunjay/x/mritunjay/types"
)

var _ = strconv.Itoa(0)

func CmdSayGoodmorning() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "say-goodmorning [name]",
		Short: "Query say-goodmorning",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqName := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySayGoodMorningRequest{

				Name: reqName,
			}

			res, err := queryClient.SayGoodMorning(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}



