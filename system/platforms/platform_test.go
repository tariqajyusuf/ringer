package platforms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tariqajyusuf/ringer/system"
)

type mockPlatform struct {
	addErr      error
	removeErr   error
	stub        string
	installed   bool
	enabledFor  system.Kernel
}

func (m *mockPlatform) AddPackage(_ string) error                 { return m.addErr }
func (m *mockPlatform) RemovePackage(_ string) error              { return m.removeErr }
func (m *mockPlatform) PlatformInfo() string                      { return "mock" }
func (m *mockPlatform) PlatformStub() string                      { return m.stub }
func (m *mockPlatform) EnabledForSystem(s system.SystemInfo) bool { return s.Kernel == m.enabledFor }
func (m *mockPlatform) IsInstalled() bool                         { return m.installed }
func (m *mockPlatform) SetupPackageManager(_ bool) error          { return nil }

func newTestBroker(stub string, mock *mockPlatform) *Broker {
	return &Broker{
		Platforms:          map[string]Platform{stub: mock},
		preferred_platform: stub,
	}
}

func TestBroker_PreferredPlatform(t *testing.T) {
	b := newTestBroker("mock", &mockPlatform{stub: "mock"})
	assert.Equal(t, "mock", b.PreferredPlatform())
}

func TestBroker_SetPreferredPlatform_Valid(t *testing.T) {
	b := newTestBroker("mock", &mockPlatform{stub: "mock"})
	err := b.SetPreferredPlatform("mock")
	require.NoError(t, err)
	assert.Equal(t, "mock", b.PreferredPlatform())
}

func TestBroker_SetPreferredPlatform_Invalid(t *testing.T) {
	b := newTestBroker("mock", &mockPlatform{stub: "mock"})
	err := b.SetPreferredPlatform("nonexistent")
	assert.Error(t, err)
}

func TestBroker_AddPackage_PropagatesError(t *testing.T) {
	installErr := &InstallError{message: "failed"}
	b := newTestBroker("mock", &mockPlatform{addErr: installErr})
	err := b.AddPackage("somepkg")
	assert.ErrorIs(t, err, installErr)
}

func TestBroker_AddPackage_Success(t *testing.T) {
	b := newTestBroker("mock", &mockPlatform{})
	err := b.AddPackage("somepkg")
	assert.NoError(t, err)
}

func TestBroker_RemovePackage_PropagatesError(t *testing.T) {
	removeErr := &PackageNotFound{message: "not found"}
	b := newTestBroker("mock", &mockPlatform{removeErr: removeErr})
	err := b.RemovePackage("somepkg")
	assert.ErrorIs(t, err, removeErr)
}

func TestBroker_RemovePackage_Success(t *testing.T) {
	b := newTestBroker("mock", &mockPlatform{})
	err := b.RemovePackage("somepkg")
	assert.NoError(t, err)
}

func TestBroker_SkippedPlatforms(t *testing.T) {
	b := &Broker{
		Platforms:        make(map[string]Platform),
		skippedPlatforms: []string{"homebrew"},
	}
	assert.Equal(t, []string{"homebrew"}, b.SkippedPlatforms())
}

func TestDefaultPlatformForSystem_MacOS(t *testing.T) {
	assert.Equal(t, "homebrew", DefaultPlatformForSystem(system.SystemInfo{Kernel: system.MacOS}))
}

func TestDefaultPlatformForSystem_Linux(t *testing.T) {
	assert.Equal(t, "homebrew", DefaultPlatformForSystem(system.SystemInfo{Kernel: system.Linux}))
}

func TestDefaultPlatformForSystem_Windows(t *testing.T) {
	assert.Equal(t, "winget", DefaultPlatformForSystem(system.SystemInfo{Kernel: system.Windows}))
}

func TestDefaultPlatformForSystem_Unknown(t *testing.T) {
	assert.Equal(t, "", DefaultPlatformForSystem(system.SystemInfo{Kernel: system.Unknown}))
}
