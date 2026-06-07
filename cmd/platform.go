package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system"
	"github.com/tariqajyusuf/ringer/system/platforms"
)

var platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Manage platform package managers",
	Long:  `Commands for listing and installing supported platform package managers.`,
}

var platformListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available platform package managers",
	Long:  `Shows which platform package managers are installed and ready to use on this system.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := io.LoadConfig()
		broker := platforms.NewBroker(verbose, cfg.PreferredPlatform)
		sys_info := system.GetSystemInfo()

		fmt.Printf("Platform managers for %s\n\n", sys_info.Distro)

		if len(broker.Platforms) > 0 {
			fmt.Println("Installed:")
			for key := range broker.Platforms {
				preferred := ""
				if key == broker.PreferredPlatform() {
					preferred = " (preferred)"
				}
				fmt.Printf("  - %s%s\n", key, preferred)
			}
		}

		if skipped := broker.SkippedPlatforms(); len(skipped) > 0 {
			fmt.Println("\nNot installed:")
			for _, key := range skipped {
				fmt.Printf("  - %s  (run: ringer platform install %s)\n", key, key)
			}
		}

		if len(broker.Platforms) == 0 && len(broker.SkippedPlatforms()) == 0 {
			fmt.Println("No platform managers are supported on this OS.")
			fmt.Println("\nTo request support for your platform:")
			fmt.Println("  https://github.com/tariqajyusuf/ringer/issues")
		}
	},
}

func init() {
	rootCmd.AddCommand(platformCmd)
	platformCmd.AddCommand(platformListCmd)
}
