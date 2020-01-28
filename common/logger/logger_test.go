package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/alienspaces/holyragingmages/common/config"
)

func TestLogger(t *testing.T) {

	// environment
	e, err := config.NewConfig([]config.Item{}, false)
	assert.Nil(t, err, "Config initialized without error")

	envVars := map[string]string{
		// logger
		"APP_LOG_LEVEL": "debug",
	}
	for key, val := range envVars {
		assert.NoError(t, os.Setenv(key, val), "Set environment value")
	}

	l, err := NewLogger(e)
	assert.NoError(t, err, "NewLogger returns without error")
	assert.NotNil(t, l, "NewLogger is not nil")

	l.Debug("Test debug")
	l.Info("Test info")
	l.Warn("Test warn")
	l.Error("Test error")
}
