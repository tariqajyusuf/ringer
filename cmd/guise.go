package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	ringerio "github.com/tariqajyusuf/ringer/io"
	"github.com/tariqajyusuf/ringer/system"
	"github.com/tariqajyusuf/ringer/system/platforms"
	"gopkg.in/yaml.v3"
)

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
		var guise ringerio.Guise
		err = yaml.Unmarshal(guise_file, &guise)
		if err != nil {
			fmt.Printf("Could not parse guise file: %v\n", err)
			return
		}
		cfg, _ := ringerio.LoadConfig()
		broker := platforms.NewBroker(verbose, cfg.PreferredPlatform)
		if len(broker.Platforms) == 0 {
			printNoPlatformMessage(broker)
			return
		}
		kernel := system.GetSystemInfo().Kernel
		fmt.Printf("Found %d packages to install\n", len(guise.Packages))
		for _, pkgName := range guise.Packages {
			pkg, err := ringerio.LocatePackage(pkgName)
			if err != nil {
				fmt.Printf("Could not locate package %s: %v\n", pkgName, err)
				continue
			}
			if err := pkg.CheckOSAllowed(kernel); err != nil {
				fmt.Println(err)
				continue
			}
			addHelper(broker, pkg)
		}
	},
}

func init() {
	rootCmd.AddCommand(guiseCmd)
}
