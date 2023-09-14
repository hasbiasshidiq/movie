package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"movie/x/movie/types"
)

func CmdCreateReview() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-review [movie-id] [star] [comment]",
		Short: "Create a new review",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMovieId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argStar, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argComment := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateReview(clientCtx.GetFromAddress().String(), argMovieId, argStar, argComment)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateReview() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-review [id] [movie-id] [star] [comment]",
		Short: "Update a review",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argMovieId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argStar, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			argComment := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateReview(clientCtx.GetFromAddress().String(), id, argMovieId, argStar, argComment)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteReview() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-review [id]",
		Short: "Delete a review by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteReview(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
