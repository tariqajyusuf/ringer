package io

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/tariqajyusuf/ringer/system"
)

// TODO: This needs to be more dynamic if the app will be more self-contained.
var PackageDataDir = "data/packages"

// kernelToOSString maps system.Kernel values to the YAML os strings used in package files.
var kernelToOSString = map[system.Kernel]string{
	system.MacOS:   "macos",
	system.Linux:   "linux",
	system.Windows: "windows",
}

// PlatformEntry holds the platform-specific data for one package manager entry.
type PlatformEntry struct {
	PackageName string `yaml:"name"`
}

// Package is a definition of an application that can be installed on different
// platforms. Each package contains the name, description, and platform-specific
// package names needed to install the application on different operating systems.
type Package struct {
	Name        string                   `yaml:"name"`
	Description string                   `yaml:"description"`
	OS          []string                 `yaml:"os,omitempty"`
	Platforms   map[string]PlatformEntry `yaml:"platforms,flow"`
}

// CheckOSAllowed returns an error if the current OS is not in the package's OS
// restriction list. If the list is empty, all OSes are allowed.
func (p *Package) CheckOSAllowed(kernel system.Kernel) error {
	if len(p.OS) == 0 {
		return nil
	}
	current, known := kernelToOSString[kernel]
	if !known {
		return fmt.Errorf("package %q restricts OS but current OS is unknown; allowed: %v", p.Name, p.OS)
	}
	for _, allowed := range p.OS {
		if allowed == current {
			return nil
		}
	}
	return fmt.Errorf("package %q is not supported on %s (supported OS: %v)", p.Name, current, p.OS)
}

func LocatePackage(name string) (*Package, error) {
	bytes, err := os.ReadFile(PackageDataDir + "/" + name + ".package.yaml")
	if err != nil {
		return nil, err
	}
	var pkg Package
	err = yaml.Unmarshal(bytes, &pkg)
	return &pkg, err
}
