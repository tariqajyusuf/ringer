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

// removeCmd represents the remove command
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
		broker := platforms.NewBroker()
		guise, err := io.LocateGuise(package_name)
		if err != nil {
			fmt.Printf("Could not locate guise for package %s: %v\n", package_name, err)
			return
		}
		fmt.Printf("%+v\n", guise)
		if platform, ok := guise.Platforms[broker.PreferredPlatform()]; !ok {
			fmt.Printf("Package %s is not defined for platform %s\n", package_name, broker.PreferredPlatform())
		} else {
			broker.RemovePackage(platform.PackageName)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
