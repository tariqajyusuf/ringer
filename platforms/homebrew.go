package platforms

import (
	"io"
	"os/exec"
	"strings"
)

/*
Homebrew package platform (https://brew.sh/).
*/
type Homebrew struct {
	Platform
}

// TODO: Deal with sudo authorization if needed.
func (h Homebrew) AddPackage(name string) error {
	return h.runBrew("install", name)
}

func (h Homebrew) RemovePackage(name string) error {
	return h.runBrew("remove", name)
}

func (h Homebrew) PlatformInfo() string {
	runner := exec.Command("brew", "-v")
	output, err := runner.Output()
	if err != nil {
		return "Unknown"
	}
	return string(output)
}

func (h Homebrew) PlatformStub() string {
	return "homebrew"
}

func (h Homebrew) EnabledForSystem(system SystemInfo) bool {
	switch system.kernel {
	case MacOS, Linux:
		return true
	default:
		return false
	}
}

func (h Homebrew) SetupPackageManager() error {
	if h.PlatformInfo() != "Unknown" {
		return nil
	}

	// Install Homebrew based on brew.sh instructions.
	runner := exec.Command(
		"/bin/bash",
		"-c",
		"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	// TODO: Log better when there is an issue.
	return runner.Run()
}

func (h Homebrew) runBrew(verb string, packageName string) error {
	runner := exec.Command("brew", verb, packageName)
	runner.Env = append(runner.Env, "NONINTERACTIVE=1")
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
		if strings.Contains(last_line, "No formulae or casks found") {
			return PackageNotFound{message: last_line}
		}
	}
	return nil
}
