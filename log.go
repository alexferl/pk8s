package pk8s

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLog() {
	zerolog.FormattedLevels = map[zerolog.Level]string{
		zerolog.DebugLevel: "DEBUG",
		zerolog.FatalLevel: "ERROR", // use fatal as error for the auto os.Exit(1)
	}
	zerolog.SetGlobalLevel(zerolog.FatalLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, PartsExclude: []string{"time"}})
}
