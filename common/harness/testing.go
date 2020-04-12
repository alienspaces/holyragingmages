package harness

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/config"
	"gitlab.com/alienspaces/holyragingmages/common/log"
	"gitlab.com/alienspaces/holyragingmages/common/prepare"
	"gitlab.com/alienspaces/holyragingmages/common/store"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/repositor"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
)

// RepositoriesFunc - callback function to return repositories
type RepositoriesFunc func(l logger.Logger, p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error)

// CreateDataFunc - callback function that creates test data
type CreateDataFunc func() error

// RemoveDataFunc - callback function that removes test data
type RemoveDataFunc func() error

// DataConfig - configuration for creating test data
type DataConfig struct{}

// Data - contains test data
type Data struct{}

// Testing -
type Testing struct {
	Config  configurer.Configurer
	Log     logger.Logger
	Store   storer.Storer
	Prepare preparer.Preparer
	Model   modeller.Modeller

	// Modeller function
	ModellerFunc func() (modeller.Modeller, error)

	// Composable functions
	CreateDataFunc CreateDataFunc
	RemoveDataFunc RemoveDataFunc

	// DataConfig
	DataConfig DataConfig

	// Data
	Data Data

	// Private
	tx *sqlx.Tx
}

// NewTesting -
func NewTesting(m modeller.Modeller) (t *Testing, err error) {

	t = &Testing{
		Model: m,
	}

	return t, nil
}

// Setup -
func (t *Testing) Setup() (err error) {

	// configurer
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

	// logger
	t.Log, err = log.NewLogger(t.Config)
	if err != nil {
		return err
	}

	// storer
	t.Store, err = store.NewStore(t.Config, t.Log)
	if err != nil {
		return err
	}

	err = t.Store.Init()
	if err != nil {
		t.Log.Warn("Failed storer init >%v<", err)
		return err
	}

	tx, err := t.Store.GetTx()
	if err != nil {
		t.Log.Warn("Failed getting database tx >%v<", err)
		return err
	}

	t.tx = tx

	// preparer
	t.Prepare, err = prepare.NewPrepare(t.Log)
	if err != nil {
		t.Log.Warn("Failed new preparer >%v<", err)
		return err
	}

	err = t.Prepare.Init(tx)
	if err != nil {
		t.Log.Warn("Failed preparer init >%v<", err)
		return err
	}

	t.Log.Info("Preparer ready")

	// modeller
	t.Model, err = t.ModellerFunc()
	if err != nil {
		t.Log.Warn("Failed new modeller >%v<", err)
		return err
	}

	err = t.Model.Init(t.Prepare, tx)
	if err != nil {
		t.Log.Warn("Failed modeller init >%v<", err)
		return err
	}

	t.Log.Info("Repositories ready")

	// data function is expected to create and manage its own store
	if t.CreateDataFunc != nil {
		t.Log.Info("Creating test data")
		err := t.CreateDataFunc()
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
