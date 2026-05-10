package io

/*
Guise files are declarations of a desired system state. A list of packages is
declared based on their package names.
*/
//lint:ignore U1000 Ignore unused for now
type Guise struct {
	Packages []string `yaml:"packages,flow"`
}
