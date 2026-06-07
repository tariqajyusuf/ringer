package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ringerio "github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system"
)

func TestAddCmd_NoArgs(t *testing.T) {
	out := captureStdout(t, func() {
		rootCmd.SetArgs([]string{"add"})
		rootCmd.Execute() //nolint:errcheck
	})
	assert.Contains(t, out, "Please provide a package name to add")
}


func TestAddHelper_OSBlocked(t *testing.T) {
	dir := t.TempDir()
	ringerio.PackageDataDir = dir
	t.Cleanup(func() { ringerio.PackageDataDir = "data/packages" })

	content := []byte(`name: "macOS Only App"
description: "Only runs on macOS."
os:
  - macos
platforms:
  homebrew:
    name: "macos-only-app"
`)
	err := os.WriteFile(filepath.Join(dir, "macos-only-app.package.yaml"), content, 0600)
	require.NoError(t, err)

	pkg, err := ringerio.LocatePackage("macos-only-app")
	require.NoError(t, err)

	sysErr := pkg.CheckOSAllowed(system.Linux)
	require.Error(t, sysErr)
	assert.Contains(t, sysErr.Error(), "not supported")
}
