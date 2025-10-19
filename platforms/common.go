package platforms

import "github.com/hashicorp/go-version"

type RingerError struct {
	error
	message string
}

func (e RingerError) Error() string {
	return e.message
}

/*
Indicates that the package provided was not found. This usually indicates
that there is a misconfiguration in the config manifest that points to the
wrong package.
*/
type PackageNotFound RingerError

/*
Indicates the package was not installed successfully.
*/
type InstallError RingerError

/*
If the package installation requires a higher privilege on the system, this will
fire.
*/
type AuthorizationError RingerError

/*
Provides information about system context that indicates which platforms apply
to the current runtime environment.
*/
//lint:ignore U1000 Ignore unused for now
type SystemInfo struct {
	kernel  Kernel
	distro  string
	version version.Version
}

type Kernel int

const (
	Unknown Kernel = iota
	Windows
	Linux
	MacOS
)
