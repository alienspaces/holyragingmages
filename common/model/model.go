package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Configurer -
type Configurer interface {
	Get(key string) string
	Set(key string, value string)
	Add(key string, required bool) (err error)
}

// Logger -
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// Storer -
type Storer interface {
	GetDb() (*sqlx.DB, error)
	GetTx() (*sqlx.Tx, error)
}

// Model -
type Model struct {
	Config Configurer
	Log    Logger
	Store  Storer
	Tx     *sqlx.Tx
}

// NewModel - intended for testing only, maybe move into test files..
func NewModel(c Configurer, l Logger, s Storer) (m *Model, err error) {

	m = &Model{
		Config: c,
		Log:    l,
		Store:  s,
	}

	return m, nil
}

// Init -
func (m *Model) Init(tx *sqlx.Tx) (err error) {

	// tx required
	if tx == nil {
		m.Log.Warn("Failed init, tx is required")
		return fmt.Errorf("Failed init, tx is required")
	}

	m.Tx = tx

	return nil
}
