package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TittleAllocationKeyPrefix is the prefix to retrieve all TittleAllocation
	TittleAllocationKeyPrefix = "TittleAllocation/value/"
)

// TittleAllocationKey returns the store key to retrieve a TittleAllocation from the index fields
func TittleAllocationKey(
	movieTitle string,
) []byte {
	var key []byte

	movieTitleBytes := []byte(movieTitle)
	key = append(key, movieTitleBytes...)
	key = append(key, []byte("/")...)

	return key
}
