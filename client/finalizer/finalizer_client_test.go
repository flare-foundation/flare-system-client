package finalizer

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFinalizerClientType0(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := setupTest(0)
	require.NoError(t, err)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Len(t, clients.eth.sentTxs, 1)

	cupaloy.SnapshotT(t, clients.eth.sentTxs[0])
}

func TestFinalizerClientType1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := setupTest(1)
	require.NoError(t, err)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Len(t, clients.eth.sentTxs, 1)

	cupaloy.SnapshotT(t, clients.eth.sentTxs[0])
}

func TestFinalizerClientSendTxErr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clients, err := setupTest(0)
	require.NoError(t, err)

	clients.eth.sendTxErr = errors.New("sendRawTx error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.eth.hasAnyCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

func TestFinalizerClientFetchTxsErr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clients, err := setupTest(0)
	require.NoError(t, err)

	clients.db.fetchTxsErr = errors.New("fetchTxs error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	require.Eventually(
		t, clients.db.hasAnyFetchTxsCalls, 10*time.Second, 100*time.Millisecond,
	)

	cancel()
	err = eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}

func TestFinalizerClientFetchLogsErr(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := setupTest(0)
	require.NoError(t, err)

	clients.db.fetchLogsErr = errors.New("fetchLogs error")

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return clients.finalizer.Run(ctx)
	})

	err = eg.Wait()
	require.True(t, errors.Is(err, clients.db.fetchLogsErr), "unexpected error: %v", err)

	t.Logf("sent transactions: %d", len(clients.eth.sentTxs))
	require.Empty(t, clients.eth.sentTxs)
}
