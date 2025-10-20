package platforms

import (
	"fmt"

	"github.com/tariqajyusuf/ringer/system"
)

// TODO: Add support for standard linux commands and package managers (apt, yum, pacman, etc).

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
	EnabledForSystem(system system.SystemInfo) bool

	/*
		Installs any prerequisite programs, if needed, and ensures we're ready to
		go. This should be expected to run every time the program is run. If nothing
		is needed, then should simply return. The error can be AuthorizationError if
		the application is not permitted to run under the authorized user context.
	*/
	SetupPackageManager() error
}

/*
The Broker is how you interact with all platforms built into Ringer. Any new
platform created in future versions will be registered here. The broker will
only allow calls to platforms that are enabled for the current system and will
handle any necessary setup prior to package installation/removal.
*/
type Broker struct {
	Platforms          map[string]Platform
	preferred_platform string
}

/*
Creates and registers a broker with all available platforms. This should only be
run once.
*/
func NewBroker() *Broker {
	// TODO: We need to include some sort of state management to remember where we
	// got different packages.
	b := &Broker{
		Platforms:          make(map[string]Platform),
		preferred_platform: "",
	}
	possible_platforms := map[string]Platform{}
	possible_platforms["homebrew"] = &Homebrew{}
	possible_platforms["winget"] = &Winget{}

	for key, platform := range possible_platforms {
		if platform.EnabledForSystem(system.GetSystemInfo()) {
			if err := platform.SetupPackageManager(); err != nil {
				// TODO: Log the error
				println("error")
			}
			if len(b.preferred_platform) == 0 {
				b.preferred_platform = key
			}
			b.Platforms[key] = platform
		}
	}
	return b
}

func (b *Broker) PreferredPlatform() string {
	return b.preferred_platform
}

func (b *Broker) SetPreferredPlatform(name string) error {
	if _, ok := b.Platforms[name]; !ok {
		return fmt.Errorf("platform %s is not available", name)
	}
	b.preferred_platform = name
	return nil
}

// TODO: Try every package manager until one works.
func (b *Broker) AddPackage(name string) {
	fmt.Printf("Installing via %s\n", b.preferred_platform)
	err := b.Platforms[b.preferred_platform].AddPackage(name)
	// TODO: We will probably need to handle errors more intelligently down the
	// line but for now we will just regurgitate from the command line.
	if err != nil {
		fmt.Printf("Error installing package via %s error: %s\n", b.preferred_platform, err)
	} else {
		fmt.Printf("Successfully installed package via %s\n", b.preferred_platform)
		return
	}
}

// TODO: Try every package manager until one works.
func (b *Broker) RemovePackage(name string) {
	fmt.Printf("Removing via %s\n", b.preferred_platform)
	err := b.Platforms[b.preferred_platform].RemovePackage(name)
	// TODO: We will probably need to handle errors more intelligently down the
	// line but for now we will just regurgitate from the command line.
	if err != nil {
		fmt.Printf("Error removing package via %s error: %s\n", b.preferred_platform, err)
	} else {
		fmt.Printf("Successfully removed package via %s\n", b.preferred_platform)
		return
	}
}
