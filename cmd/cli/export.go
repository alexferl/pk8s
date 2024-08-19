package cli

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func exportCmd() *cobra.Command {
	return &cobra.Command{
		Use:           "export",
		Short:         "Generates Kubernetes manifests for all stacks and charts in your app",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			c := exec.Command("go", "run", path)
			output, err := c.CombinedOutput()
			if err != nil {
				msg := fmt.Sprintf("command execution failed: %v", err)
				return Err(cmd, msg)
			}

			Msg(cmd, string(output))

			return nil
		},
	}
}
