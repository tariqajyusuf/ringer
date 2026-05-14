package platforms

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tariqajyusuf/ringer/system"
)

func TestWinget_PlatformStub(t *testing.T) {
	assert.Equal(t, "winget", Winget{}.PlatformStub())
}

func TestWinget_EnabledForSystem_Windows(t *testing.T) {
	assert.True(t, Winget{}.EnabledForSystem(system.SystemInfo{Kernel: system.Windows}))
}

func TestWinget_EnabledForSystem_MacOS(t *testing.T) {
	assert.False(t, Winget{}.EnabledForSystem(system.SystemInfo{Kernel: system.MacOS}))
}

func TestWinget_EnabledForSystem_Linux(t *testing.T) {
	assert.False(t, Winget{}.EnabledForSystem(system.SystemInfo{Kernel: system.Linux}))
}

func TestWinget_ParseOutput_NotFound(t *testing.T) {
	w := Winget{}
	output := []byte("No package found matching input criteria.\n")
	err := w.parseOutput(output)
	require.Error(t, err)
	var target *PackageNotFound
	assert.True(t, errors.As(err, &target))
}

func TestWinget_ParseOutput_GenericError(t *testing.T) {
	w := Winget{}
	output := []byte("An unexpected error occurred.\n")
	err := w.parseOutput(output)
	require.Error(t, err)
	var target *InstallError
	assert.True(t, errors.As(err, &target))
}
