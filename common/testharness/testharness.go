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
	Repositories []repositor.Repositor

	// composable functions
	RepositoriesFunc func(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error)
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
func (t *TestHarness) Setup(tx *sqlx.Tx) error {

	p, err := prepare.NewPrepare(t.Log, tx)
	if err != nil {
		t.Log.Warn("Failed new preparer >%v<", err)
		return err
	}

	t.Log.Info("Preparer ready >%v<", p)

	repositories, err := t.RepositoriesFunc(t.Log, p, tx)
	if err != nil {
		t.Log.Warn("Failed repositories >%v<", err)
		return err
	}

	t.Repositories = repositories

	return nil
}

// Teardown -
func (t *TestHarness) Teardown(tx *sqlx.Tx) error {

	return nil
}
