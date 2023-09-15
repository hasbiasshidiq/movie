package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ReviewsAllocationKeyPrefix is the prefix to retrieve all ReviewsAllocation
	ReviewsAllocationKeyPrefix = "ReviewsAllocation/value/"
)

// ReviewsAllocationKey returns the store key to retrieve a ReviewsAllocation from the index fields
func ReviewsAllocationKey(
	movieId uint64,
) []byte {
	var key []byte

	movieIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(movieIdBytes, movieId)
	key = append(key, movieIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
