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
		  - PackageNotFound, if the underlying manager cannot find the package.
			- InstallError, if the underlying manager cannot install the package.
			- AuthorizationError, if the package manager is unable to install under
			  the user context.
	*/
	AddPackage(name string) error

	/*
		Removes the package based on the platform-specific package name. The error
		can be either:
			- PackageNotFound, if the underlying manager cannot find the package.
			- AuthorizationError, if the package manager is unable to install under
			  the user context.
	*/
	RemovePackage(name string) error

	/*
		Gets platform information.
	*/
	PlatformInfo() string

	/*
		For guise files, the platform stub is what we use to identify the translated
		package name.
	*/
	PlatformStub() string

	/*
		For platforms to nominate themselves for an option based on the base system
		information. This allows for multiple platforms to nominate themselves
		(e.g. Debian platforms can install along with generic Linux installers).
	*/
	EnabledForSystem(system SystemInfo) bool

	/*
		Installs any prerequisite programs, if needed, and ensures we're ready to
		go. This should be expected to run every time the program is run. If nothing
		is needed, then should simply return. The error can be AuthorizationError if
		the application is not permitted to run under the authorized user context.
	*/
	SetupPackageManager() error
}
