package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
)

// Log -
type Log struct {
	log    *logrus.Logger
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

var levelMap = map[Level]logrus.Level{
	// DebugLevel -
	DebugLevel: logrus.DebugLevel,
	// InfoLevel -
	InfoLevel: logrus.InfoLevel,
	// WarnLevel -
	WarnLevel: logrus.WarnLevel,
	// ErrorLevel -
	ErrorLevel: logrus.ErrorLevel,
}

// NewLogger returns a logger
func NewLogger(c configurer.Configurer) (*Log, error) {

	l := Log{
		log:    logrus.New(),
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

	// create a new instance of the logger
	l.log.SetFormatter(&logrus.JSONFormatter{})

	// output to stdout instead of the default stderr
	l.log.SetOutput(os.Stdout)

	// log level
	configLevel := l.Config.Get("APP_LOG_LEVEL")
	switch configLevel {
	case "debug", "Debug", "DEBUG":
		l.log.SetLevel(DebugLevel)
	case "info", "Info", "INFO":
		l.log.SetLevel(InfoLevel)
	case "warn", "Warn", "WARN":
		l.log.SetLevel(WarnLevel)
	case "error", "Error", "ERROR":
		l.log.SetLevel(ErrorLevel)
	default:
		l.log.SetLevel(DebugLevel)
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
		l.log.SetLevel(lvl)
	}
}

// Debug -
func (l *Log) Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.log.Debug(msg)
}

// Info -
func (l *Log) Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.log.Info(msg)
}

// Warn -
func (l *Log) Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.log.Warn(msg)
}

// Error -
func (l *Log) Error(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.log.Error(msg)
}
