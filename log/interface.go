package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type console struct {
	appName string
	logger  *zerolog.Logger
}

// NewConsoleLogger creates a new console logger
func NewConsoleLogger(appName string) ILogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &console{
		appName: appName,
	}
}

func (c *console) Info(msg string) {
	log.Info().Str("app", c.appName).Msg(msg)
}

func (c *console) Error(err error, msg string) {
	log.Error().Str("app", c.appName).Err(err).Msg(msg)
}
