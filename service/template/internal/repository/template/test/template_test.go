package test

// NOTE: repository tests are run is the public space so we are
// able to use common setup and teardown tooling for all repositories

import (
	"testing"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/harness"
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

	rec := r.NewRecord()

	t.Logf("Have record >%#v<", rec)

	// create test record
	err = r.CreateTestRecord(rec)
	if err != nil {
		t.Fatalf("Failed creating record >%v<", err)
	}
}
