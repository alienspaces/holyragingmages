package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/database"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
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

	d, err := database.NewDatabase(c, l)
	if err != nil {
		t.Fatalf("Failed new env >%v<", err)
	}

	r := Runner{}

	err = r.Init(c, l, d)
	assert.NoError(t, err, "Init returns without error")
}
