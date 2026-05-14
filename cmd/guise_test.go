package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGuiseCmd_NoArgs(t *testing.T) {
	out := captureStdout(t, func() {
		rootCmd.SetArgs([]string{"guise"})
		rootCmd.Execute() //nolint:errcheck
	})
	assert.Contains(t, out, "Please provide a guise file")
}

func TestGuiseCmd_FileNotFound(t *testing.T) {
	out := captureStdout(t, func() {
		rootCmd.SetArgs([]string{"guise", "/tmp/__ringer_nonexistent__.yml"})
		rootCmd.Execute() //nolint:errcheck
	})
	assert.Contains(t, out, "Could not read guise file")
}

func TestGuiseCmd_MalformedYAML(t *testing.T) {
	f, err := os.CreateTemp("", "ringer-test-*.yml")
	require.NoError(t, err)
	t.Cleanup(func() { _ = os.Remove(f.Name()) })
	_, err = f.WriteString("packages: [unclosed\n")
	require.NoError(t, err)
	require.NoError(t, f.Close())

	out := captureStdout(t, func() {
		rootCmd.SetArgs([]string{"guise", filepath.Clean(f.Name())})
		rootCmd.Execute() //nolint:errcheck
	})
	assert.Contains(t, out, "Could not parse guise file")
}
