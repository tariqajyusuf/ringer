package system

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-version"
)

type SystemProfile struct {
	SPSoftwareDataType []struct {
		KernelVersion string `json:"kernel_version"`
		OSVersion     string `json:"os_version"`
	} `json:"SPSoftwareDataType"`
}

func GetMacOSInfo() SystemInfo {
	runner := exec.Command(
		"system_profiler",
		"SPSoftwareDataType",
		"-json",
		"-detailLevel",
		"mini")
	output, err := runner.Output()
	// TODO: Better error handling
	if err != nil {
		return SystemInfo{Kernel: MacOS}
	}
	return parseMacOSProfile(output)
}

func parseMacOSProfile(output []byte) SystemInfo {
	system_profile := SystemProfile{}
	if err := json.Unmarshal(output, &system_profile); err != nil {
		return SystemInfo{Kernel: MacOS}
	}
	if len(system_profile.SPSoftwareDataType) == 0 {
		return SystemInfo{Kernel: MacOS}
	}
	os_version := system_profile.SPSoftwareDataType[0].OSVersion
	split_str := strings.Split(os_version, " ")
	if len(split_str) < 2 {
		return SystemInfo{Kernel: MacOS, Distro: os_version}
	}
	v, err := version.NewVersion(split_str[1])
	if err != nil {
		return SystemInfo{Kernel: MacOS, Distro: os_version}
	}
	return SystemInfo{
		Kernel:  MacOS,
		Distro:  os_version,
		Version: v,
	}
}
