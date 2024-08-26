package cmd

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/alexferl/pk8s"
)

func example() string {
	var sb strings.Builder

	w := tabwriter.NewWriter(&sb, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "pk8s import cert-manager.crds.yaml\tImports CRDs from a file")
	fmt.Fprintln(w, "pk8s import my_name:=cert-manager.crds.yaml\tImports CRDs from a file using a custom package name")
	fmt.Fprintln(w, "pk8s import https://github.com/cert-manager/cert-manager/releases/download/v1.15.3/cert-manager.crds.yaml\tImports CRDs from an HTTP link")
	fmt.Fprintln(w, "pk8s import my_name:=https://github.com/cert-manager/cert-manager/releases/download/v1.15.3/cert-manager.crds.yaml\tImports CRDs from an HTTP link with a custom package name")
	fmt.Fprintln(w, "kubectl get crds managedcertificates.networking.gke.io -o yaml | go run ./pk8s import\tImports CRDs from a Kubernetes cluster")
	fmt.Fprintln(w, "kubectl get crds managedcertificates.networking.gke.io -o yaml | go run ./pk8s import my_name:=\tImports CRDs from a Kubernetes cluster")
	w.Flush()

	return sb.String()
}

func importCmd() *cobra.Command {
	return &cobra.Command{
		Use:           "import",
		Short:         "Imports API objects by generating Go structs",
		SilenceErrors: true,
		SilenceUsage:  true,
		Example:       example(),
		Args:          cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var path string
			if len(args) > 0 {
				path = args[0]
			}

			overwrite, _ := cmd.Flags().GetBool("overwrite")
			importer := pk8s.NewImporter(&pk8s.ImporterConfig{Overwrite: overwrite})
			data, err := importer.Read(path)
			if err != nil {
				return Err(cmd, err.Error())
			}

			err = importer.Import(data)
			if err != nil {
				return Err(cmd, err.Error())
			}

			return nil
		},
	}
}
