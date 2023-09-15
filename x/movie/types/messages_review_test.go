package types

import (
	"testing"

	"movie/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateReview_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateReview
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateReview{
				Creator: "invalid_address", Star: 2,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateReview{
				Creator: sample.AccAddress(), Star: 2,
			},
		}, {
			name: "Invalid Star Value",
			msg: MsgCreateReview{
				Creator: sample.AccAddress(), Star: 10,
			},
			err: ErrInvalidValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateReview_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateReview
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateReview{
				Creator: "invalid_address", Star: 3,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateReview{
				Creator: sample.AccAddress(), Star: 4,
			},
		}, {
			name: "Invalid Star Value",
			msg: MsgUpdateReview{
				Creator: sample.AccAddress(), Star: 10,
			},
			err: ErrInvalidValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteReview_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteReview
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteReview{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteReview{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
