package finalizer

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromSignedPayloadEmpty(t *testing.T) {
	cases := []struct {
		name    string
		payload []byte
	}{
		{"nil", nil},
		{"empty", []byte{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s submitSignaturesPayload
			err := s.FromSignedPayload(payloadMessage{payload: tc.payload})
			require.Error(t, err)
		})
	}
}

func TestExtractPayloadsZeroLengthDoesNotPanic(t *testing.T) {
	// 4-byte selector + 1-byte protocol + 4-byte votingRound + 2-byte length=0
	data := make([]byte, 4+1+4+2)
	binary.BigEndian.PutUint32(data[5:9], 1)
	binary.BigEndian.PutUint16(data[9:11], 0)

	payloads, err := ExtractPayloads(data)
	require.NoError(t, err)
	require.Len(t, payloads, 1)

	var s submitSignaturesPayload
	err = s.FromSignedPayload(payloads[0])
	require.Error(t, err)
}
