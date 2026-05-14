package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestGuise_Unmarshal_ValidList(t *testing.T) {
	input := "packages:\n  - git\n  - vscode\n"
	var g Guise
	require.NoError(t, yaml.Unmarshal([]byte(input), &g))
	assert.Equal(t, []string{"git", "vscode"}, g.Packages)
}

func TestGuise_Unmarshal_Empty(t *testing.T) {
	input := "packages: []\n"
	var g Guise
	require.NoError(t, yaml.Unmarshal([]byte(input), &g))
	assert.Empty(t, g.Packages)
}

func TestGuise_Unmarshal_MissingKey(t *testing.T) {
	input := "{}"
	var g Guise
	require.NoError(t, yaml.Unmarshal([]byte(input), &g))
	assert.Nil(t, g.Packages)
}
