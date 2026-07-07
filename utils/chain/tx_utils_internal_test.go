package chain

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

// encodeRevertReason builds revert data as returned by a contract: Error(string) selector + ABI-encoded string.
func encodeRevertReason(t *testing.T, reason string) []byte {
	t.Helper()

	encoded, err := abi.Arguments{{Type: abiString}}.Pack(reason)
	require.NoError(t, err)

	return append(append([]byte{}, errorSig...), encoded...)
}

func TestUnpackError(t *testing.T) {
	tests := []struct {
		name      string
		result    []byte
		want      string
		expectErr bool
	}{
		{
			name:      "valid revert reason",
			result:    encodeRevertReason(t, "nonce too low"),
			want:      "nonce too low",
			expectErr: false,
		},
		{
			name:      "empty result",
			result:    []byte{},
			want:      "<tx result not Error(string)>",
			expectErr: true,
		},
		{
			name:      "result shorter than selector",
			result:    []byte{0x08, 0xc3},
			want:      "<tx result not Error(string)>",
			expectErr: true,
		},
		{
			name:      "unknown selector",
			result:    []byte{0xde, 0xad, 0xbe, 0xef, 0x01, 0x02},
			want:      "<tx result not Error(string)>",
			expectErr: true,
		},
		{
			name:      "valid selector with malformed payload",
			result:    append(append([]byte{}, errorSig...), 0x01, 0x02, 0x03),
			want:      "<invalid tx result>",
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := unpackError(tc.result)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.want, got)
		})
	}
}
