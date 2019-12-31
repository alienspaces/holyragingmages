package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/env"
)

func TestLogger(t *testing.T) {

	// environment
	e, err := env.NewEnv([]env.Item{}, false)
	assert.Nil(t, err, "Env initialized without error")

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

	l.Debug().Msg("Test debug")
	l.Info().Msg("Test info")
	l.Warn().Msg("Test warn")
	l.Error().Msg("Test error")
}
