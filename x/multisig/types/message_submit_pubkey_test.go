package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSubmitPubkey_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmitPubkey
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmitPubkey{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSubmitPubkey{
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
