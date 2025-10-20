/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tariqajyusuf/ringer/system"
	"github.com/tariqajyusuf/ringer/system/platforms"
)

// platformsCmd represents the platforms command
var platformsCmd = &cobra.Command{
	Use:   "platforms",
	Short: "Shows supported platforms",
	Long: `Checks the system and identifies supported platforms for
package management.`,
	Run: func(cmd *cobra.Command, args []string) {
		sys_info := system.GetSystemInfo()
		broker := platforms.NewBroker()

		fmt.Printf("Platforms supported for %s\n", sys_info.Distro)
		for key := range broker.Platforms {
			fmt.Printf(" - %s\n", key)
		}
	},
}

func init() {
	rootCmd.AddCommand(platformsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// platformsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// platformsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
