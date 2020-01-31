package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/store"
)

// NewDefaultDependencies -
func NewDefaultDependencies() (service.Configurer, service.Logger, service.Storer, error) {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		return nil, nil, nil, err
	}

	l, err := logger.NewLogger(c)
	if err != nil {
		return nil, nil, nil, err
	}

	s, err := store.NewStore(c, l)
	if err != nil {
		return nil, nil, nil, err
	}

	return c, l, s, nil
}
func TestRunner(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	r := NewRunner()

	err = r.Init(c, l, s)
	assert.NoError(t, err, "Init returns without error")
}
