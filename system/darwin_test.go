package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMacOSProfile_Valid(t *testing.T) {
	input := `{"SPSoftwareDataType":[{"kernel_version":"Darwin 23.5.0","os_version":"macOS 14.5 (23F79)"}]}`
	info := parseMacOSProfile([]byte(input))
	assert.Equal(t, MacOS, info.Kernel)
	assert.Equal(t, "macOS 14.5 (23F79)", info.Distro)
	require.NotNil(t, info.Version)
	assert.Equal(t, "14.5.0", info.Version.String())
}

func TestParseMacOSProfile_EmptyDataType(t *testing.T) {
	input := `{"SPSoftwareDataType":[]}`
	info := parseMacOSProfile([]byte(input))
	assert.Equal(t, MacOS, info.Kernel)
	assert.Empty(t, info.Distro)
	assert.Nil(t, info.Version)
}

func TestParseMacOSProfile_InvalidJSON(t *testing.T) {
	info := parseMacOSProfile([]byte("not json"))
	assert.Equal(t, MacOS, info.Kernel)
	assert.Nil(t, info.Version)
}

func TestParseMacOSProfile_BadVersion(t *testing.T) {
	input := `{"SPSoftwareDataType":[{"os_version":"macOS notaversion (build)"}]}`
	require.NotPanics(t, func() {
		info := parseMacOSProfile([]byte(input))
		assert.Equal(t, MacOS, info.Kernel)
	})
}

func TestParseMacOSProfile_NoSpaceInVersion(t *testing.T) {
	input := `{"SPSoftwareDataType":[{"os_version":"macOS"}]}`
	require.NotPanics(t, func() {
		info := parseMacOSProfile([]byte(input))
		assert.Equal(t, MacOS, info.Kernel)
	})
}
