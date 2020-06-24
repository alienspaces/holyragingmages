package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/storer"
)

// Model -
type Model struct {
	Config       configurer.Configurer
	Log          logger.Logger
	Store        storer.Storer
	Repositories map[string]repositor.Repositor
	Tx           *sqlx.Tx

	// composable functions
	RepositoriesFunc func(p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error)
}

var _ modeller.Modeller = &Model{}

// NewModel - intended for testing only, maybe move into test files..
func NewModel(c configurer.Configurer, l logger.Logger, s storer.Storer) (m *Model, err error) {

	m = &Model{
		Config: c,
		Log:    l,
		Store:  s,
	}

	return m, nil
}

// Init -
func (m *Model) Init(p preparer.Preparer, tx *sqlx.Tx) (err error) {

	// tx required
	if tx == nil {
		m.Log.Warn("Failed init, tx is required")
		return fmt.Errorf("Failed init, tx is required")
	}

	if m.RepositoriesFunc == nil {
		m.RepositoriesFunc = m.NewRepositories
	}

	// assign database tx for possible custom SQL execution in model functions
	m.Tx = tx

	// repositories
	repositories, err := m.RepositoriesFunc(p, tx)
	if err != nil {
		m.Log.Warn("Failed repositories func >%v<", err)
		return err
	}

	m.Repositories = make(map[string]repositor.Repositor)
	for _, r := range repositories {
		m.Repositories[r.TableName()] = r
	}

	return nil
}

// NewRepositories - default repositor.RepositoriesFunc, override this function for custom repositories
func (m *Model) NewRepositories(p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error) {

	m.Log.Info("** repositor.Repositories **")

	return nil, nil
}
