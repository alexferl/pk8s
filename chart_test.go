package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexferl/pk8s/k8s"
)

func TestNewChart(t *testing.T) {
	stack := &Stack{}
	namespace := "custom"
	chart := NewChart(stack, "test-chart", &ChartParams{Namespace: &namespace})

	assert.Equal(t, "test-chart", chart.Name)
	assert.Equal(t, namespace, *chart.Params.Namespace)
	assert.Contains(t, stack.charts, chart)
}

func TestNewChart_NilParams(t *testing.T) {
	stack := &Stack{}
	chart := NewChart(stack, "test-chart", nil)

	assert.Equal(t, "test-chart", chart.Name)
	assert.Contains(t, stack.charts, chart)
}

func TestChart_Append(t *testing.T) {
	chart := &Chart{
		Name: "test-chart",
		Params: &ChartParams{
			Namespace: k8s.String("test"),
			Labels:    &map[string]string{"app": "test"},
		},
	}

	deployment := &k8s.DeploymentV1{}

	chart.Append(deployment)

	assert.Equal(t, "test-chart", (*deployment.Metadata.Name)[:10])
	assert.Equal(t, "test", *deployment.Metadata.Namespace)
	assert.Equal(t, "test", (*deployment.Metadata.Labels)["app"])
}

func TestChart_String(t *testing.T) {
	stack := &Stack{}
	chart := NewChart(stack, "test-chart", &ChartParams{
		DisableNameHashes: true,
	})

	deployment := &k8s.DeploymentV1{}

	chart.Append(deployment)

	assert.Equal(t, deploymentYAML, chart.String())
}

var chartYAML = deploymentYAML + "---\napiVersion: v1\nkind: Service\nmetadata:\n  name: test-chart\n"

func TestChart_String_Multiple(t *testing.T) {
	stack := &Stack{}
	chart := NewChart(stack, "test-chart", &ChartParams{
		DisableNameHashes: true,
	})

	deployment := &k8s.DeploymentV1{}
	service := &k8s.ServiceV1{}

	chart.Append(deployment, service)

	assert.Equal(t, chartYAML, chart.String())
}

type MockAPIObject struct {
	Metadata *MockMetadata
}

type MockMetadata struct {
	Name      string
	Namespace string
	Labels    map[string]string
}

func Test_modifyMetadataField(t *testing.T) {
	obj := &MockAPIObject{
		Metadata: &MockMetadata{},
	}

	err := modifyMetadataField(obj, "Name", "new-name")
	assert.NoError(t, err)
	assert.Equal(t, "new-name", obj.Metadata.Name)

	err = modifyMetadataField(obj, "Namespace", "new-namespace")
	assert.NoError(t, err)
	assert.Equal(t, "new-namespace", obj.Metadata.Namespace)

	err = modifyMetadataField(obj, "Labels", map[string]string{"key": "value"})
	assert.NoError(t, err)
	assert.Equal(t, "value", obj.Metadata.Labels["key"])

	err = modifyMetadataField(123, "Name", "new-name")
	assert.Error(t, err)

	err = modifyMetadataField(obj, "NonExistentField", "value")
	assert.Error(t, err)

	// test namespace shouldn't be overwritten if already set
	obj = &MockAPIObject{
		Metadata: &MockMetadata{Namespace: "custom"},
	}

	err = modifyMetadataField(obj, "Namespace", "new-namespace")
	assert.NoError(t, err)
	assert.Equal(t, "custom", obj.Metadata.Namespace)
}
