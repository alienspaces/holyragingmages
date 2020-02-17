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

// Testing -
type Testing struct {
	Store        storer.Storer
	Log          logger.Logger
	Config       configurer.Configurer
	Repositories map[string]repositor.Repositor

	// composable functions
	RepositoriesFunc RepositoriesFunc

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

	t.Log.Info("Preparer ready >%v<", p)

	repositories, err := t.RepositoriesFunc(t.Log, p, t.tx)
	if err != nil {
		t.Log.Warn("Failed repositories >%v<", err)
		return err
	}

	t.Repositories = make(map[string]repositor.Repositor)
	for _, r := range repositories {
		err := r.Init(p, t.tx)
		if err != nil {
			t.Log.Warn("Failed initialising repository >%s< >%v<", r.TableName(), err)
			return err
		}
		t.Repositories[r.TableName()] = r
	}

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
