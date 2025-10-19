package io

/*
Circle files are declarations of a desired system state. A list of packages is
declared based on their Guise names.
*/
type Circle struct {
	Marshalable
	packages []string `yaml:"packages,flow"`
}
