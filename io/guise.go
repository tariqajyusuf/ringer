package io

import (
	"os"

	"gopkg.in/yaml.v3"
)

// TODO: Add validity checks for guises.

/*
Guises are definitions of applications that can be installed on different
platforms. Each guise contains the name, description, and platform-specific
package names needed to install the application on different operating systems.
*/
//lint:ignore U1000 Ignore unused for now
type Guise struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Platforms   map[string]struct {
		PackageName string `yaml:"name"`
	} `yaml:"platforms,flow"`
}

func LocateGuise(name string) (*Guise, error) {
	// TODO: This needs to be more dynamic if the app will be more self-contained.
	bytes, err := os.ReadFile("data/guises/" + name + ".guise.yaml")
	if err != nil {
		return nil, err
	}
	var guise Guise
	err = yaml.Unmarshal(bytes, &guise)
	return &guise, err
}
