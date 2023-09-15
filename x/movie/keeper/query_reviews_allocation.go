package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"movie/x/movie/types"
)

func (k Keeper) ReviewsAllocationAll(goCtx context.Context, req *types.QueryAllReviewsAllocationRequest) (*types.QueryAllReviewsAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reviewsAllocations []types.ReviewsAllocation
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	reviewsAllocationStore := prefix.NewStore(store, types.KeyPrefix(types.ReviewsAllocationKeyPrefix))

	pageRes, err := query.Paginate(reviewsAllocationStore, req.Pagination, func(key []byte, value []byte) error {
		var reviewsAllocation types.ReviewsAllocation
		if err := k.cdc.Unmarshal(value, &reviewsAllocation); err != nil {
			return err
		}

		reviewsAllocations = append(reviewsAllocations, reviewsAllocation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReviewsAllocationResponse{ReviewsAllocation: reviewsAllocations, Pagination: pageRes}, nil
}

func (k Keeper) ReviewsAllocation(goCtx context.Context, req *types.QueryGetReviewsAllocationRequest) (*types.QueryGetReviewsAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetReviewsAllocation(
		ctx,
		req.MovieId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetReviewsAllocationResponse{ReviewsAllocation: val}, nil
}
