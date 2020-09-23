package model

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/core/model"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/storer"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/repository/account"
)

// Model -
type Model struct {
	model.Model
	// Allows auth token verification to be mocked for testing
	VerifyAuthTokenFunc func(provider, token string) (*VerifiedData, error)
}

// NewModel -
func NewModel(c configurer.Configurer, l logger.Logger, s storer.Storer) (*Model, error) {

	m := &Model{
		Model: model.Model{
			Config: c,
			Log:    l,
			Store:  s,
		},
		VerifyAuthTokenFunc: nil,
	}

	m.VerifyAuthTokenFunc = m.verifyAuthToken
	m.RepositoriesFunc = m.NewRepositories

	return m, nil
}

// NewRepositories - Custom repositories for this model
func (m *Model) NewRepositories(p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error) {

	repositoryList := []repositor.Repositor{}

	tr, err := account.NewRepository(m.Log, p, tx)
	if err != nil {
		m.Log.Warn("Failed new account repository >%v<", err)
		return nil, err
	}

	repositoryList = append(repositoryList, tr)

	return repositoryList, nil
}

// AccountRepository -
func (m *Model) AccountRepository() *account.Repository {

	r := m.Repositories[account.TableName]
	if r == nil {
		m.Log.Warn("Repository >%s< is nil", account.TableName)
		return nil
	}

	return r.(*account.Repository)
}
