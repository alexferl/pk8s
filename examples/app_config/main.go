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
	app := pk8s.NewWithConfig(pk8s.AppConfig{
		OutputPath:          k8s.String("custom"),
		OutputFileExtension: k8s.String(".gen.yaml"),
	})
	stack := pk8s.NewStack(app, "dev")
	NewChart(stack, "hello", nil)

	app.Export()
}
