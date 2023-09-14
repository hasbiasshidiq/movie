package keeper

import (
	"context"
	"movie/x/movie/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TittleAllocationAll(goCtx context.Context, req *types.QueryAllTittleAllocationRequest) (*types.QueryAllTittleAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tittleAllocations []types.TittleAllocation
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	tittleAllocationStore := prefix.NewStore(store, types.KeyPrefix(types.TittleAllocationKeyPrefix))

	pageRes, err := query.Paginate(tittleAllocationStore, req.Pagination, func(key []byte, value []byte) error {
		var tittleAllocation types.TittleAllocation
		if err := k.cdc.Unmarshal(value, &tittleAllocation); err != nil {
			return err
		}

		tittleAllocations = append(tittleAllocations, tittleAllocation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTittleAllocationResponse{TittleAllocation: tittleAllocations, Pagination: pageRes}, nil
}

func (k Keeper) TittleAllocation(goCtx context.Context, req *types.QueryGetTittleAllocationRequest) (*types.QueryGetTittleAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTittleAllocation(
		ctx,
		req.MovieTitle,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTittleAllocationResponse{TittleAllocation: val}, nil
}
