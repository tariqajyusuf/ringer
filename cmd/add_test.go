package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCmd_NoArgs(t *testing.T) {
	out := captureStdout(t, func() {
		rootCmd.SetArgs([]string{"add"})
		rootCmd.Execute() //nolint:errcheck
	})
	assert.Contains(t, out, "Please provide a package name to add")
}
