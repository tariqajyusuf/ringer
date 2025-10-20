package platforms

import (
	"errors"
	"io"
	"os/exec"
	"strings"

	"github.com/tariqajyusuf/ringer/system"
)

/*
Winget package platform (https://learn.microsoft.com/en-us/windows/package-manager/winget/).
*/
type Winget struct {
	Platform
}

func (w Winget) AddPackage(name string) error {
	return w.runWinget("install", name)
}

func (w Winget) RemovePackage(name string) error {
	return w.runWinget("uninstall", name)
}

func (w Winget) PlatformInfo() string {
	runner := exec.Command("winget", "--version")
	output, err := runner.Output()
	if err != nil {
		return "Unknown"
	}
	return string(output)
}

func (w Winget) PlatformStub() string {
	return "winget"
}

func (w Winget) EnabledForSystem(sysinfo system.SystemInfo) bool {
	switch sysinfo.Kernel {
	case system.Windows:
		return true
	default:
		return false
	}
}

func (w Winget) SetupPackageManager() error {
	if w.PlatformInfo() != "Unknown" {
		return nil
	}

	// Winget is included by default in Windows 10 (from version 1809) and Windows 11.
	return errors.New("winget is not installed; please install it from the Microsoft Store (https://apps.microsoft.com/detail/9nblggh4nns1)")
}

func (w Winget) runWinget(verb string, packageName string) error {
	runner := exec.Command("winget", verb, packageName, "--accept-source-agreements", "--accept-package-agreements")
	err_pipe, err := runner.StderrPipe()
	// TODO: Better handling for output failure
	if err != nil {
		return err
	}
	if err := runner.Start(); err != nil {

		stderr, err := io.ReadAll(err_pipe)
		// TODO: Warn of system state if output cannot be read
		if err != nil {
			return err
		}

		lines := strings.Split(string(stderr), "\n")
		last_line := lines[len(lines)-1]
		if strings.Contains(last_line, "No package found matching input criteria") {
			return PackageNotFound{message: last_line}
		}
	}
	return nil
}
