package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/flare-foundation/flare-system-client/config"

	"github.com/stretchr/testify/require"
)

func parseTOML(t *testing.T, body string) *Client {
	t.Helper()
	path := filepath.Join(t.TempDir(), "cfg.toml")
	require.NoError(t, os.WriteFile(path, []byte(body), 0o600))
	cfg := defaultConfig()
	require.NoError(t, config.ParseConfigFile(cfg, path, false))
	return cfg
}

func TestSubmitEnabledParsing(t *testing.T) {
	t.Run("absent keys default to enabled", func(t *testing.T) {
		cfg := parseTOML(t, `
[submit1]
start_offset = "75s"
[submit2]
start_offset = "20s"
[submit_signatures]
start_offset = "45s"
`)
		require.True(t, cfg.Submit1.Enabled)
		require.True(t, cfg.Submit2.Enabled)
		require.True(t, cfg.SubmitSignatures.Enabled)
	})

	t.Run("explicit opt-out parses, including embedded submit_signatures", func(t *testing.T) {
		cfg := parseTOML(t, `
[submit1]
enabled = false
[submit2]
enabled = true
[submit_signatures]
enabled = false
`)
		require.False(t, cfg.Submit1.Enabled)
		require.True(t, cfg.Submit2.Enabled)
		require.False(t, cfg.SubmitSignatures.Enabled)
	})

	t.Run("legacy enabled = true keys still parse", func(t *testing.T) {
		cfg := parseTOML(t, `
[submit1]
enabled = true
start_offset = "75s"
`)
		require.True(t, cfg.Submit1.Enabled)
	})
}
