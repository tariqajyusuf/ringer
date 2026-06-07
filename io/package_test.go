package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tariqajyusuf/ringer/system"
)

func TestLocatePackage_KnownPackage(t *testing.T) {
	PackageDataDir = "../data/packages"
	t.Cleanup(func() { PackageDataDir = "data/packages" })

	pkg, err := LocatePackage("git")
	require.NoError(t, err)
	require.NotNil(t, pkg)
	assert.Equal(t, "Git", pkg.Name)
	assert.NotEmpty(t, pkg.Platforms["homebrew"].PackageName)
}

func TestLocatePackage_MissingPackage(t *testing.T) {
	PackageDataDir = "../data/packages"
	t.Cleanup(func() { PackageDataDir = "data/packages" })

	pkg, err := LocatePackage("__nonexistent__")
	assert.Error(t, err)
	assert.Nil(t, pkg)
}

func TestLocatePackage_MalformedYAML(t *testing.T) {
	dir := t.TempDir()
	PackageDataDir = dir
	t.Cleanup(func() { PackageDataDir = "data/packages" })

	err := os.WriteFile(filepath.Join(dir, "bad.package.yaml"), []byte(":\tinvalid: yaml::\n"), 0600)
	require.NoError(t, err)

	pkg, err := LocatePackage("bad")
	assert.Error(t, err)
	_ = pkg
}

func TestLocatePackage_MultiPlatform(t *testing.T) {
	PackageDataDir = "../data/packages"
	t.Cleanup(func() { PackageDataDir = "data/packages" })

	pkg, err := LocatePackage("git")
	require.NoError(t, err)
	assert.NotEmpty(t, pkg.Platforms["homebrew"].PackageName)
	assert.NotEmpty(t, pkg.Platforms["windows"].PackageName)
}

func TestLocatePackage_OSFieldParsed(t *testing.T) {
	PackageDataDir = "../data/packages"
	t.Cleanup(func() { PackageDataDir = "data/packages" })

	pkg, err := LocatePackage("iterm")
	require.NoError(t, err)
	assert.Equal(t, []string{"macos"}, pkg.OS)
}

func TestCheckOSAllowed_NoRestriction(t *testing.T) {
	pkg := &Package{Name: "Git"}
	assert.NoError(t, pkg.CheckOSAllowed(system.MacOS))
	assert.NoError(t, pkg.CheckOSAllowed(system.Linux))
	assert.NoError(t, pkg.CheckOSAllowed(system.Windows))
}

func TestCheckOSAllowed_Allowed(t *testing.T) {
	pkg := &Package{Name: "iTerm", OS: []string{"macos"}}
	assert.NoError(t, pkg.CheckOSAllowed(system.MacOS))
}

func TestCheckOSAllowed_Blocked(t *testing.T) {
	pkg := &Package{Name: "iTerm", OS: []string{"macos"}}
	err := pkg.CheckOSAllowed(system.Linux)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "linux")
	assert.Contains(t, err.Error(), "macos")
}

func TestCheckOSAllowed_UnknownKernel(t *testing.T) {
	pkg := &Package{Name: "iTerm", OS: []string{"macos"}}
	assert.Error(t, pkg.CheckOSAllowed(system.Unknown))
}
