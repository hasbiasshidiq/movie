package movie

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"movie/x/movie/keeper"
	"movie/x/movie/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the movie
	for _, elem := range genState.MovieList {
		k.SetMovie(ctx, elem)
	}

	// Set movie count
	k.SetMovieCount(ctx, genState.MovieCount)
	// Set all the review
	for _, elem := range genState.ReviewList {
		k.SetReview(ctx, elem)
	}

	// Set review count
	k.SetReviewCount(ctx, genState.ReviewCount)
	// Set all the tittleAllocation
	for _, elem := range genState.TittleAllocationList {
		k.SetTittleAllocation(ctx, elem)
	}
	// Set all the reviewsAllocation
	for _, elem := range genState.ReviewsAllocationList {
		k.SetReviewsAllocation(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MovieList = k.GetAllMovie(ctx)
	genesis.MovieCount = k.GetMovieCount(ctx)
	genesis.ReviewList = k.GetAllReview(ctx)
	genesis.ReviewCount = k.GetReviewCount(ctx)
	genesis.TittleAllocationList = k.GetAllTittleAllocation(ctx)
	genesis.ReviewsAllocationList = k.GetAllReviewsAllocation(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
