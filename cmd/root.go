package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "pk8s",
	Version: "0.0.0",
}

func init() {
	importCmd := importCmd()
	rootCmd.AddCommand(initCmd(&CommandExecutor{}))
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(exportCmd())

	importCmd.PersistentFlags().Bool("overwrite", false, "overwrite existing custom resource definition files")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
