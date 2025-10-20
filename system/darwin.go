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
	system_profile := SystemProfile{}
	err = json.Unmarshal(output, &system_profile)
	// TODO: Better error handling
	if err != nil {
		return SystemInfo{Kernel: MacOS}
	}
	split_str := strings.Split(system_profile.SPSoftwareDataType[0].OSVersion, " ")
	major_version := split_str[1]
	minor_version := strings.Trim(split_str[1], "()")
	return SystemInfo{
		Kernel:  MacOS,
		Distro:  system_profile.SPSoftwareDataType[0].OSVersion,
		Version: version.Must(version.NewVersion(major_version + "+" + minor_version)),
	}
}
