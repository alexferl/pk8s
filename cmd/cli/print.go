package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Err(cmd *cobra.Command, msg string) error {
	red := color.New(color.FgRed)
	red.Fprint(cmd.OutOrStderr(), "ERROR ")

	bold := color.New(color.Bold)
	bold.Fprintln(cmd.OutOrStderr(), msg)

	return fmt.Errorf(msg)
}

func Msg(cmd *cobra.Command, msg string) {
	fmt.Fprintf(cmd.OutOrStdout(), msg)
}
