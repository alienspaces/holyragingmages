package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/common/database"
	"gitlab.com/alienspaces/holyragingmages/common/env"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
)

func TestRunner(t *testing.T) {

	e, err := env.NewEnv(nil, false)
	if err != nil {
		t.Fatalf("Failed new env >%v<", err)
	}

	l, err := logger.NewLogger(e)
	if err != nil {
		t.Fatalf("Failed new env >%v<", err)
	}

	d, err := database.NewDatabase(e, l)
	if err != nil {
		t.Fatalf("Failed new env >%v<", err)
	}

	r := Runner{}

	err = r.Init(e, l, d)
	assert.NoError(t, err, "Init returns without error")
}
