package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Config -
type Config interface {
	Get(key string) string
}

// Logger -
type Logger struct {
	log *logrus.Logger
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
func NewLogger(c Config) (*Logger, error) {

	l := Logger{
		log: logrus.New(),
	}

	err := l.Init(c)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// Init initializes logger
func (l *Logger) Init(c Config) error {

	// Create a new instance of the logger. You can have any number of instances.
	l.log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	l.log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	l.log.SetLevel(DebugLevel)

	return nil
}

// Printf -
func (l *Logger) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

// Level -
func (l *Logger) Level(level Level) {
	if lvl, ok := levelMap[level]; ok {
		l.log.SetLevel(lvl)
	}
}

// Debug -
func (l *Logger) Debug(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args)
	l.log.Debug(msg)
}

// Info -
func (l *Logger) Info(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args)
	l.log.Info(msg)
}

// Warn -
func (l *Logger) Warn(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args)
	l.log.Warn(msg)
}

// Error -
func (l *Logger) Error(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args)
	l.log.Error(msg)
}
