package service

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
	"gitlab.com/alienspaces/holyragingmages/common/prepare"
	"gitlab.com/alienspaces/holyragingmages/common/store"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
)

// TestRunner - allow Init and Run functions to be defined by tests
type TestRunner struct {
	Runner
	InitFunc func(c configurer.Configurer, l logger.Logger, s storer.Storer, p preparer.Preparer) error
	RunFunc  func(args map[string]interface{}) error
}

func (rnr *TestRunner) Init(c configurer.Configurer, l logger.Logger, s storer.Storer, p preparer.Preparer) error {
	if rnr.InitFunc == nil {
		return rnr.Runner.Init(c, l, s, p)
	}
	return rnr.InitFunc(c, l, s, p)
}

func (rnr *TestRunner) Run(args map[string]interface{}) error {
	if rnr.RunFunc == nil {
		return rnr.Runner.Run(args)
	}
	return rnr.RunFunc(args)
}

// NewDefaultDependencies -
func NewDefaultDependencies() (configurer.Configurer, logger.Logger, storer.Storer, preparer.Preparer, error) {

	// configurer
	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	configVars := []string{
		// logger
		"APP_LOG_LEVEL",
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
	}
	for _, key := range configVars {
		err = c.Add(key, true)
		if err != nil {
			return nil, nil, nil, nil, err
		}
	}

	// logger
	l, err := log.NewLogger(c)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// storer
	s, err := store.NewStore(c, l)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// preparer
	p, err := prepare.NewPrepare(l)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return c, l, s, p, nil
}

func TestNewService(t *testing.T) {

	c, l, s, p, err := NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	tr := TestRunner{}

	ts, err := NewService(c, l, s, p, &tr)
	require.NoError(t, err, "NewService returns without error")
	require.NotNil(t, ts, "Test service is not nil")
}
