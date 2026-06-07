package platforms

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tariqajyusuf/ringer/system"
)

func TestHomebrew_PlatformStub(t *testing.T) {
	assert.Equal(t, "homebrew", Homebrew{}.PlatformStub())
}

func TestHomebrew_EnabledForSystem_MacOS(t *testing.T) {
	assert.True(t, Homebrew{}.EnabledForSystem(system.SystemInfo{Kernel: system.MacOS}))
}

func TestHomebrew_EnabledForSystem_Linux(t *testing.T) {
	assert.True(t, Homebrew{}.EnabledForSystem(system.SystemInfo{Kernel: system.Linux}))
}

func TestHomebrew_EnabledForSystem_Windows(t *testing.T) {
	assert.False(t, Homebrew{}.EnabledForSystem(system.SystemInfo{Kernel: system.Windows}))
}

func TestHomebrew_EnabledForSystem_Unknown(t *testing.T) {
	assert.False(t, Homebrew{}.EnabledForSystem(system.SystemInfo{Kernel: system.Unknown}))
}

func TestHomebrew_IsInstalled_WhenNotPresent(t *testing.T) {
	// brew is not installed in this environment
	assert.False(t, Homebrew{}.IsInstalled())
}

func TestHomebrew_ParseOutput_NoFormulaeError(t *testing.T) {
	h := Homebrew{}
	output := []byte("Error: No formulae or casks with that name.\n")
	err := h.parseOutput(output)
	require.Error(t, err)
	var target *PackageNotFound
	assert.True(t, errors.As(err, &target))
}

func TestHomebrew_ParseOutput_GenericError(t *testing.T) {
	h := Homebrew{}
	output := []byte("Error: Something else went wrong.\n")
	err := h.parseOutput(output)
	require.Error(t, err)
	var target *InstallError
	assert.True(t, errors.As(err, &target))
}
