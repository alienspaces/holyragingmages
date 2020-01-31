package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
	"gitlab.com/alienspaces/holyragingmages/common/store"
)

func TestRunner(t *testing.T) {

	c, err := config.NewConfig(nil, false)
	if err != nil {
		t.Fatalf("Failed new config >%v<", err)
	}

	l, err := logger.NewLogger(c)
	if err != nil {
		t.Fatalf("Failed new env >%v<", err)
	}

	s, err := store.NewStore(c, l)
	if err != nil {
		t.Fatalf("Failed new env >%v<", err)
	}

	r := Runner{}

	err = r.Init(c, l, s)
	assert.NoError(t, err, "Init returns without error")
}
