/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system/platforms"
	"gopkg.in/yaml.v3"
)

// circleCmd represents the circle command
var circleCmd = &cobra.Command{
	Use:   "circle [circle file]",
	Short: "Install based on a circle file",
	Long: `Installs all the packages defined in the circle file. A circle file is a declaration
of a desired system state. It contains a list of packages based on their Guise names.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a circle file")
			return
		}
		circle_file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Printf("Could not read circle file: %v\n", err)
			return
		}
		var circle io.Circle
		err = yaml.Unmarshal(circle_file, &circle)
		if err != nil {
			fmt.Printf("Could not parse circle file: %v\n", err)
			return
		}
		fmt.Printf("Found %d packages to install", len(circle.Packages))
		broker := platforms.NewBroker()
		for _, pkg := range circle.Packages {
			addHelper(broker, pkg)
		}
	},
}

func init() {
	rootCmd.AddCommand(circleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// circleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// circleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
