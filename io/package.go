package io

import (
	"os"

	"gopkg.in/yaml.v3"
)

// TODO: Add validity checks for packages.

/*
Packages are definitions of applications that can be installed on different
platforms. Each package contains the name, description, and platform-specific
package names needed to install the application on different operating systems.
*/
//lint:ignore U1000 Ignore unused for now
type Package struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Platforms   map[string]struct {
		PackageName string `yaml:"name"`
	} `yaml:"platforms,flow"`
}

func LocatePackage(name string) (*Package, error) {
	// TODO: This needs to be more dynamic if the app will be more self-contained.
	bytes, err := os.ReadFile("data/packages/" + name + ".package.yaml")
	if err != nil {
		return nil, err
	}
	var pkg Package
	err = yaml.Unmarshal(bytes, &pkg)
	return &pkg, err
}
