package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system"
	"github.com/tariqajyusuf/ringer/system/platforms"
)

var platformInstallCmd = &cobra.Command{
	Use:   "install [platform name]",
	Short: "Install a platform package manager",
	Long:  `Installs the specified platform package manager so it can be used with ringer.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a platform name (e.g. homebrew, winget)")
			return
		}
		name := args[0]

		// Build a map of all possible platforms to find the one requested.
		possible := map[string]platforms.Platform{
			"homebrew": &platforms.Homebrew{},
			"winget":   &platforms.Winget{},
		}

		platform, ok := possible[name]
		if !ok {
			fmt.Printf("Unknown platform %q. Supported platforms: homebrew, winget\n", name)
			return
		}

		if !platform.EnabledForSystem(system.GetSystemInfo()) {
			fmt.Printf("Platform %q is not supported on this operating system.\n", name)
			return
		}

		if platform.IsInstalled() {
			fmt.Printf("%s is already installed.\n", name)
		} else {
			fmt.Printf("Installing %s...\n", name)
			if err := platform.SetupPackageManager(verbose); err != nil {
				fmt.Printf("Could not install %s: %s\n", name, err)
				return
			}
			fmt.Printf("Successfully installed %s.\n", name)
			fmt.Println("\nRestart your shell to pick up any new PATH changes.")
		}

		cfg, _ := io.LoadConfig()
		cfg.PreferredPlatform = name
		if err := io.SaveConfig(cfg); err != nil {
			fmt.Printf("Warning: could not save preferred platform to config: %s\n", err)
		} else {
			fmt.Printf("Set %s as your preferred platform.\n", name)
		}
	},
}

func init() {
	platformCmd.AddCommand(platformInstallCmd)
}
