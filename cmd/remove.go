package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ringerio "github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system"
	"github.com/tariqajyusuf/ringer/system/platforms"
)

var removeCmd = &cobra.Command{
	Use:   "remove [package name]",
	Short: "Remove a package",
	Long: `Remove a package using the preferred underlying package manager for
this system.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a package name to remove")
			return
		}
		package_name := args[0]
		cfg, _ := ringerio.LoadConfig()
		broker := platforms.NewBroker(verbose, cfg.PreferredPlatform)
		if len(broker.Platforms) == 0 {
			printNoPlatformMessage(broker)
			return
		}
		pkg, err := ringerio.LocatePackage(package_name)
		if err != nil {
			fmt.Printf("Could not locate package %s: %v\n", package_name, err)
			return
		}
		sysinfo := system.GetSystemInfo()
		if err := pkg.CheckOSAllowed(sysinfo.Kernel); err != nil {
			fmt.Println(err)
			return
		}
		if verbose {
			fmt.Printf("%+v\n", pkg)
		}
		if platform, ok := pkg.Platforms[broker.PreferredPlatform()]; !ok {
			fmt.Printf("Package %s is not defined for platform %s\n", package_name, broker.PreferredPlatform())
		} else if err := broker.RemovePackage(platform.PackageName); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
