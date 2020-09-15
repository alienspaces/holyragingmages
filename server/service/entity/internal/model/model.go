package model

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/core/model"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/storer"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/repository/entity"
)

const (
	startingAttributePoints int64 = 32
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

	tr, err := entity.NewRepository(m.Log, p, tx)
	if err != nil {
		m.Log.Warn("Failed new mage repository >%v<", err)
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

// EntityRepository -
func (m *Model) EntityRepository() *entity.Repository {

	r := m.Repositories[entity.TableName]
	if r == nil {
		m.Log.Warn("Repository >%s< is nil", entity.TableName)
		return nil
	}

	return r.(*entity.Repository)
}
