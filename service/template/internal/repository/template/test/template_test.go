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

func TestCreateOne(t *testing.T) {

	// harness
	h, err := harness.NewTesting()
	if err != nil {
		t.Fatalf("Failed new test harness >%v<", err)
	}

	err = h.Setup()
	if err != nil {
		t.Fatalf("Failed test harness setup >%v<", err)
	}

	defer h.Teardown()

	// repository
	r := h.TemplateRepository()
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

		err = r.CreateOne(rec)
		if err != nil {
			t.Fatalf("Failed creating record >%v<", err)
		}
		if tc.err == true {
			assert.Error(t, err, "CreateOne returns error")
			continue
		}
		assert.NoError(t, err, "CreateOne returns without error")
		assert.NotEmpty(t, rec.CreatedAt, "CreateOne returns record with CreatedAt")
	}
}

func TestGetRec(t *testing.T) {

	// harness
	h, err := harness.NewTesting()
	if err != nil {
		t.Fatalf("Failed new test harness >%v<", err)
	}

	err = h.Setup()
	if err != nil {
		t.Fatalf("Failed test harness setup >%v<", err)
	}

	defer h.Teardown()

	// repository
	r := h.TemplateRepository()
	if r == nil {
		t.Fatalf("Repository >%s< is nil", template.RepositoryTableName)
	}

	tests := []struct {
		name string
		id   func() string
		err  bool
	}{
		{
			name: "With ID",
			id: func() string {
				return h.Data.TemplateRecs[0].ID
			},
			err: false,
		},
		{
			name: "Without ID",
			id: func() string {
				return ""
			},
			err: true,
		},
	}

	for _, tc := range tests {

		rec, err := r.GetOne(tc.id(), false)
		if err != nil {
			t.Fatalf("Failed getting record >%v<", err)
		}
		if tc.err == true {
			assert.Error(t, err, "GetOne returns error")
			continue
		}
		assert.NoError(t, err, "GetOne returns without error")
		assert.NotEmpty(t, rec, "GetOne returns record")
	}
}
