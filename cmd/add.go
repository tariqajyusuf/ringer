package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ringerio "github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system"
	"github.com/tariqajyusuf/ringer/system/platforms"
)

var addCmd = &cobra.Command{
	Use:   "add [package name]",
	Short: "Adds a package",
	Long: `Add a package using the preferred underlying package manager for
this system.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a package name to add")
			return
		}
		package_name := args[0]
		pkg, err := ringerio.LocatePackage(package_name)
		if err != nil {
			fmt.Printf("Could not locate package %s: %v\n", package_name, err)
			return
		}
		if err := pkg.CheckOSAllowed(system.GetSystemInfo().Kernel); err != nil {
			fmt.Println(err)
			return
		}
		cfg, _ := ringerio.LoadConfig()
		broker := platforms.NewBroker(verbose, cfg.PreferredPlatform)
		if len(broker.Platforms) == 0 {
			printNoPlatformMessage(broker)
			return
		}
		addHelper(broker, pkg)
	},
}

func addHelper(broker *platforms.Broker, pkg *ringerio.Package) {
	if platform, ok := pkg.Platforms[broker.PreferredPlatform()]; !ok {
		fmt.Printf("Package %s is not defined for platform %s\n", pkg.Name, broker.PreferredPlatform())
	} else if err := broker.AddPackage(platform.PackageName); err != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
