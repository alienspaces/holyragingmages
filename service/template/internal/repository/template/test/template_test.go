package test

// NOTE: repository tests are run is the public space so we are
// able to use common setup and teardown tooling for all repositories

import (
	"testing"

	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/testharness"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/repository/template"
)

// NewRepositories - Custom repositories for this model
func newRepositories(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error) {

	repositoryList := []repositor.Repositor{}

	tr, err := template.NewRepository(l, p, tx)
	if err != nil {
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

func TestCreateRec(t *testing.T) {

	// harness
	harness, err := testharness.NewTestHarness(newRepositories)
	if err != nil {
		t.Fatalf("Failed new test harness >%v<", err)
	}

	err = harness.Setup()
	if err != nil {
		t.Fatalf("Failed test harness setup >%v<", err)
	}

	defer harness.Teardown()

	// repository
	r := harness.Repository(template.RepositoryTableName)
	if r == nil {
		t.Fatalf("Repository >%s< is nil", template.RepositoryTableName)
	}

	rp := r.(*template.Repository)
	rec := rp.NewRecord()

	t.Logf("Have record >%#v<", rec)

	// create test record
	err = rp.CreateTestRecord(rec)
	if err != nil {
		t.Fatalf("Failed creating record >%v<", err)
	}
}
