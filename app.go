package pk8s

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/alexferl/pk8s/k8s"
)

func init() {
	InitLog()
}

type App struct {
	config AppConfig
	stacks []*Stack
}

type AppConfig struct {
	// OutputPath defines the directory to export to.
	// Optional. Default: "dist".
	OutputPath *string

	// OutputFileExtension defines the file extension to use for generated YAML files.
	// Optional. Default: ".yaml".
	OutputFileExtension *string
}

var DefaultConfig = AppConfig{
	OutputPath:          k8s.String("dist"),
	OutputFileExtension: k8s.String(".yaml"),
}

func newApp(config AppConfig) *App {
	return &App{
		config: config,
		stacks: make([]*Stack, 0),
	}
}

func New() *App {
	return newApp(DefaultConfig)
}

func NewWithConfig(config AppConfig) *App {
	if config.OutputPath == nil {
		config.OutputPath = DefaultConfig.OutputPath
	}

	if config.OutputFileExtension == nil {
		config.OutputFileExtension = DefaultConfig.OutputFileExtension
	}

	return newApp(config)
}

func (a *App) build(writeToFs bool, printManifestList bool) string {
	var sb strings.Builder

	if writeToFs {
		err := removeDirectory(*a.config.OutputPath)
		if err != nil {
			log.Fatal().Msgf("failed clearing directory: %v", err)
		}
	}

	if printManifestList {
		fmt.Println("Generated Kubernetes manifests:")
	}
	for _, stack := range a.stacks {
		err := createDirectory(*a.config.OutputPath, false)
		if err != nil {
			log.Fatal().Msgf("failed creating directory: %v", err)
		}
		if printManifestList {
			fmt.Printf("  %s:\n", stack.Name)
		}
		sb.WriteString(stack.build(writeToFs))
	}

	return sb.String()
}

// Export exports all manifests to the output directory.
func (a *App) Export() {
	a.build(true, true)
}

// Print prints all YAML objects across all stacks and charts in this app.
func (a *App) Print() {
	fmt.Printf("%s", a.build(false, false))
}

// String returns all YAML objects across all stacks and charts in this app.
func (a *App) String() string {
	return a.build(false, false)
}
