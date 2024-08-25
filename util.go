package pk8s

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func removeDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	err = os.Mkdir(path, 0o755)
	if err != nil {
		return err
	}

	return nil
}

func createDirectory(path string, errOnExist bool) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0o755)
		if err != nil {
			return err
		}
		log.Debug().Msgf("directory '%s' successfully created", path)
		return nil
	}
	log.Debug().Msgf("directory '%s' already exists", path)
	if errOnExist {
		return fmt.Errorf("directory '%s' already exists", path)
	}
	return nil
}

func RunGoFmt(dir string) error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go fmt failed: %v", err)
	}
	return nil
}

func RunGoImports(dir string) error {
	cmd := exec.Command("goimports", "-w", ".")
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("goimports failed: %v", err)
	}
	return nil
}
