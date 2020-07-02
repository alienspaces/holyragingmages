package model

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/core/model"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/storer"

	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/repository/spell"
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

	tr, err := spell.NewRepository(m.Log, p, tx)
	if err != nil {
		m.Log.Warn("Failed new spell repository >%v<", err)
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

// SpellRepository -
func (m *Model) SpellRepository() *spell.Repository {

	r := m.Repositories[spell.TableName]
	if r == nil {
		m.Log.Warn("Repository >%s< is nil", spell.TableName)
		return nil
	}

	return r.(*spell.Repository)
}
