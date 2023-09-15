package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "movie/testutil/keeper"
	"movie/x/movie"
	"movie/x/movie/keeper"
	"movie/x/movie/types"
)

func setupMsgServerReview(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MovieKeeper(t)
	movie.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestReviewMsgServerCreate(t *testing.T) {

	t.Run("Test Increment Id", func(t *testing.T) {
		srv, ctx := setupMsgServerReview(t)
		creator := "A"

		createdMovie, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
		require.NoError(t, err)

		resp, err := srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: createdMovie.Id})
		require.NoError(t, err)
		require.Equal(t, 0, int(resp.Id))

		resp, err = srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: createdMovie.Id})
		require.NoError(t, err)
		require.Equal(t, 1, int(resp.Id))

		resp, err = srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: createdMovie.Id})
		require.NoError(t, err)
		require.Equal(t, 2, int(resp.Id))

	})
	t.Run("Test Movie Id doesn't exist", func(t *testing.T) {
		srv, ctx := setupMsgServerReview(t)
		creator := "A"

		createdMovie, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
		require.NoError(t, err)

		movieId := createdMovie.Id + 10
		_, err = srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: movieId})
		require.Equal(t, fmt.Sprintf("Can't create review since movie with id %d doesn't exist: movie doesn't exist", movieId), err.Error())

	})
}

func TestReviewMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateReview
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateReview{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateReview{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateReview{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServerReview(t)

			createdMovie, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
			require.NoError(t, err)

			_, err = srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: createdMovie.Id})
			require.NoError(t, err)

			tc.request.MovieId = createdMovie.Id
			_, err = srv.UpdateReview(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}

	t.Run("Test Movie Id doesn't exist", func(t *testing.T) {
		srv, ctx := setupMsgServerReview(t)
		creator := "A"

		createdMovie, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
		require.NoError(t, err)

		reviewCreate := &types.MsgCreateReview{Creator: creator, MovieId: createdMovie.Id}
		createdReview, err := srv.CreateReview(ctx, reviewCreate)
		require.NoError(t, err)

		movieId := createdMovie.Id + 10
		reviewUpdate := &types.MsgUpdateReview{Creator: creator, Id: createdReview.Id, MovieId: movieId}
		_, err = srv.UpdateReview(ctx, reviewUpdate)

		require.Equal(t, fmt.Sprintf("Can't update review since movie with id %d doesn't exist: movie doesn't exist", movieId), err.Error())

	})
}

func TestReviewMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteReview
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteReview{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteReview{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteReview{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServerReview(t)
			createdMovie, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
			require.NoError(t, err)

			_, err = srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: createdMovie.Id})
			require.NoError(t, err)
			_, err = srv.DeleteReview(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
