package pk8s

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexferl/pk8s/k8s"
)

func TestNew(t *testing.T) {
	app := New()

	assert.Equal(t, DefaultConfig.OutputPath, app.config.OutputPath)
	assert.Equal(t, DefaultConfig.OutputFileExtension, app.config.OutputFileExtension)
	assert.Empty(t, app.stacks)
}

func TestNewWithConfig(t *testing.T) {
	config := AppConfig{OutputPath: k8s.String("custom"), OutputFileExtension: k8s.String(".k8s.yaml")}
	app := NewWithConfig(config)

	assert.Equal(t, "custom", *app.config.OutputPath)
	assert.Equal(t, ".k8s.yaml", *app.config.OutputFileExtension)
	assert.Empty(t, app.stacks)
}

func TestNewWithConfig_Empty(t *testing.T) {
	config := AppConfig{}
	app := NewWithConfig(config)

	assert.Empty(t, app.stacks)
}

var deploymentYAML = "---\napiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: test-chart\n"

func TestApp_Export(t *testing.T) {
	tempDir := t.TempDir()

	config := AppConfig{OutputPath: &tempDir}
	app := NewWithConfig(config)
	stack := NewStack(app, "test-stack")
	chart := NewChart(stack, "test-chart", &ChartParams{DisableNameHashes: true})

	filePath := fmt.Sprintf("%s/%s/%s.yaml", tempDir, stack.Name, chart.Name)
	deployment := &k8s.DeploymentV1{}
	chart.Append(deployment)

	app.Export()

	assert.DirExists(t, tempDir)
	assert.FileExists(t, filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed reading file: %v", err)
	}
	assert.Equal(t, deploymentYAML, string(data))
}

func TestApp_Print(t *testing.T) {
	app := New()
	stack := NewStack(app, "test-stack")
	chart := NewChart(stack, "test-chart", &ChartParams{DisableNameHashes: true})
	deployment := &k8s.DeploymentV1{}
	chart.Append(deployment)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	app.Print()

	w.Close()
	os.Stdout = old

	var buf strings.Builder
	_, _ = io.Copy(&buf, r)

	assert.Equal(t, deploymentYAML, buf.String())
}

func TestApp_String(t *testing.T) {
	app := New()
	stack := NewStack(app, "test-stack")
	chart := NewChart(stack, "test-chart", &ChartParams{DisableNameHashes: true})
	deployment := &k8s.DeploymentV1{}
	chart.Append(deployment)

	assert.Equal(t, deploymentYAML, app.String())
}
