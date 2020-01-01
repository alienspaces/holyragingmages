package database

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/env"
)

func TestNewDatabase(t *testing.T) {

	// environment
	e, err := env.NewEnv([]env.Item{}, false)
	assert.Nil(t, err, "Env initialized without error")

	envVars := []string{
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
	}
	for _, key := range envVars {
		assert.NoError(t, e.Add(env.Item{
			Key:      key,
			Required: true,
		}), "Add environment item")
	}

	l := log.New(os.Stdout,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	// database
	db, err := NewDatabase(l, e)
	if assert.Nil(t, err, "Database initialized without error") {
		assert.NotNil(t, db, "Database connected")
	}
}
