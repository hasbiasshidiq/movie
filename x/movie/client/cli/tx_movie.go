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

func CmdCreateMovie() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-movie [title] [plot] [year] [genre] [language] [is-published]",
		Short: "Create a new movie",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argPlot := args[1]
			argYear, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argGenre := args[3]
			argLanguage := args[4]
			argIsPublished, err := cast.ToBoolE(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMovie(clientCtx.GetFromAddress().String(), argTitle, argPlot, argYear, argGenre, argLanguage, argIsPublished)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateMovie() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-movie [id] [title] [plot] [year] [genre] [language] [is-published]",
		Short: "Update a movie",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argTitle := args[1]

			argPlot := args[2]

			argYear, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argGenre := args[4]

			argLanguage := args[5]

			argIsPublished, err := cast.ToBoolE(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMovie(clientCtx.GetFromAddress().String(), id, argTitle, argPlot, argYear, argGenre, argLanguage, argIsPublished)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteMovie() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-movie [id]",
		Short: "Delete a movie by id",
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

			msg := types.NewMsgDeleteMovie(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
