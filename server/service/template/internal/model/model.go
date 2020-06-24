package model

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/common/model"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/storer"

	"gitlab.com/alienspaces/holyragingmages/server/service/template/internal/repository/template"
)

// Model -
type Model struct {
	model.Model
}

// NewModel -
func NewModel(c configurer.Configurer, l logger.Logger, s storer.Storer) (*Model, error) {

	m := &Model{
		model.Model{
			Config: c,
			Log:    l,
			Store:  s,
		},
	}

	m.RepositoriesFunc = m.NewRepositories

	return m, nil
}

// NewRepositories - Custom repositories for this model
func (m *Model) NewRepositories(p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error) {

	repositoryList := []repositor.Repositor{}

	tr, err := template.NewRepository(m.Log, p, tx)
	if err != nil {
		m.Log.Warn("Failed new template repository >%v<", err)
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

// TemplateRepository -
func (m *Model) TemplateRepository() *template.Repository {

	r := m.Repositories[template.TableName]
	if r == nil {
		m.Log.Warn("Repository >%s< is nil", template.TableName)
		return nil
	}

	return r.(*template.Repository)
}
