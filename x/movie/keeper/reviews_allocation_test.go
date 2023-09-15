package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "movie/testutil/keeper"
	"movie/testutil/nullify"
	"movie/x/movie/keeper"
	"movie/x/movie/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNReviewsAllocation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ReviewsAllocation {
	items := make([]types.ReviewsAllocation, n)
	for i := range items {
		items[i].MovieId = uint64(i)

		keeper.SetReviewsAllocation(ctx, items[i])
	}
	return items
}

func TestReviewsAllocationGet(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNReviewsAllocation(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetReviewsAllocation(ctx,
			item.MovieId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestReviewsAllocationRemove(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNReviewsAllocation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveReviewsAllocation(ctx,
			item.MovieId,
		)
		_, found := keeper.GetReviewsAllocation(ctx,
			item.MovieId,
		)
		require.False(t, found)
	}
}

func TestReviewsAllocationGetAll(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNReviewsAllocation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllReviewsAllocation(ctx)),
	)
}
