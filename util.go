package pk8s

import (
	"fmt"
	"os"

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
