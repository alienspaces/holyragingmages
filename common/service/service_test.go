package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/database"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
)

// TestRunner - allow Init and Run functions to be defined by tests
type TestRunner struct {
	Runner
	InitFunc func(c Configurer, l Logger, s Storer) error
	RunFunc  func(args map[string]interface{}) error
}

func (rnr *TestRunner) Init(c Configurer, l Logger, s Storer) error {
	if rnr.InitFunc == nil {
		return rnr.Runner.Init(c, l, s)
	}
	return rnr.InitFunc(c, l, s)
}

func (rnr *TestRunner) Run(args map[string]interface{}) error {
	if rnr.RunFunc == nil {
		return rnr.Runner.Run(args)
	}
	return rnr.RunFunc(args)
}

// NewDefaultDependencies -
func NewDefaultDependencies() (Configurer, Logger, Storer, error) {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, nil, err
	}

	l, err := logger.NewLogger(c)
	if err != nil {
		return nil, nil, nil, err
	}

	d, err := database.NewDatabase(c, l)
	if err != nil {
		return nil, nil, nil, err
	}

	return c, l, d, nil
}

func TestNewService(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	tr := TestRunner{}

	ts, err := NewService(c, l, s, &tr)
	if assert.NoError(t, err, "NewService returns without error") {
		assert.NotNil(t, ts, "Test service is not nil")
	}
}
