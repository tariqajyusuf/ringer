package platforms

import "github.com/hashicorp/go-version"

/*
Indicates that the package provided was not found. This usually indicates
that there is a misconfiguration in the config manifest that points to the
wrong package.
*/
type PackageNotFound error

/*
Indicates the package was not installed successfully.
*/
type InstallError error

/*
If the package installation requires a higher privilege on the system, this will
fire.
*/
type AuthorizationError error

/*
Provides information about system context that indicates which platforms apply
to the current runtime environment.
*/
type SystemInfo struct {
	kernel  string
	distro  string
	version version.Version
}
