package store

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
)

func TestNewStore(t *testing.T) {

	// config
	c, err := config.NewConfig([]config.Item{}, false)
	require.NoError(t, err, "NewConfig returns without error")

	configVars := []string{
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
	}
	for _, key := range configVars {
		require.NoError(t, c.Add(key, true), "Add config item")
	}

	l, err := log.NewLogger(c)
	require.NoError(t, err, "NewLogger returns without error")

	// database
	s, err := NewStore(c, l)
	require.Nil(t, err, "NewStore returns without error")
	require.NotNil(t, s, "NewStore returns a store")
}
