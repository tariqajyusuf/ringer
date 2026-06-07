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
		For package files, the platform stub is what we use to identify the translated
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
		Returns true if the platform manager is already installed on this system
		and ready to use. Unlike SetupPackageManager, this does not attempt to
		install anything.
	*/
	IsInstalled() bool

	/*
		Installs any prerequisite programs needed to use this platform manager.
		Should only be called explicitly by the user via `ringer platform install`.
		When verbose is true, install output is streamed to stdout/stderr.
		The error can be AuthorizationError if elevated privileges are required.
	*/
	SetupPackageManager(verbose bool) error
}

/*
The Broker is how you interact with all platforms built into Ringer. Any new
platform created in future versions will be registered here. The broker will
only allow calls to platforms that are enabled for the current system and are
already installed.
*/
type Broker struct {
	Platforms          map[string]Platform
	preferred_platform string
	verbose            bool
	skippedPlatforms   []string
}

// DefaultPlatformForSystem returns the canonical platform manager for the given OS.
func DefaultPlatformForSystem(info system.SystemInfo) string {
	switch info.Kernel {
	case system.MacOS, system.Linux:
		return "homebrew"
	case system.Windows:
		return "winget"
	default:
		return ""
	}
}

/*
Creates and registers a broker with all available, already-installed platforms.
preferredPlatform is the user's saved preference (empty string to use the OS default).
*/
func NewBroker(verbose bool, preferredPlatform string) *Broker {
	// TODO: We need to include some sort of state management to remember where we
	// got different packages.
	b := &Broker{
		Platforms:          make(map[string]Platform),
		preferred_platform: "",
		verbose:            verbose,
	}
	possible_platforms := map[string]Platform{}
	possible_platforms["homebrew"] = &Homebrew{verbose: verbose}
	possible_platforms["winget"] = &Winget{verbose: verbose}

	sysinfo := system.GetSystemInfo()
	for key, platform := range possible_platforms {
		if platform.EnabledForSystem(sysinfo) {
			if !platform.IsInstalled() {
				b.skippedPlatforms = append(b.skippedPlatforms, key)
				continue
			}
			b.Platforms[key] = platform
		}
	}

	// Resolve preferred platform: user config → OS default → first available.
	switch {
	case preferredPlatform != "" && b.Platforms[preferredPlatform] != nil:
		b.preferred_platform = preferredPlatform
	case b.Platforms[DefaultPlatformForSystem(sysinfo)] != nil:
		b.preferred_platform = DefaultPlatformForSystem(sysinfo)
	default:
		for key := range b.Platforms {
			b.preferred_platform = key
			break
		}
	}

	return b
}

// SkippedPlatforms returns platform names that are supported on this OS but not installed.
func (b *Broker) SkippedPlatforms() []string {
	return b.skippedPlatforms
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
func (b *Broker) AddPackage(name string) error {
	fmt.Printf("Installing via %s\n", b.preferred_platform)
	err := b.Platforms[b.preferred_platform].AddPackage(name)
	if err != nil {
		fmt.Printf("Error installing package via %s: %s\n", b.preferred_platform, err)
	}
	return err
}

// TODO: Try every package manager until one works.
func (b *Broker) RemovePackage(name string) error {
	if b.verbose {
		fmt.Printf("Removing via %s\n", b.preferred_platform)
	}
	err := b.Platforms[b.preferred_platform].RemovePackage(name)
	if err != nil {
		fmt.Printf("Error removing package via %s: %s\n", b.preferred_platform, err)
	} else if b.verbose {
		fmt.Printf("Successfully removed package via %s\n", b.preferred_platform)
	}
	return err
}
