package platforms

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tariqajyusuf/ringer/system"
)

/*
Homebrew package platform (https://brew.sh/).
*/
type Homebrew struct {
	Platform
	verbose bool
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

func (h Homebrew) IsInstalled() bool {
	return h.PlatformInfo() != "Unknown"
}

func (h Homebrew) EnabledForSystem(sysinfo system.SystemInfo) bool {
	switch sysinfo.Kernel {
	case system.MacOS, system.Linux:
		return true
	default:
		return false
	}
}

func (h Homebrew) SetupPackageManager(verbose bool) error {
	if h.PlatformInfo() != "Unknown" {
		return nil
	}

	// Install Homebrew based on brew.sh instructions.
	// curl | bash is equivalent to: bash -c "$(curl ...)" — the latter only works
	// when a parent shell expands $() before passing it to bash -c, which exec.Command
	// does not do.
	runner := exec.Command(
		"/bin/bash",
		"-c",
		"curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | /bin/bash")
	runner.Env = os.Environ()
	runner.Env = append(runner.Env, "NONINTERACTIVE=1")
	if verbose {
		runner.Stdout = os.Stdout
		runner.Stderr = os.Stderr
	}
	if err := runner.Run(); err != nil {
		return err
	}

	if system.GetSystemInfo().Kernel == system.Linux {
		return h.linuxPostInstall(verbose)
	}
	return nil
}

const linuxBrewBin = "/home/linuxbrew/.linuxbrew/bin/brew"

func (h Homebrew) linuxPostInstall(verbose bool) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Add brew to PATH in ~/.bashrc (idempotent).
	shellenvLine := fmt.Sprintf(`eval "$(%s shellenv bash)"`, linuxBrewBin)
	if err := appendLineIfAbsent(filepath.Join(home, ".bashrc"), shellenvLine); err != nil {
		return err
	}

	// Install build-essential — requires sudo, best-effort.
	h.runSilently(verbose, "sudo", "apt-get", "install", "-y", "build-essential")

	// Install GCC — use full path since PATH isn't updated in the current process yet.
	h.runSilently(verbose, linuxBrewBin, "install", "gcc")

	return nil
}

// appendLineIfAbsent appends line to path only if it isn't already present.
func appendLineIfAbsent(path, line string) error {
	content, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if strings.Contains(string(content), line) {
		return nil
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close() //nolint:errcheck
	_, err = fmt.Fprintf(f, "\n%s\n", line)
	return err
}

// runSilently runs a command, streaming output only when verbose. Errors are ignored
// because these post-install steps are best-effort.
func (h Homebrew) runSilently(verbose bool, name string, args ...string) {
	cmd := exec.Command(name, args...)
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	_ = cmd.Run()
}

func (h Homebrew) runBrew(verb string, packageName string) error {
	runner := exec.CommandContext(context.Background(), "brew", verb, packageName)
	runner.Env = os.Environ()
	runner.Env = append(runner.Env, "NONINTERACTIVE=1")
	if h.verbose {
		var buf bytes.Buffer
		runner.Stdout = io.MultiWriter(os.Stdout, &buf)
		runner.Stderr = io.MultiWriter(os.Stderr, &buf)
		if err := runner.Run(); err != nil {
			return h.parseOutput(buf.Bytes())
		}
		return nil
	}
	if out, err := runner.CombinedOutput(); err != nil {
		return h.parseOutput(out)
	}
	return nil
}

func (h Homebrew) parseOutput(bytes []byte) error {
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if strings.Contains(line, "No formulae or casks") {
			return &PackageNotFound{message: line}
		}
	}
	return &InstallError{message: string(bytes)}
}
