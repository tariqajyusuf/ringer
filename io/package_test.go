package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocatePackage_KnownPackage(t *testing.T) {
	packageDataDir = "../data/packages"
	t.Cleanup(func() { packageDataDir = "data/packages" })

	pkg, err := LocatePackage("git")
	require.NoError(t, err)
	require.NotNil(t, pkg)
	assert.Equal(t, "Git", pkg.Name)
	assert.NotEmpty(t, pkg.Platforms["homebrew"].PackageName)
}

func TestLocatePackage_MissingPackage(t *testing.T) {
	packageDataDir = "../data/packages"
	t.Cleanup(func() { packageDataDir = "data/packages" })

	pkg, err := LocatePackage("__nonexistent__")
	assert.Error(t, err)
	assert.Nil(t, pkg)
}

func TestLocatePackage_MalformedYAML(t *testing.T) {
	dir := t.TempDir()
	packageDataDir = dir
	t.Cleanup(func() { packageDataDir = "data/packages" })

	err := os.WriteFile(filepath.Join(dir, "bad.package.yaml"), []byte(":\tinvalid: yaml::\n"), 0600)
	require.NoError(t, err)

	pkg, err := LocatePackage("bad")
	assert.Error(t, err)
	_ = pkg
}

func TestLocatePackage_MultiPlatform(t *testing.T) {
	packageDataDir = "../data/packages"
	t.Cleanup(func() { packageDataDir = "data/packages" })

	pkg, err := LocatePackage("git")
	require.NoError(t, err)
	assert.NotEmpty(t, pkg.Platforms["homebrew"].PackageName)
	assert.NotEmpty(t, pkg.Platforms["windows"].PackageName)
}
