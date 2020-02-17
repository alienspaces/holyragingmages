package harness

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/harness"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/repository/template"
)

// Testing -
type Testing struct {
	harness.Testing
}

// NewTesting -
func NewTesting() (*Testing, error) {

	// harness
	h := Testing{}

	h.RepositoriesFunc = newRepositories

	return &h, nil
}

// TemplateRepository -
func (t *Testing) TemplateRepository() *template.Repository {

	r := t.Repository(template.RepositoryTableName)
	if r == nil {
		return nil
	}

	return r.(*template.Repository)
}

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
