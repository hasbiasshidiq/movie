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

func createNTittleAllocation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TittleAllocation {
	items := make([]types.TittleAllocation, n)
	for i := range items {
		items[i].MovieTitle = strconv.Itoa(i)

		keeper.SetTittleAllocation(ctx, items[i])
	}
	return items
}

func TestTittleAllocationGet(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNTittleAllocation(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTittleAllocation(ctx,
			item.MovieTitle,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTittleAllocationRemove(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNTittleAllocation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTittleAllocation(ctx,
			item.MovieTitle,
		)
		_, found := keeper.GetTittleAllocation(ctx,
			item.MovieTitle,
		)
		require.False(t, found)
	}
}

func TestTittleAllocationGetAll(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNTittleAllocation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTittleAllocation(ctx)),
	)
}
