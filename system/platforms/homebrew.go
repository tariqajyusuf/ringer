package platforms

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/tariqajyusuf/ringer/system"
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

func (h Homebrew) EnabledForSystem(sysinfo system.SystemInfo) bool {
	switch sysinfo.Kernel {
	case system.MacOS, system.Linux:
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
	runner.Env = os.Environ()
	runner.Env = append(runner.Env, "NONINTERACTIVE=1")
	// TODO: Log better when there is an issue.
	return runner.Run()
}

func (h Homebrew) runBrew(verb string, packageName string) error {
	runner := exec.CommandContext(context.Background(), "brew", verb, packageName)
	runner.Env = os.Environ()
	runner.Env = append(runner.Env, "NONINTERACTIVE=1")
	if bytes, err := runner.CombinedOutput(); err != nil {
		return h.parseOutput(bytes)
	}
	return nil
}

func (h Homebrew) parseOutput(bytes []byte) error {
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Error") {
			if strings.Contains(line, "No formulae or casks") {
				return errors.New(line)
			}
		}
	}
	return nil
}
