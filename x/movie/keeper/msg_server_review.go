package keeper

import (
	"context"
	"fmt"

	"movie/x/movie/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateReview(goCtx context.Context, msg *types.MsgCreateReview) (*types.MsgCreateReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var review = types.Review{
		Creator: msg.Creator,
		MovieId: msg.MovieId,
		Star:    msg.Star,
		Comment: msg.Comment,
	}

	_, isFound := k.GetMovie(ctx, msg.MovieId)
	if !isFound {
		return nil, sdkerrors.Wrap(types.ErrMovieDoesNotExist, fmt.Sprintf("Can't create review since movie with id %d doesn't exist", msg.MovieId))
	}

	id := k.AppendReview(
		ctx,
		review,
	)

	k.UpdateMovieToReviewsMap(ctx, msg.MovieId, id)

	return &types.MsgCreateReviewResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMovieToReviewsMap(ctx sdk.Context, movieId uint64, newReviewId uint64) {

	reviewsAllocation, _ := k.GetReviewsAllocation(ctx, movieId)
	reviewIds := reviewsAllocation.ReviewIds

	reviewIds = append(reviewIds, newReviewId)
	reviewsAllocation = types.ReviewsAllocation{MovieId: movieId, ReviewIds: reviewIds}

	k.SetReviewsAllocation(ctx, reviewsAllocation)
}

func (k msgServer) UpdateReview(goCtx context.Context, msg *types.MsgUpdateReview) (*types.MsgUpdateReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var review = types.Review{
		Creator: msg.Creator,
		Id:      msg.Id,
		MovieId: msg.MovieId,
		Star:    msg.Star,
		Comment: msg.Comment,
	}

	// Checks that the element exists
	val, found := k.GetReview(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Checks if movie exist
	_, isFound := k.GetMovie(ctx, msg.MovieId)
	if !isFound {
		return nil, sdkerrors.Wrap(types.ErrMovieDoesNotExist, fmt.Sprintf("Can't update review since movie with id %d doesn't exist", msg.MovieId))
	}

	k.SetReview(ctx, review)

	return &types.MsgUpdateReviewResponse{}, nil
}

func (k msgServer) DeleteReview(goCtx context.Context, msg *types.MsgDeleteReview) (*types.MsgDeleteReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetReview(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveReview(ctx, msg.Id)

	return &types.MsgDeleteReviewResponse{}, nil
}
