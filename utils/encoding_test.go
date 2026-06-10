package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransformSignatureRoundTrip(t *testing.T) {
	rsv := make([]byte, SignatureLength)
	for i := range 64 {
		rsv[i] = byte(i + 1)
	}
	rsv[64] = 1 // V - 27

	vrs, err := TransformSignatureRSVtoVRS(rsv)
	require.NoError(t, err)
	require.Len(t, vrs, SignatureLength)
	require.Equal(t, byte(28), vrs[0])
	require.Equal(t, rsv[:64], vrs[1:])

	back, err := TransformSignatureVRStoRSV(vrs)
	require.NoError(t, err)
	require.Equal(t, rsv, back)
}

func TestTransformSignatureLengthChecks(t *testing.T) {
	for _, length := range []int{0, 1, 64, 66, 130} {
		input := make([]byte, length)

		_, err := TransformSignatureVRStoRSV(input)
		require.Error(t, err, "VRStoRSV must reject length %d", length)

		_, err = TransformSignatureRSVtoVRS(input)
		require.Error(t, err, "RSVtoVRS must reject length %d", length)
	}
}
