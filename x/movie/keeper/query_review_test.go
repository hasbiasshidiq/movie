package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "movie/testutil/keeper"
	"movie/testutil/nullify"
	"movie/x/movie/types"
)

func TestReviewQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNReview(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetReviewRequest
		response *types.QueryGetReviewResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetReviewRequest{Id: msgs[0].Id},
			response: &types.QueryGetReviewResponse{Review: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetReviewRequest{Id: msgs[1].Id},
			response: &types.QueryGetReviewResponse{Review: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetReviewRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Review(wctx, tc.request)
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

func TestReviewQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNReview(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllReviewRequest {
		return &types.QueryAllReviewRequest{
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
			resp, err := keeper.ReviewAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Review), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Review),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ReviewAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Review), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Review),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ReviewAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Review),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ReviewAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
