package store

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
)

func TestNewStore(t *testing.T) {

	// config
	c, err := config.NewConfig([]config.Item{}, false)
	if err != nil {
		t.Fatalf("Failed new config >%v<", err)
	}

	configVars := []string{
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
	}
	for _, key := range configVars {
		assert.NoError(t, c.Add(config.Item{
			Key:      key,
			Required: true,
		}), "Add configironment item")
	}

	l, err := logger.NewLogger(c)
	if err != nil {
		t.Fatalf("Failed new logger >%v<", err)
	}

	// database
	s, err := NewStore(c, l)
	if assert.Nil(t, err, "NewStore returns without error") {
		assert.NotNil(t, s, "NewStore returns a store")
	}
}
