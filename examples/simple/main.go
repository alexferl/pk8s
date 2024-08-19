package main

import (
	"github.com/alexferl/pk8s"
	"github.com/alexferl/pk8s/k8s"
)

type WebServiceParams struct {
	pk8s.ChartParams
	Replicas int32
}

func NewWebService(stack *pk8s.Stack, name string, params WebServiceParams) *pk8s.Chart {
	chart := pk8s.NewChart(stack, name, &params.ChartParams)

	service := &k8s.ServiceV1{
		Spec: &k8s.ServiceSpecV1{
			Ports:    []k8s.ServicePortV1{{Port: 80, TargetPort: k8s.String("http")}},
			Selector: params.ChartParams.Labels,
			Type:     k8s.String("LoadBalancer"),
		},
	}

	deployment := &k8s.DeploymentV1{
		Spec: &k8s.DeploymentSpecV1{
			Replicas: k8s.Int32(params.Replicas),
			Selector: k8s.LabelSelectorV1{
				MatchLabels: params.ChartParams.Labels,
			},
			Template: k8s.PodTemplateSpecV1{
				Metadata: &k8s.ObjectMetaV1{
					Labels: params.ChartParams.Labels,
				},
				Spec: &k8s.PodSpecV1{
					Containers: []k8s.ContainerV1{
						{
							Name:  "nginx",
							Image: k8s.String("nginx:1.27.1"),
							Ports: []k8s.ContainerPortV1{
								{
									Name:          k8s.String("http"),
									Protocol:      k8s.String("TCP"),
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	chart.Append(deployment, service)

	return chart
}

func main() {
	app := pk8s.New()
	stack := pk8s.NewStack(app, "dev")

	name := "hello"
	labels := map[string]string{
		"app": name,
	}
	NewWebService(stack, name, WebServiceParams{
		ChartParams: pk8s.ChartParams{
			Labels: &labels,
		},
		Replicas: 2,
	})

	app.Export()
}
