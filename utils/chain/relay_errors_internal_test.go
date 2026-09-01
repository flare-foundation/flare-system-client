package chain

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

// The selectors must match the ERR_* constants Relay.sol folds into its
// assembly; a typo in relayErrorNames would silently stop matching on chain.
func TestRelayErrorSelectors(t *testing.T) {
	for name, selector := range map[string]string{
		"AlreadyRelayed()":       "0xd0ebeb4b",
		"TooShortMessage()":      "0x43a69646",
		"MessageTooOld()":        "0x4ed02d0d",
		"WrongSignature()":       "0x356a4418",
		"NoRandomNumber()":       "0xd76adcd1",
		"InvalidVotingRoundId()": "0x01ed5f84",
	} {
		got, ok := relayErrorName(hexutil.MustDecode(selector))
		require.True(t, ok, selector)
		require.Equal(t, name, got)
	}

	_, ok := relayErrorName([]byte{0xde, 0xad, 0xbe, 0xef})
	require.False(t, ok)

	_, ok = relayErrorName([]byte{0xd0, 0xeb, 0xeb})
	require.False(t, ok)
}

func TestDecodeRevertCustomError(t *testing.T) {
	alreadyRelayed := hexutil.MustDecode("0xd0ebeb4b")

	t.Run("bare selector", func(t *testing.T) {
		reason, err := decodeRevert(alreadyRelayed)
		require.NoError(t, err)
		require.Equal(t, "AlreadyRelayed()", reason)
	})

	t.Run("selector with trailing bytes", func(t *testing.T) {
		reason, err := decodeRevert(append(alreadyRelayed, 0x01, 0x02))
		require.NoError(t, err)
		require.Equal(t, "AlreadyRelayed()", reason)
	})

	t.Run("unknown selector stays undecodable", func(t *testing.T) {
		_, err := decodeRevert([]byte{0xde, 0xad, 0xbe, 0xef})
		require.ErrorIs(t, err, errRevertUndecodable)
	})

	t.Run("Error(string) still decodes", func(t *testing.T) {
		reason, err := decodeRevert(encodeRevertReason(t, "Already relayed"))
		require.NoError(t, err)
		require.Equal(t, "Already relayed", reason)
	})
}

func TestAnnotateRevert(t *testing.T) {
	t.Run("custom error is appended", func(t *testing.T) {
		err := annotateRevert(stubDataError{msg: "execution reverted", data: "0xd0ebeb4b"})
		require.ErrorContains(t, err, "AlreadyRelayed()")
	})

	t.Run("reason already in the message is not repeated", func(t *testing.T) {
		revert := hexutil.Encode(encodeRevertReason(t, "Already relayed"))
		err := annotateRevert(stubDataError{msg: "execution reverted: Already relayed", data: revert})
		require.Equal(t, "execution reverted: Already relayed", err.Error())
	})

	t.Run("undecodable revert is left alone", func(t *testing.T) {
		err := annotateRevert(stubDataError{msg: "execution reverted", data: "0xdeadbeef"})
		require.Equal(t, "execution reverted", err.Error())
	})

	t.Run("plain error and nil pass through", func(t *testing.T) {
		plain := errors.New("connection reset")
		require.Equal(t, plain, annotateRevert(plain))
		require.NoError(t, annotateRevert(nil))
	})
}
