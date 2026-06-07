package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withTempConfigDir(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	orig, err := os.UserHomeDir()
	require.NoError(t, err)
	t.Setenv("HOME", dir)
	t.Cleanup(func() { t.Setenv("HOME", orig) })
}

func TestLoadConfig_MissingFile(t *testing.T) {
	withTempConfigDir(t)
	cfg, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "", cfg.PreferredPlatform)
}

func TestLoadConfig_ValidFile(t *testing.T) {
	withTempConfigDir(t)
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".ringer")
	require.NoError(t, os.MkdirAll(dir, 0o700))
	require.NoError(t, os.WriteFile(
		filepath.Join(dir, "config.yaml"),
		[]byte("preferred_platform: homebrew\n"),
		0o600,
	))

	cfg, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "homebrew", cfg.PreferredPlatform)
}

func TestSaveAndLoadConfig_RoundTrip(t *testing.T) {
	withTempConfigDir(t)
	cfg := &Config{PreferredPlatform: "winget"}
	require.NoError(t, SaveConfig(cfg))

	loaded, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "winget", loaded.PreferredPlatform)
}

func TestSaveConfig_CreatesDirectory(t *testing.T) {
	withTempConfigDir(t)
	home, _ := os.UserHomeDir()
	cfg := &Config{PreferredPlatform: "homebrew"}
	require.NoError(t, SaveConfig(cfg))
	_, err := os.Stat(filepath.Join(home, ".ringer", "config.yaml"))
	assert.NoError(t, err)
}

func TestLoadConfig_MalformedYAML(t *testing.T) {
	withTempConfigDir(t)
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".ringer")
	require.NoError(t, os.MkdirAll(dir, 0o700))
	require.NoError(t, os.WriteFile(
		filepath.Join(dir, "config.yaml"),
		[]byte("preferred_platform: [unclosed\n"),
		0o600,
	))

	_, err := LoadConfig()
	assert.Error(t, err)
}
