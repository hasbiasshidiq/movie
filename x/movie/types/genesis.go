package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MovieList:            []Movie{},
		ReviewList:           []Review{},
		TittleAllocationList: []TittleAllocation{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in movie
	movieIdMap := make(map[uint64]bool)
	movieCount := gs.GetMovieCount()
	for _, elem := range gs.MovieList {
		if _, ok := movieIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for movie")
		}
		if elem.Id >= movieCount {
			return fmt.Errorf("movie id should be lower or equal than the last id")
		}
		movieIdMap[elem.Id] = true
	}
	// Check for duplicated ID in review
	reviewIdMap := make(map[uint64]bool)
	reviewCount := gs.GetReviewCount()
	for _, elem := range gs.ReviewList {
		if _, ok := reviewIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for review")
		}
		if elem.Id >= reviewCount {
			return fmt.Errorf("review id should be lower or equal than the last id")
		}
		reviewIdMap[elem.Id] = true
	}
	// Check for duplicated index in tittleAllocation
	tittleAllocationIndexMap := make(map[string]struct{})

	for _, elem := range gs.TittleAllocationList {
		index := string(TittleAllocationKey(elem.MovieTitle))
		if _, ok := tittleAllocationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tittleAllocation")
		}
		tittleAllocationIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
