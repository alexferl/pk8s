package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexferl/pk8s/k8s"
)

func TestNewStack(t *testing.T) {
	app := &App{}
	stack := NewStack(app, "test-stack")

	assert.Equal(t, "test-stack", stack.Name)
	assert.Contains(t, app.stacks, stack)
}

func TestStack_String(t *testing.T) {
	app := &App{}
	stack := NewStack(app, "test-stack")
	chart := NewChart(stack, "test-chart", &ChartParams{
		DisableNameHashes: true,
	})

	deployment := &k8s.DeploymentV1{}

	chart.Append(deployment)

	assert.Equal(t, deploymentYAML, stack.String())
}
