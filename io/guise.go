package io

/*
Guises are definitions of applications that can be installed on different
platforms. Each guise contains the name, description, and platform-specific
package names needed to install the application on different operating systems.
*/
//lint:ignore U1000 Ignore unused for now
type Guise struct {
	Marshalable
	name        string `yaml:"name"`
	description string `yaml:"description"`
	platforms   []struct {
		platform    string `yaml:"platform"`
		packageName string `yaml:"name"`
	} `yaml:"platforms,flow"`
}
