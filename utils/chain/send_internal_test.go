package chain

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// mayBeInMempool must classify ambiguous broadcast errors as possibly-in-mempool
// (so their hash is tracked) and only known synchronous rejections as not.
func TestMayBeInMempool(t *testing.T) {
	// Definitive rejections: the node never admitted the tx.
	require.False(t, mayBeInMempool(errors.New("nonce too low")))
	require.False(t, mayBeInMempool(errors.New("replacement transaction underpriced")))
	require.False(t, mayBeInMempool(errors.New("insufficient funds for gas * price + value")))
	require.False(t, mayBeInMempool(errors.New("intrinsic gas too low")))

	// Ambiguous failures: the tx may already be in the mempool.
	require.True(t, mayBeInMempool(context.DeadlineExceeded))
	require.True(t, mayBeInMempool(errors.New("read tcp 1.2.3.4:443: connection reset by peer")))
	require.True(t, mayBeInMempool(errors.New("unexpected EOF")))
	require.True(t, mayBeInMempool(errors.New("http2: server sent GOAWAY and closed the connection")))
	require.True(t, mayBeInMempool(errors.New("502 Bad Gateway")))
}

func TestIsAlreadyKnown(t *testing.T) {
	require.True(t, isAlreadyKnown(errors.New("already known")))
	require.False(t, isAlreadyKnown(errors.New("nonce too low")))
	require.False(t, isAlreadyKnown(nil))
}
