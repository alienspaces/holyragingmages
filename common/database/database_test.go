package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/env"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
)

func TestNewDatabase(t *testing.T) {

	// environment
	e, err := env.NewEnv([]env.Item{}, false)
	assert.Nil(t, err, "Env initialized without error")

	envVars := map[string]string{
		// logger
		"APP_LOG_LEVEL": "debug",
		// database
		"APP_DB_HOST":     "postgres",
		"APP_DB_PORT":     "5432",
		"APP_DB_NAME":     "test",
		"APP_DB_USER":     "test_user",
		"APP_DB_PASSWORD": "test_user_password",
	}
	for key, val := range envVars {
		assert.NoError(t, os.Setenv(key, val), "Set environment value")
	}

	// logger
	l, err := logger.NewLogger(e)
	assert.NotNil(t, l, "Logger initialized without error")

	// database
	db, err := NewDatabase(l, e)
	if assert.Nil(t, err, "Database initialized without error") {
		assert.NotNil(t, db, "Database connected")
	}
}
