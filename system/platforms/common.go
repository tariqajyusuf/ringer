package platforms

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
