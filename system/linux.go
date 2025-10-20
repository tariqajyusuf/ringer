package system

import (
	"os"
	"strings"

	"github.com/hashicorp/go-version"
)

func GetLinuxInfo() SystemInfo {
	release_info, err := os.ReadFile("/etc/*-release")
	// TODO: Better error handling
	if err != nil {
		return SystemInfo{Kernel: Linux}
	}
	system_info := make(map[string]string)
	for _, line := range strings.Split(string(release_info), "\n") {
		values := strings.Split(line, "=")
		system_info[values[0]] = values[1]
	}
	result := SystemInfo{Kernel: Linux}
	if val, ok := system_info["ID_LIKE"]; ok {
		result.Distro = val
	} else {
		result.Distro = system_info["ID"]
	}
	result.Version = version.Must(version.NewVersion(system_info["VERSION_ID"]))
	return result
}
