package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Stack struct {
	Name string

	app    *App
	charts []*Chart
	sb     strings.Builder
}

func NewStack(app *App, name string) *Stack {
	s := &Stack{
		Name:   name,
		app:    app,
		charts: make([]*Chart, 0),
	}
	app.stacks = append(app.stacks, s)
	return app.stacks[len(app.stacks)-1]
}

func (s *Stack) build(writeToFs bool) string {
	var sb strings.Builder
	for _, chart := range s.charts {
		sb.WriteString(chart.String())

		if writeToFs {
			path := fmt.Sprintf("%s/%s", *s.app.config.OutputPath, s.Name)
			err := createDirectory(path)
			if err != nil {
				log.Fatal().Msgf("failed creating directory: %v", err)
			}

			chartPath := fmt.Sprintf("%s/%s%s", path, chart.Name, *s.app.config.OutputFileExtension)
			fmt.Printf("    - %s\n", chartPath)
			err = os.WriteFile(chartPath, []byte(sb.String()), 0o644)
			if err != nil {
				log.Fatal().Msgf("failed writing chart: %v", err)
			}

			log.Debug().Msgf("wrote chart '%s'", chartPath)
		}

		s.sb.WriteString(sb.String())
		sb.Reset()
	}

	return s.sb.String()
}

// String returns a string of all YAML objects across all charts in this stack.
func (s *Stack) String() string {
	return s.build(false)
}
