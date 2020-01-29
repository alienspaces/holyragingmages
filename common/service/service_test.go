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
	InitFunc func() error
	RunFunc  func() error
}

func (r *TestRunner) Init(c Configurer, l Logger, s Storer) error {
	if r.InitFunc == nil {
		l.Info("InitFunc is nil")
		return nil
	}
	return r.InitFunc()
}

func (r *TestRunner) Run(args map[string]interface{}) error {
	if r.RunFunc == nil {
		return nil
	}
	return r.RunFunc()
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
