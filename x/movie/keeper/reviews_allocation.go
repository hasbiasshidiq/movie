package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"movie/x/movie/types"
)

// SetReviewsAllocation set a specific reviewsAllocation in the store from its index
func (k Keeper) SetReviewsAllocation(ctx sdk.Context, reviewsAllocation types.ReviewsAllocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewsAllocationKeyPrefix))
	b := k.cdc.MustMarshal(&reviewsAllocation)
	store.Set(types.ReviewsAllocationKey(
		reviewsAllocation.MovieId,
	), b)
}

// GetReviewsAllocation returns a reviewsAllocation from its index
func (k Keeper) GetReviewsAllocation(
	ctx sdk.Context,
	movieId uint64,

) (val types.ReviewsAllocation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewsAllocationKeyPrefix))

	b := store.Get(types.ReviewsAllocationKey(
		movieId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveReviewsAllocation removes a reviewsAllocation from the store
func (k Keeper) RemoveReviewsAllocation(
	ctx sdk.Context,
	movieId uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewsAllocationKeyPrefix))
	store.Delete(types.ReviewsAllocationKey(
		movieId,
	))
}

// GetAllReviewsAllocation returns all reviewsAllocation
func (k Keeper) GetAllReviewsAllocation(ctx sdk.Context) (list []types.ReviewsAllocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewsAllocationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ReviewsAllocation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
