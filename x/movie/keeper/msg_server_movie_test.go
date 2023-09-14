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

func setupMsgServerMovie(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MovieKeeper(t)
	movie.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
func TestMovieMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServerMovie(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator, Title: fmt.Sprintf("Kimi No Nawa %d", i)})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestMovieMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateMovie
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateMovie{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMovie{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMovie{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServerMovie(t)
			_, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateMovie(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMovieMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteMovie
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteMovie{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteMovie{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteMovie{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServerMovie(t)

			_, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteMovie(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMovieMsgServerCreateDuplicateTitle(t *testing.T) {
	creator := "A"

	t.Run("Teset Duplicate Title", func(t *testing.T) {
		srv, ctx := setupMsgServerMovie(t)

		_, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator, Title: "Kimi No Nawa"})
		require.NoError(t, err)

		movieResponse, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator, Title: "Kimi No Nawa"})

		require.Nil(t, movieResponse)
		require.Equal(t, types.ErrMovieTitleAlreadyExist.Error(), err.Error())
	})

}

func TestMovieMsgServerUpdateDuplicateTitle(t *testing.T) {
	creator := "A"

	t.Run("Teset Duplicate Title", func(t *testing.T) {
		srv, ctx := setupMsgServerMovie(t)

		movieCreate1 := &types.MsgCreateMovie{Creator: creator, Title: "Kimi No Nawa"}
		movieCreate2 := &types.MsgCreateMovie{Creator: creator, Title: "Gintama"}

		createdMovie1, err := srv.CreateMovie(ctx, movieCreate1)
		require.NoError(t, err)

		_, err = srv.CreateMovie(ctx, movieCreate2)
		require.NoError(t, err)

		// Ensure that a duplicate title error does not occur when updating a movie with an ID that corresponds to the title
		_, err = srv.UpdateMovie(ctx, &types.MsgUpdateMovie{Creator: creator, Id: createdMovie1.Id, Title: "Kimi No Nawa"})
		require.NoError(t, err)
		_, err = srv.UpdateMovie(ctx, &types.MsgUpdateMovie{Creator: creator, Id: createdMovie1.Id, Title: "Kimi No Nawa (Updated)"})
		require.NoError(t, err)

		// A movie with title "Gintama" has already been in existence.
		updateResponse, err := srv.UpdateMovie(ctx, &types.MsgUpdateMovie{Creator: creator, Id: createdMovie1.Id, Title: "Gintama"})
		require.Nil(t, updateResponse)
		require.Equal(t, types.ErrMovieTitleAlreadyExist.Error(), err.Error())
	})

}
