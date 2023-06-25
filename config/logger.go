package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetLogger() zerolog.Logger {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return log.Logger

}
