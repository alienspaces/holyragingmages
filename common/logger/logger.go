package logger

import (
	"os"

	"github.com/rs/zerolog"

	"gitlab.com/alienspaces/holyragingmages/common/env"
)

// Hook -
type Hook struct {
	PackageName string
}

// Run -
func (h Hook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Str("package", h.PackageName)
	}
}

// NewLogger returns a logger
func NewLogger(e *env.Env) (*zerolog.Logger, error) {

	err := Init(e)
	if err != nil {
		return nil, err
	}

	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	v := e.Get("APP_LOG_PRETTY")
	if v != "" && v != "0" {
		l = l.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return &l, nil
}

// Init initializes logger
func Init(e *env.Env) error {

	err := e.Init([]env.Item{
		env.Item{Key: "APP_LOG_LEVEL", Required: false},
	}, false)
	if err != nil {
		return err
	}

	// log level
	logLevel := e.Get("APP_LOG_LEVEL")

	// Default log level is error
	zeroLogLevel := zerolog.ErrorLevel

	if logLevel == "debug" {
		zeroLogLevel = zerolog.DebugLevel
	}

	if logLevel == "info" {
		zeroLogLevel = zerolog.InfoLevel
	}

	if logLevel == "warn" {
		zeroLogLevel = zerolog.WarnLevel
	}

	if logLevel == "error" {
		zeroLogLevel = zerolog.ErrorLevel
	}

	zerolog.SetGlobalLevel(zeroLogLevel)

	return nil
}
