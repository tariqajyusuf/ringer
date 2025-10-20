package system

import (
	"encoding/json"
	"os/exec"

	"github.com/hashicorp/go-version"
)

type ComputerInfo struct {
	OsName    string `json:"OsName"`
	OsVersion string `json:"OsVersion"`
}

func GetWindowsInfo() SystemInfo {
	output, err := exec.Command("powershell", "-NoProfile", "-Command", "Get-ComputerInfo -Property @('OsName', 'OsVersion') | ConvertTo-Json").Output()
	// TODO: Better error handling
	if err != nil {
		return SystemInfo{Kernel: Windows}
	}
	computer_info := &ComputerInfo{}
	err = json.Unmarshal(output, &computer_info)
	// TODO: Better error handling
	if err != nil {
		return SystemInfo{Kernel: Windows}
	}
	return SystemInfo{
		Kernel:  Windows,
		Distro:  computer_info.OsName,
		Version: version.Must(version.NewVersion(computer_info.OsVersion)),
	}
}
