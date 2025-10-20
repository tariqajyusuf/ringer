package system

import (
	"os/exec"
	"strings"

	"github.com/hashicorp/go-version"
)

/*
Provides information about system context that indicates which platforms apply
to the current runtime environment.
*/
//lint:ignore U1000 Ignore unused for now
type SystemInfo struct {
	Kernel  Kernel
	Distro  string
	Version *version.Version
}

type Kernel int

const (
	Unknown Kernel = iota
	Windows
	Linux
	MacOS
)

func GetSystemInfo() SystemInfo {
	// uname works for all Unix-like systems
	cmd := exec.Command("uname", "-s")
	output, err := cmd.Output()
	if err != nil {
		// This is probably a Windows system
		return GetWindowsInfo()
	}
	switch strings.TrimSpace(string(output)) {
	case "Darwin":
		return GetMacOSInfo()
	case "Linux":
		return GetLinuxInfo()
	default:
		return SystemInfo{Kernel: Unknown}
	}
}
