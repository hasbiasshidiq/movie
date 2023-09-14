package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"movie/x/movie/types"
)

func (k Keeper) MovieAll(goCtx context.Context, req *types.QueryAllMovieRequest) (*types.QueryAllMovieResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var movies []types.Movie
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	movieStore := prefix.NewStore(store, types.KeyPrefix(types.MovieKey))

	pageRes, err := query.Paginate(movieStore, req.Pagination, func(key []byte, value []byte) error {
		var movie types.Movie
		if err := k.cdc.Unmarshal(value, &movie); err != nil {
			return err
		}

		movies = append(movies, movie)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMovieResponse{Movie: movies, Pagination: pageRes}, nil
}

func (k Keeper) Movie(goCtx context.Context, req *types.QueryGetMovieRequest) (*types.QueryGetMovieResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	movie, found := k.GetMovie(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMovieResponse{Movie: movie}, nil
}
