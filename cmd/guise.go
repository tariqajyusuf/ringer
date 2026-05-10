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

// guiseCmd represents the guise command
var guiseCmd = &cobra.Command{
	Use:   "guise [guise file]",
	Short: "Install based on a guise file",
	Long: `Installs all the packages defined in the guise file. A guise file is a declaration
of a desired system state. It contains a list of packages based on their package names.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a guise file")
			return
		}
		guise_file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Printf("Could not read guise file: %v\n", err)
			return
		}
		var guise io.Guise
		err = yaml.Unmarshal(guise_file, &guise)
		if err != nil {
			fmt.Printf("Could not parse guise file: %v\n", err)
			return
		}
		fmt.Printf("Found %d packages to install", len(guise.Packages))
		broker := platforms.NewBroker()
		for _, pkg := range guise.Packages {
			addHelper(broker, pkg)
		}
	},
}

func init() {
	rootCmd.AddCommand(guiseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// guiseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guiseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
