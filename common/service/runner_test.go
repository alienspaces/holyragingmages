package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunnerInit(t *testing.T) {

	c, l, s, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed new default dependencies >%v<", err)
	}

	tr := TestRunner{}

	err = tr.Init(c, l, s)
	assert.NoError(t, err, "Runner Init returns without error")
}
