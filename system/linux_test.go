package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseReleaseFile_Standard(t *testing.T) {
	content := "ID=ubuntu\nID_LIKE=debian\nVERSION_ID=22.04\n"
	info := parseReleaseFile(content)
	assert.Equal(t, "ubuntu", info["ID"])
	assert.Equal(t, "debian", info["ID_LIKE"])
	assert.Equal(t, "22.04", info["VERSION_ID"])
}

func TestParseReleaseFile_QuotedValues(t *testing.T) {
	content := `ID="ubuntu"` + "\n" + `VERSION_ID="22.04"` + "\n"
	info := parseReleaseFile(content)
	assert.Equal(t, "ubuntu", info["ID"])
	assert.Equal(t, "22.04", info["VERSION_ID"])
}

func TestParseReleaseFile_BlankAndCommentLines(t *testing.T) {
	content := "# Comment line\n\nID=fedora\n\nVERSION_ID=39\n"
	require.NotPanics(t, func() {
		info := parseReleaseFile(content)
		assert.Equal(t, "fedora", info["ID"])
		assert.Equal(t, "39", info["VERSION_ID"])
	})
}

func TestParseReleaseFile_ValueWithEquals(t *testing.T) {
	content := "KEY=value=with=equals\n"
	info := parseReleaseFile(content)
	assert.Equal(t, "value=with=equals", info["KEY"])
}
