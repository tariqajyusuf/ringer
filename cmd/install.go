/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system/platforms"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a package",
	Long: `Install a package using the preferred underlying package manager for
this system.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a package name to install")
			return
		}
		package_name := args[0]
		broker := platforms.NewBroker()
		guise, err := io.LocateGuise(package_name)
		if err != nil {
			fmt.Printf("Could not locate guise for package %s: %v\n", package_name, err)
			return
		}
		fmt.Printf("%+v\n", guise)
		if platform, ok := guise.Platforms[broker.PreferredPlatform()]; !ok {
			fmt.Printf("Package %s is not available for platform %s\n", package_name, broker.PreferredPlatform())
		} else {
			broker.AddPackage(platform.PackageName)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
