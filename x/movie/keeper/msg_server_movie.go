package keeper

import (
	"context"
	"fmt"

	"movie/x/movie/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMovie(goCtx context.Context, msg *types.MsgCreateMovie) (*types.MsgCreateMovieResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var movie = types.Movie{
		Creator:     msg.Creator,
		Title:       msg.Title,
		Plot:        msg.Plot,
		Year:        msg.Year,
		Genre:       msg.Genre,
		Language:    msg.Language,
		IsPublished: msg.IsPublished,
	}

	_, isExist := k.GetTittleAllocation(ctx, msg.Title)
	if isExist {
		return nil, types.ErrMovieTitleAlreadyExist
	}

	id := k.AppendMovie(ctx, movie)

	k.SetTittleAllocation(ctx, types.TittleAllocation{MovieTitle: msg.Title, MovieId: id})

	return &types.MsgCreateMovieResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMovie(goCtx context.Context, msg *types.MsgUpdateMovie) (*types.MsgUpdateMovieResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var movie = types.Movie{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Title:       msg.Title,
		Plot:        msg.Plot,
		Year:        msg.Year,
		Genre:       msg.Genre,
		Language:    msg.Language,
		IsPublished: msg.IsPublished,
	}

	// Checks that the element exists
	val, found := k.GetMovie(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	allocatedMovieID, isExist := k.GetTittleAllocation(ctx, msg.Title)
	if allocatedMovieID.MovieId != msg.Id && isExist {
		return nil, types.ErrMovieTitleAlreadyExist
	}

	k.SetMovie(ctx, movie)
	k.SetTittleAllocation(ctx, types.TittleAllocation{MovieTitle: msg.Title, MovieId: msg.Id})

	return &types.MsgUpdateMovieResponse{}, nil
}

func (k msgServer) DeleteMovie(goCtx context.Context, msg *types.MsgDeleteMovie) (*types.MsgDeleteMovieResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMovie(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if val.IsPublished {
		return nil, types.ErrCannotDeletePublishedMovie
	}

	reviews, isReviewExistOnMovie := k.GetReviewsAllocation(ctx, val.Id)
	if isReviewExistOnMovie && len(reviews.ReviewIds) > 0 {
		return nil, types.ErrCannotDeleteReviewedMovie
	}

	k.RemoveMovie(ctx, msg.Id)

	return &types.MsgDeleteMovieResponse{}, nil
}
