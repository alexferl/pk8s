package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "pk8s",
	Version: "0.0.0",
}

func init() {
	rootCmd.AddCommand(exportCmd())
	rootCmd.AddCommand(initCmd(&CommandExecutor{}))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
