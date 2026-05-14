package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveCmd_NoArgs(t *testing.T) {
	out := captureStdout(t, func() {
		rootCmd.SetArgs([]string{"remove"})
		rootCmd.Execute() //nolint:errcheck
	})
	assert.Contains(t, out, "Please provide a package name to remove")
}
