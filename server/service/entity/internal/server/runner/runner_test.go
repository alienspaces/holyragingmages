package runner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/harness"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

const (
	testAccountID string = "5de1cd8d-e136-47b9-82cd-b42b2a0e13eb"
)

func NewTestHarness() (*harness.Testing, error) {

	// harness
	config := harness.DataConfig{
		AccountEntityConfig: []harness.AccountEntityConfig{
			{
				Record: record.AccountEntity{},
				EntityConfig: []harness.EntityConfig{
					{
						Record: record.Entity{},
					},
				},
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

	// Test harness
	th, err := NewTestHarness()
	require.NoError(t, err, "New test data returns without error")

	//  Test dependencies
	c, l, s, err := th.NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	r := NewRunner()

	err = r.Init(c, l, s)
	require.NoError(t, err, "Init returns without error")
}
