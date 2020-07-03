package runner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/server/core/config"
	"gitlab.com/alienspaces/holyragingmages/server/core/log"
	"gitlab.com/alienspaces/holyragingmages/server/core/store"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/storer"
	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/harness"
)

// NewDefaultDependencies -
func NewDefaultDependencies() (configurer.Configurer, logger.Logger, storer.Storer, error) {

	// configurer
	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, nil, err
	}

	configVars := []string{
		// logger
		"APP_SERVER_LOG_LEVEL",
		// database
		"APP_SERVER_DB_HOST",
		"APP_SERVER_DB_PORT",
		"APP_SERVER_DB_NAME",
		"APP_SERVER_DB_USER",
		"APP_SERVER_DB_PASSWORD",
	}
	for _, key := range configVars {
		err = c.Add(key, true)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	// logger
	l, err := log.NewLogger(c)
	if err != nil {
		return nil, nil, nil, err
	}

	// storer
	s, err := store.NewStore(c, l)
	if err != nil {
		return nil, nil, nil, err
	}

	err = s.Init()
	if err != nil {
		return nil, nil, nil, err
	}

	return c, l, s, nil
}

func NewTestHarness() (*harness.Testing, error) {

	// harness
	config := harness.DataConfig{}

	h, err := harness.NewTesting(config)
	if err != nil {
		return nil, err
	}

	// harness commit data
	h.CommitData = true

	return h, nil
}

func TestNewRunner(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	r := NewRunner()

	err = r.Init(c, l, s)
	require.NoError(t, err, "Init returns without error")
}
