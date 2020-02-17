package test

// NOTE: repository tests are run is the public space so we are
// able to use common setup and teardown tooling for all repositories

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/harness"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/repository/template"
)

func TestCreateRec(t *testing.T) {

	// harness
	harness, err := harness.NewTesting()
	if err != nil {
		t.Fatalf("Failed new test harness >%v<", err)
	}

	err = harness.Setup()
	if err != nil {
		t.Fatalf("Failed test harness setup >%v<", err)
	}

	defer harness.Teardown()

	// repository
	r := harness.TemplateRepository()
	if r == nil {
		t.Fatalf("Repository >%s< is nil", template.RepositoryTableName)
	}

	tests := []struct {
		name string
		rec  func() *record.Template
		err  bool
	}{
		{
			name: "Without ID",
			rec: func() *record.Template {
				return r.NewRecord()
			},
			err: false,
		},
		{
			name: "With ID",
			rec: func() *record.Template {
				rec := r.NewRecord()
				id, _ := uuid.NewRandom()
				rec.ID = id.String()
				return rec
			},
			err: false,
		},
	}

	for _, tc := range tests {

		rec := tc.rec()

		err = r.CreateTestRecord(rec)
		if err != nil {
			t.Fatalf("Failed creating record >%v<", err)
		}
		if tc.err == true {
			assert.Error(t, err, "CreateTestRecord returns error")
			continue
		}
		assert.NoError(t, err, "CreateTestRecord returns without error")
		assert.NotEmpty(t, rec.CreatedAt, "CreateTestRecord returns record with CreatedAt")
	}
}
