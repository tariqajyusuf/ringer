package platforms

/*
Platform represents a operating system environment.

When Ringer is installed, it needs to undersatnd what environment it's
running in. THe platform struct provides a common interface for how Ringer
should install packages. For example, to install Visual Studio code on Windows,
we could use `winget install Microsoft.VisualStudioCode` but on Mac we
can use `brew install visual-studio-code`.
*/
type Platform interface {
	/*
		Installs the package based on the platform-specific package name. For
		example, `Windows.InstallPackage("Microsoft.VisualStudioCode")`. The error
		can be either one of:
		  - PackageNotFound
			- InstallError
			- AuthorizationError
	*/
	AddPackage(name string) error

	/*
		Removes the package based on the platform-specific package name. The error
		can be either:
			- PackageNotFound
			- AuthorizationError
	*/
	RemovePackage(name string) error

	/*
		Gets information about the package based on the platform-specific package
		name. Package search is handled at the application level. The error will
		usually be PackageNotFound.
	*/
	PackageInfo(name string) error

	/*
		Gets platform information.
	*/
	GetPlatformInfo() string
}
