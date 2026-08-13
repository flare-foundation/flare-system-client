package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// newChainIDRPCStub answers eth_chainId with the given hex id.
func newChainIDRPCStub(t *testing.T, hexID string) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		body, _ := io.ReadAll(r.Body)
		require.NoError(t, json.Unmarshal(body, &req))
		require.Equal(t, "eth_chainId", req.Method)

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":%q}`, req.ID, hexID)
	}))
	t.Cleanup(srv.Close)

	return srv
}

func TestVerifyChainID(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		srv := newChainIDRPCStub(t, "0x72") // 114
		cfg := &Chain{ChainID: 114, EthRPCURL: srv.URL}
		require.NoError(t, cfg.VerifyChainID(context.Background()))
	})

	t.Run("mismatch is ErrChainIDMismatch and names both ids", func(t *testing.T) {
		srv := newChainIDRPCStub(t, "0xe") // 14
		cfg := &Chain{ChainID: 114, EthRPCURL: srv.URL}
		err := cfg.VerifyChainID(context.Background())
		require.ErrorIs(t, err, ErrChainIDMismatch)
		require.ErrorContains(t, err, "config has 114")
		require.ErrorContains(t, err, "node reports 14")
	})

	t.Run("node id beyond int64 is a mismatch, not a panic", func(t *testing.T) {
		srv := newChainIDRPCStub(t, "0x1ffffffffffffffff")
		cfg := &Chain{ChainID: 114, EthRPCURL: srv.URL}
		require.ErrorIs(t, cfg.VerifyChainID(context.Background()), ErrChainIDMismatch)
	})

	t.Run("unreachable node errors but is not a mismatch", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
		url := srv.URL
		srv.Close()

		cfg := &Chain{ChainID: 114, EthRPCURL: url}
		err := cfg.VerifyChainID(context.Background())
		require.Error(t, err)
		require.NotErrorIs(t, err, ErrChainIDMismatch)
	})
}
