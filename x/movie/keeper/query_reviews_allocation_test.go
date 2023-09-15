package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "movie/testutil/keeper"
	"movie/testutil/nullify"
	"movie/x/movie/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestReviewsAllocationQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNReviewsAllocation(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetReviewsAllocationRequest
		response *types.QueryGetReviewsAllocationResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetReviewsAllocationRequest{
				MovieId: msgs[0].MovieId,
			},
			response: &types.QueryGetReviewsAllocationResponse{ReviewsAllocation: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetReviewsAllocationRequest{
				MovieId: msgs[1].MovieId,
			},
			response: &types.QueryGetReviewsAllocationResponse{ReviewsAllocation: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetReviewsAllocationRequest{
				MovieId: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ReviewsAllocation(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestReviewsAllocationQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNReviewsAllocation(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllReviewsAllocationRequest {
		return &types.QueryAllReviewsAllocationRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ReviewsAllocationAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ReviewsAllocation), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ReviewsAllocation),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ReviewsAllocationAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ReviewsAllocation), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ReviewsAllocation),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ReviewsAllocationAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ReviewsAllocation),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ReviewsAllocationAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
