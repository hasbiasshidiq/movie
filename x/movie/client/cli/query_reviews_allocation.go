package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/spf13/cast"
	"movie/x/movie/types"
)

func CmdListReviewsAllocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-reviews-allocation",
		Short: "list all reviewsAllocation",
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

			params := &types.QueryAllReviewsAllocationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ReviewsAllocationAll(cmd.Context(), params)
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

func CmdShowReviewsAllocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-reviews-allocation [movie-id]",
		Short: "shows a reviewsAllocation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argMovieId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetReviewsAllocationRequest{
				MovieId: argMovieId,
			}

			res, err := queryClient.ReviewsAllocation(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
