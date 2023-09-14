package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"movie/x/movie/types"
)

func CmdListTittleAllocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tittle-allocation",
		Short: "list all tittleAllocation",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTittleAllocationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TittleAllocationAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowTittleAllocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tittle-allocation [movie-title]",
		Short: "shows a tittleAllocation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argMovieTitle := args[0]

			params := &types.QueryGetTittleAllocationRequest{
				MovieTitle: argMovieTitle,
			}

			res, err := queryClient.TittleAllocation(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
