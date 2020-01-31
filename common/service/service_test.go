package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
	"gitlab.com/alienspaces/holyragingmages/common/model"
	"gitlab.com/alienspaces/holyragingmages/common/store"
)

// TestRunner - allow Init and Run functions to be defined by tests
type TestRunner struct {
	Runner
	InitFunc func(c Configurer, l Logger, s Storer, m Modeller) error
	RunFunc  func(args map[string]interface{}) error
}

func (rnr *TestRunner) Init(c Configurer, l Logger, s Storer, m Modeller) error {
	if rnr.InitFunc == nil {
		return rnr.Runner.Init(c, l, s, m)
	}
	return rnr.InitFunc(c, l, s, m)
}

func (rnr *TestRunner) Run(args map[string]interface{}) error {
	if rnr.RunFunc == nil {
		return rnr.Runner.Run(args)
	}
	return rnr.RunFunc(args)
}

// NewDefaultDependencies -
func NewDefaultDependencies() (Configurer, Logger, Storer, Modeller, error) {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	l, err := logger.NewLogger(c)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	s, err := store.NewStore(c, l)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	m, err := model.NewModel()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return c, l, s, m, nil
}

func TestNewService(t *testing.T) {

	c, l, s, m, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	tr := TestRunner{}

	ts, err := NewService(c, l, s, m, &tr)
	if assert.NoError(t, err, "NewService returns without error") {
		assert.NotNil(t, ts, "Test service is not nil")
	}
}
