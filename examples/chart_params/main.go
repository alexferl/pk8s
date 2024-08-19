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

	name := "hello"
	labels := map[string]string{
		"app": name,
	}
	NewChart(stack, name, &ChartParams{
		pk8s.ChartParams{
			DisableNameHashes: true,
			Labels:            &labels,
			Namespace:         k8s.String("custom"),
		},
	})

	app.Export()
}
