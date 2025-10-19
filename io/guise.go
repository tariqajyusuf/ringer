package io

/*
Guises are definitions of applications that can be installed on different
platforms. Each guise contains the name, description, and platform-specific
package names needed to install the application on different operating systems.
*/
type Guise struct {
	Marshalable
	name        string `yaml:"name"`
	description string `yaml:"description"`
	platforms   []struct {
		platform    string `yaml:"platform"`
		packageName string `yaml:"packageName"`
	} `yaml:"platforms,flow"`
}
