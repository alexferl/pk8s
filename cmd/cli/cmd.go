package cli

import (
	"os/exec"
)

type Executor interface {
	Init(path string) ([]byte, error)
	Tidy(path string) ([]byte, error)
}

type CommandExecutor struct{}

func (c *CommandExecutor) Init(path string) ([]byte, error) {
	cmd := exec.Command("go", "mod", "init", path)
	cmd.Dir = path
	return cmd.CombinedOutput()
}

func (c *CommandExecutor) Tidy(path string) ([]byte, error) {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = path
	return cmd.CombinedOutput()
}
