package model

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/common/model"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/storer"

	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/repository/item"
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

	tr, err := item.NewRepository(m.Log, p, tx)
	if err != nil {
		m.Log.Warn("Failed new item repository >%v<", err)
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

// ItemRepository -
func (m *Model) ItemRepository() *item.Repository {

	r := m.Repositories[item.TableName]
	if r == nil {
		m.Log.Warn("Repository >%s< is nil", item.TableName)
		return nil
	}

	return r.(*item.Repository)
}
