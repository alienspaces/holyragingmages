package harness

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/harness"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/repository/template"
)

// Testing -
type Testing struct {
	harness.Testing
	Data *Data
}

// Data -
type Data struct {
	TemplateRecs []*record.Template
}

// NewTesting -
func NewTesting() (*Testing, error) {

	// harness
	h := Testing{}

	h.RepositoriesFunc = h.CreateRepositories
	h.DataFunc = h.CreateData
	h.Data = &Data{}

	return &h, nil
}

// CreateRepositories - Custom repositories
func (t *Testing) CreateRepositories(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error) {

	repositoryList := []repositor.Repositor{}

	tr, err := template.NewRepository(l, p, tx)
	if err != nil {
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

// TemplateRepository -
func (t *Testing) TemplateRepository() *template.Repository {

	r := t.Repository(template.RepositoryTableName)
	if r == nil {
		return nil
	}

	return r.(*template.Repository)
}

// CreateData - Custom data
func (t *Testing) CreateData() error {

	tr := t.TemplateRepository()
	rec := tr.NewRecord()

	t.Log.Warn("Test record >%#v<", rec)

	err := tr.CreateTestRecord(rec)
	if err != nil {
		t.Log.Warn("Failed creating test template record >%v<", err)
		return err
	}

	t.Data.TemplateRecs = append(t.Data.TemplateRecs, rec)

	return nil
}
