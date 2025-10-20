package io

/*
Circle files are declarations of a desired system state. A list of packages is
declared based on their Guise names.
*/
//lint:ignore U1000 Ignore unused for now
type Circle struct {
	Packages []string `yaml:"packages,flow"`
}
