package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/config"
)

func TestLogger(t *testing.T) {

	// config
	c, err := config.NewConfig([]config.Item{}, false)
	assert.Nil(t, err, "Config initialized without error")

	envVars := map[string]string{
		// logger
		"APP_LOG_LEVEL": "debug",
	}
	for key, val := range envVars {
		assert.NoError(t, os.Setenv(key, val), "Set environment value")
	}

	l, err := NewLogger(c)
	assert.NoError(t, err, "NewLogger returns without error")
	assert.NotNil(t, l, "NewLogger is not nil")

	l.Debug("Test level >%s<", "debug")
	l.Info("Test level >%s<", "info")
	l.Warn("Test level >%s<", "warn")
	l.Error("Test level >%s<", "error")

	l.Context("correlation-id", "abcdefg")

	l.Debug("Test level >%s<", "debug")

	l.Context("correlation-id", "hijklmn")

	l.Debug("Test level >%s<", "debug")

	l.Level(ErrorLevel)

	l.Debug("Test level >%s<", "debug")
	l.Info("Test level >%s<", "info")
	l.Warn("Test level >%s<", "warn")
	l.Error("Test level >%s<", "error")
}
