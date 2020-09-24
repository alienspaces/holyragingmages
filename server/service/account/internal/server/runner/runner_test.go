package runner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/harness"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

func NewTestHarness() (*harness.Testing, error) {

	// harness
	config := harness.DataConfig{
		AccountConfig: []harness.AccountConfig{
			{
				Record: record.Account{},
			},
		},
	}

	h, err := harness.NewTesting(config)
	if err != nil {
		return nil, err
	}

	// harness commit data
	h.CommitData = true

	return h, nil
}

func TestNewRunner(t *testing.T) {

	// test harness
	th, err := NewTestHarness()
	require.NoError(t, err, "New test data returns without error")

	// test dependencies
	c, l, s, err := th.NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	r := NewRunner()

	err = r.Init(c, l, s)
	require.NoError(t, err, "Init returns without error")
}
