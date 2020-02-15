package testharness

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

// TestHarness -
type TestHarness struct {
	Store        storer.Storer
	Log          logger.Logger
	Config       configurer.Configurer
	Repositories map[string]repositor.Repositor

	// composable functions
	RepositoriesFunc func(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error)

	// private
	tx *sqlx.Tx
}

// NewTestHarness -
func NewTestHarness(r RepositoriesFunc) (t *TestHarness, err error) {

	t = &TestHarness{
		RepositoriesFunc: r,
	}

	t.Config, err = config.NewConfig(nil, false)
	if err != nil {
		return nil, err
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
			return nil, err
		}
	}

	t.Log, err = log.NewLogger(t.Config)
	if err != nil {
		return nil, err
	}

	t.Store, err = store.NewStore(t.Config, t.Log)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Setup -
func (t *TestHarness) Setup() error {

	err := t.Store.Init()
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
func (t *TestHarness) Teardown() error {

	err := t.tx.Rollback()
	if err != nil {
		t.Log.Warn("Failed database tx rollback >%v<", err)
		return err
	}

	return nil
}

// Tx -
func (t *TestHarness) Tx() *sqlx.Tx {
	return t.tx
}

// Repository -
func (t *TestHarness) Repository(name string) interface{} {
	return t.Repositories[name]
}
