package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"movie/x/movie/types"
)

// SetTittleAllocation set a specific tittleAllocation in the store from its index
func (k Keeper) SetTittleAllocation(ctx sdk.Context, tittleAllocation types.TittleAllocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TittleAllocationKeyPrefix))
	b := k.cdc.MustMarshal(&tittleAllocation)
	store.Set(types.TittleAllocationKey(
		tittleAllocation.MovieTitle,
	), b)
}

// GetTittleAllocation returns a tittleAllocation from its index
func (k Keeper) GetTittleAllocation(
	ctx sdk.Context,
	movieTitle string,

) (val types.TittleAllocation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TittleAllocationKeyPrefix))

	b := store.Get(types.TittleAllocationKey(
		movieTitle,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTittleAllocation removes a tittleAllocation from the store
func (k Keeper) RemoveTittleAllocation(
	ctx sdk.Context,
	movieTitle string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TittleAllocationKeyPrefix))
	store.Delete(types.TittleAllocationKey(
		movieTitle,
	))
}

// GetAllTittleAllocation returns all tittleAllocation
func (k Keeper) GetAllTittleAllocation(ctx sdk.Context) (list []types.TittleAllocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TittleAllocationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TittleAllocation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
