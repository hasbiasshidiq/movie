package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"movie/x/movie/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group movie queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListMovie())
	cmd.AddCommand(CmdShowMovie())
	cmd.AddCommand(CmdListReview())
	cmd.AddCommand(CmdShowReview())
	cmd.AddCommand(CmdListTittleAllocation())
	cmd.AddCommand(CmdShowTittleAllocation())
	cmd.AddCommand(CmdListReviewsAllocation())
	cmd.AddCommand(CmdShowReviewsAllocation())
	// this line is used by starport scaffolding # 1

	return cmd
}
