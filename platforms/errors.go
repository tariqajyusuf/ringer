package platforms

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
