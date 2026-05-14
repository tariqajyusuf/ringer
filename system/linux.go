package system

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-version"
)

func GetLinuxInfo() SystemInfo {
	matches, _ := filepath.Glob("/etc/*-release")
	if len(matches) == 0 {
		return SystemInfo{Kernel: Linux}
	}
	target := matches[0]
	for _, m := range matches {
		if m == "/etc/os-release" {
			target = m
			break
		}
	}
	releaseInfo, err := os.ReadFile(target)
	if err != nil {
		return SystemInfo{Kernel: Linux}
	}
	system_info := parseReleaseFile(string(releaseInfo))
	result := SystemInfo{Kernel: Linux}
	if val, ok := system_info["ID_LIKE"]; ok {
		result.Distro = val
	} else {
		result.Distro = system_info["ID"]
	}
	if v, err := version.NewVersion(system_info["VERSION_ID"]); err == nil {
		result.Version = v
	}
	return result
}

func parseReleaseFile(content string) map[string]string {
	info := make(map[string]string)
	for _, line := range strings.Split(content, "\n") {
		values := strings.SplitN(line, "=", 2)
		if len(values) < 2 {
			continue
		}
		info[values[0]] = strings.Trim(values[1], "\"")
	}
	return info
}
