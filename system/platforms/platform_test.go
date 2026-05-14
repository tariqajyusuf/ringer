package platforms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tariqajyusuf/ringer/system"
)

type mockPlatform struct {
	addErr    error
	removeErr error
	stub      string
}

func (m *mockPlatform) AddPackage(_ string) error                 { return m.addErr }
func (m *mockPlatform) RemovePackage(_ string) error              { return m.removeErr }
func (m *mockPlatform) PlatformInfo() string                      { return "mock" }
func (m *mockPlatform) PlatformStub() string                      { return m.stub }
func (m *mockPlatform) EnabledForSystem(_ system.SystemInfo) bool { return true }
func (m *mockPlatform) SetupPackageManager() error                { return nil }

func newTestBroker(stub string, mock *mockPlatform) *Broker {
	b := &Broker{
		Platforms:          map[string]Platform{stub: mock},
		preferred_platform: stub,
	}
	return b
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
