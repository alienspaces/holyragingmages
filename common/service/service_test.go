package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/database"
	"gitlab.com/alienspaces/holyragingmages/common/env"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
)

// TestRunner -
type TestRunner struct {
	InitFunc func() error
	RunFunc  func() error
}

func (r *TestRunner) Init(c Configurer, l Logger, s Storer) error {
	if r.InitFunc == nil {
		l.Printf("InitFunc is nil")
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

func TestNewService(t *testing.T) {

	e, err := env.NewEnv(nil, false)
	if err != nil {
		t.Fatalf("Failed new env >%v<", err)
	}

	l, err := logger.NewLogger(e)
	if err != nil {
		t.Fatalf("Failed new logger >%v<", err)
	}

	d, err := database.NewDatabase(e, l)
	if err != nil {
		t.Fatalf("Failed new database >%v<", err)
	}

	tr := TestRunner{}

	ts, err := NewService(e, l, d, &tr)
	if assert.NoError(t, err, "NewService returns without error") {
		assert.NotNil(t, ts, "Test service is not nil")
	}
}
