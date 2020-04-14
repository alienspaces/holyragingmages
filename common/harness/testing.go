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
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
)

// CreateDataFunc - callback function that creates test data
type CreateDataFunc func() error

// RemoveDataFunc - callback function that removes test data
type RemoveDataFunc func() error

// Testing -
type Testing struct {
	Config  configurer.Configurer
	Log     logger.Logger
	Store   storer.Storer
	Prepare preparer.Preparer
	Model   modeller.Modeller

	// Configuration
	CommitData bool

	// Modeller function
	ModellerFunc func() (modeller.Modeller, error)

	// Composable functions
	CreateDataFunc CreateDataFunc
	RemoveDataFunc RemoveDataFunc

	// Private
	tx *sqlx.Tx
}

// NewTesting -
func NewTesting() (t *Testing, err error) {

	t = &Testing{}

	return t, nil
}

// Init -
func (t *Testing) Init() (err error) {

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

	// preparer
	t.Prepare, err = prepare.NewPrepare(t.Log)
	if err != nil {
		t.Log.Warn("Failed new preparer >%v<", err)
		return err
	}

	t.Log.Debug("Preparer ready")

	// modeller
	t.Model, err = t.ModellerFunc()
	if err != nil {
		t.Log.Warn("Failed new modeller >%v<", err)
		return err
	}

	t.Log.Debug("Modeller ready")

	return nil
}

// InitTx -
func (t *Testing) InitTx(tx *sqlx.Tx) (err error) {

	t.Log.Debug("Initialising database tx")

	// initialise our own database tx when none is provided
	if tx == nil {
		t.Log.Debug("Starting database tx")

		tx, err = t.Store.GetTx()
		if err != nil {
			t.Log.Warn("Failed getting database tx >%v<", err)
			return err
		}
	}

	err = t.Prepare.Init(tx)
	if err != nil {
		t.Log.Warn("Failed preparer init >%v<", err)
		return err
	}

	err = t.Model.Init(t.Prepare, tx)
	if err != nil {
		t.Log.Warn("Failed modeller init >%v<", err)
		return err
	}

	t.tx = tx

	t.Log.Debug("Database tx initialised")

	return nil
}

// CommitTx -
func (t *Testing) CommitTx() (err error) {
	return t.tx.Commit()
}

// RollbackTx -
func (t *Testing) RollbackTx() (err error) {
	return t.tx.Rollback()
}

// Setup -
func (t *Testing) Setup() (err error) {

	// init
	err = t.InitTx(nil)
	if err != nil {
		t.Log.Warn("Failed init >%v<", err)
		return err
	}

	// data function is expected to create and manage its own store
	if t.CreateDataFunc != nil {
		t.Log.Debug("Creating test data")
		err := t.CreateDataFunc()
		if err != nil {
			t.Log.Warn("Failed creating data >%v<", err)
			return err
		}
	}

	// commit data when configured, otherwise we are leaving
	// it up to tests to explicitly commit or rollback
	if t.CommitData {
		t.Log.Debug("Committing database tx")
		err = t.CommitTx()
		if err != nil {
			t.Log.Warn("Failed comitting data >%v<", err)
			return err
		}
	}

	t.Log.Debug("Setup complete")

	return nil
}

// Teardown -
func (t *Testing) Teardown() (err error) {

	// init
	err = t.InitTx(nil)
	if err != nil {
		t.Log.Warn("Failed init >%v<", err)
		return err
	}

	// data function is expected to create and manage its own store
	if t.RemoveDataFunc != nil {
		t.Log.Debug("Removing test data")
		err := t.RemoveDataFunc()
		if err != nil {
			t.Log.Warn("Failed removing data >%v<", err)
			return err
		}
	}

	// commit data when configured, otherwise we are leaving
	// it up to tests to explicitly commit or rollback
	if t.CommitData {
		t.Log.Debug("Committing database tx")
		err = t.CommitTx()
		if err != nil {
			t.Log.Warn("Failed comitting data >%v<", err)
			return err
		}
	}

	t.Log.Debug("Teardown complete")

	return nil
}
