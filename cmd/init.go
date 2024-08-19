package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var initMsg = `
  Your pk8s Go project is ready! 🚀✨

    📦 pk8s export  Export Kubernetes manifests to dist/

  Deploy:
    🛠️ kubectl apply -f dist/

`

var mainGo = `
package main

import (
	"github.com/alexferl/pk8s"
	"github.com/alexferl/pk8s/k8s"
)

type ChartParams struct {
	pk8s.ChartParams
}

func NewChart(stack *pk8s.Stack, name string, params *ChartParams) *pk8s.Chart {
	var chartParams pk8s.ChartParams
	if params != nil {
		chartParams = params.ChartParams
	}
	chart := pk8s.NewChart(stack, name, &chartParams)

	deployment := &k8s.DeploymentV1{
		// ...
	}

	chart.Append(deployment)

	return chart
}

func main() {
	app := pk8s.New()
	stack := pk8s.NewStack(app, "dev")

	NewChart(stack, "hello", nil)

	app.Export()
}
`

func initCmd(executor Executor) *cobra.Command {
	return &cobra.Command{
		Use:           "init",
		Short:         "Create a new pk8s project",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			empty, err := isDirectoryEmpty(path)
			if err != nil {
				return Err(cmd, err.Error())
			}

			if !empty {
				var msg string
				if path != "." {
					msg = fmt.Sprintf("directory '%s' is not empty", path)
				} else {
					msg = "current directory is not empty"
				}

				return Err(cmd, msg)
			}

			err = os.WriteFile(fmt.Sprintf("%s/main.go", path), []byte(mainGo), 0o755)
			if err != nil {
				return Err(cmd, err.Error())
			}

			output, err := executor.Init(path)
			if err != nil {
				Msg(cmd, string(output))
				return Err(cmd, fmt.Sprintf("running go mod init: %v", err))
			}

			output, err = executor.Tidy(path)
			if err != nil {
				Msg(cmd, string(output))
				return Err(cmd, fmt.Sprintf("running go mod tidy: %v", err))
			}

			Msg(cmd, initMsg)

			return nil
		},
	}
}

func isDirectoryEmpty(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0o755)
		if err != nil {
			return false, fmt.Errorf("failed to create directory: %w", err)
		}
		return true, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to stat path: %w", err)
	}

	if !info.IsDir() {
		return false, fmt.Errorf("path is not a directory")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return false, fmt.Errorf("failed to read directory: %w", err)
	}
	return len(entries) == 0, nil
}
