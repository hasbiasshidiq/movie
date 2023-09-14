package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/movie module sentinel errors
var (
	ErrSample                 = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrMovieTitleAlreadyExist = sdkerrors.Register(ModuleName, 1101, "movie with this title is already exist")
)
