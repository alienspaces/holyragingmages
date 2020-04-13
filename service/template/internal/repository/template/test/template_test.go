package test

// NOTE: repository tests are run is the public space so we are
// able to use common setup and teardown tooling for all repositories

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/harness"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
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
	r := h.Model.(*model.Model).TemplateRepository()
	if r == nil {
		t.Fatalf("Repository >%s< is nil", template.TableName)
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

		t.Logf("Run test >%s<", tc.name)

		rec := tc.rec()

		err = r.CreateOne(rec)
		if err != nil {
			t.Fatalf("Failed creating record >%v<", err)
		}
		if tc.err == true {
			require.Error(t, err, "CreateOne returns error")
			continue
		}
		require.NoError(t, err, "CreateOne returns without error")
		require.NotEmpty(t, rec.CreatedAt, "CreateOne returns record with CreatedAt")
	}
}

func TestGetOne(t *testing.T) {

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
	r := h.Model.(*model.Model).TemplateRepository()
	if r == nil {
		t.Fatalf("Repository >%s< is nil", template.TableName)
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

		t.Logf("Run test >%s<", tc.name)

		rec, err := r.GetOne(tc.id(), false)
		if tc.err == true {
			require.Error(t, err, "GetOne returns error")
			continue
		}
		require.NoError(t, err, "GetOne returns without error")
		require.NotNil(t, rec, "GetOne returns record")
		require.NotEmpty(t, rec.ID, "Record ID is not empty")
	}
}

func TestUpdateOne(t *testing.T) {

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
	r := h.Model.(*model.Model).TemplateRepository()
	if r == nil {
		t.Fatalf("Repository >%s< is nil", template.TableName)
	}

	tests := []struct {
		name string
		rec  func() *record.Template
		err  bool
	}{
		{
			name: "With ID",
			rec: func() *record.Template {
				return h.Data.TemplateRecs[0]
			},
			err: false,
		},
		{
			name: "Without ID",
			rec: func() *record.Template {
				rec := h.Data.TemplateRecs[0]
				rec.ID = ""
				return rec
			},
			err: true,
		},
	}

	for _, tc := range tests {

		t.Logf("Run test >%s<", tc.name)

		rec := tc.rec()

		err := r.UpdateOne(rec)
		if tc.err == true {
			require.Error(t, err, "UpdateOne returns error")
			continue
		}
		require.NoError(t, err, "UpdateOne returns without error")
		require.NotEmpty(t, rec.UpdatedAt, "UpdateOne returns record with UpdatedAt")
	}
}

func TestDeleteOne(t *testing.T) {

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
	r := h.Model.(*model.Model).TemplateRepository()
	if r == nil {
		t.Fatalf("Repository >%s< is nil", template.TableName)
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

		t.Logf("Run test >%s<", tc.name)

		err := r.DeleteOne(tc.id())
		if tc.err == true {
			require.Error(t, err, "DeleteOne returns error")
			continue
		}
		require.NoError(t, err, "DeleteOne returns without error")

		rec, err := r.GetOne(tc.id(), false)
		require.Error(t, err, "GetOne returns error")
		require.Nil(t, rec, "GetOne does not return record")
	}
}
