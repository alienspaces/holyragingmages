package log

import (
	"os"

	"github.com/rs/zerolog"

	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
)

// Log -
type Log struct {
	log    zerolog.Logger
	fields map[string]interface{}
	Config configurer.Configurer
}

// Level -
type Level uint32

const (
	// DebugLevel -
	DebugLevel = 5
	// InfoLevel -
	InfoLevel = 4
	// WarnLevel -
	WarnLevel = 3
	// ErrorLevel -
	ErrorLevel = 2
)

var levelMap = map[Level]zerolog.Level{
	// DebugLevel -
	DebugLevel: zerolog.DebugLevel,
	// InfoLevel -
	InfoLevel: zerolog.InfoLevel,
	// WarnLevel -
	WarnLevel: zerolog.WarnLevel,
	// ErrorLevel -
	ErrorLevel: zerolog.ErrorLevel,
}

// NewLogger returns a logger
func NewLogger(c configurer.Configurer) (*Log, error) {

	l := Log{
		log:    zerolog.New(os.Stdout).With().Timestamp().Logger(),
		fields: make(map[string]interface{}),
		Config: c,
	}

	err := l.Init()
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// Init initializes logger
func (l *Log) Init() error {

	// log level
	configLevel := l.Config.Get("APP_LOG_LEVEL")
	switch configLevel {
	case "debug", "Debug", "DEBUG":
		l.log.Level(DebugLevel)
	case "info", "Info", "INFO":
		l.log.Level(InfoLevel)
	case "warn", "Warn", "WARN":
		l.log.Level(WarnLevel)
	case "error", "Error", "ERROR":
		l.log.Level(ErrorLevel)
	default:
		l.log.Level(DebugLevel)
	}

	return nil
}

// Printf -
func (l *Log) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

// Level -
func (l *Log) Level(level Level) {
	if lvl, ok := levelMap[level]; ok {
		l.log = l.log.Level(lvl)
	}
}

// Context - set logging
func (l *Log) Context(key, value string) {
	l.fields[key] = value
}

// Debug -
func (l *Log) Debug(msg string, args ...interface{}) {
	ctxLog := l.log.With().Fields(l.fields).Logger()
	ctxLog.Debug().Msgf(msg, args...)
}

// Info -
func (l *Log) Info(msg string, args ...interface{}) {
	ctxLog := l.log.With().Fields(l.fields).Logger()
	ctxLog.Info().Msgf(msg, args...)
}

// Warn -
func (l *Log) Warn(msg string, args ...interface{}) {
	ctxLog := l.log.With().Fields(l.fields).Logger()
	ctxLog.Warn().Msgf(msg, args...)
}

// Error -
func (l *Log) Error(msg string, args ...interface{}) {
	ctxLog := l.log.With().Fields(l.fields).Logger()
	ctxLog.Error().Msgf(msg, args...)
}
