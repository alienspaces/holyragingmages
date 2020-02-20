package harness

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
	"gitlab.com/alienspaces/holyragingmages/common/prepare"
	"gitlab.com/alienspaces/holyragingmages/common/store"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
)

// RepositoriesFunc - callback function to return repositories
type RepositoriesFunc func(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error)

// DataFunc - callback function that creates test data
type DataFunc func() error

// Data - contains test data
type Data struct{}

// Testing -
type Testing struct {
	Store        storer.Storer
	Log          logger.Logger
	Config       configurer.Configurer
	Repositories map[string]repositor.Repositor

	// Data
	Data Data

	// composable functions
	RepositoriesFunc RepositoriesFunc
	DataFunc         DataFunc

	// private
	tx *sqlx.Tx
}

// NewTesting -
func NewTesting(r RepositoriesFunc) (t *Testing, err error) {

	t = &Testing{
		RepositoriesFunc: r,
	}

	return t, nil
}

// Setup -
func (t *Testing) Setup() (err error) {

	t.Config, err = config.NewConfig(nil, false)
	if err != nil {
		return err
	}

	configVars := []string{
		// logger
		"APP_LOG_LEVEL",
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
	}
	for _, key := range configVars {
		err = t.Config.Add(key, true)
		if err != nil {
			return err
		}
	}

	t.Log, err = log.NewLogger(t.Config)
	if err != nil {
		return err
	}

	t.Store, err = store.NewStore(t.Config, t.Log)
	if err != nil {
		return err
	}

	err = t.Store.Init()
	if err != nil {
		t.Log.Warn("Failed store init >%v<", err)
		return err
	}

	tx, err := t.Store.GetTx()
	if err != nil {
		t.Log.Warn("Failed getting database tx >%v<", err)
		return err
	}

	t.tx = tx

	p, err := prepare.NewPrepare(t.Log, t.tx)
	if err != nil {
		t.Log.Warn("Failed new preparer >%v<", err)
		return err
	}

	t.Log.Info("Preparer ready")

	// repositories function is expected to create and return a list of repositors
	repositories, err := t.RepositoriesFunc(t.Log, p, t.tx)
	if err != nil {
		t.Log.Warn("Failed repositories function >%v<", err)
		return err
	}

	t.Repositories = make(map[string]repositor.Repositor)
	for _, r := range repositories {
		t.Repositories[r.TableName()] = r
	}

	t.Log.Info("Repositories ready")

	// data function is expected to create and manage its own store
	if t.DataFunc != nil {
		t.Log.Info("Creating test data")
		err := t.DataFunc()
		if err != nil {
			t.Log.Warn("Failed data function >%v<", err)
			return err
		}
	}

	t.Log.Info("Data ready")

	return nil
}

// Teardown -
func (t *Testing) Teardown() error {

	err := t.tx.Rollback()
	if err != nil {
		t.Log.Warn("Failed database tx rollback >%v<", err)
		return err
	}

	return nil
}

// Tx -
func (t *Testing) Tx() *sqlx.Tx {
	return t.tx
}

// Repository -
func (t *Testing) Repository(name string) interface{} {
	return t.Repositories[name]
}
