/*
Copyright © 2025 Tariq Yusuf <tariq@tariqyusuf.in>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tariqajyusuf/ringer/system/platforms"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ringer",
	Short: "A universal package manager and workstation setup tool",
	Long: `Ringer is a universal install command that bridges many of the common
platforms. When setting up new systems, there are usually several pieces of
software to install to get your environment just right. Ringer helps bridge that
setup process across multiple platforms.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print detailed output")
}

func printNoPlatformMessage(broker *platforms.Broker) {
	fmt.Println("No platform managers are installed on this system.")
	if skipped := broker.SkippedPlatforms(); len(skipped) > 0 {
		fmt.Println("\nSupported managers for your OS (not yet installed):")
		for _, p := range skipped {
			fmt.Printf("  - %s  (run: ringer platform install %s)\n", p, p)
		}
	}
	fmt.Println("\nTo request support for a new platform:")
	fmt.Println("  https://github.com/tariqajyusuf/ringer/issues")
}
