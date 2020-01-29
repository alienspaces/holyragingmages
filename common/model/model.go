package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Logger -
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// Model -
type Model struct {
	Log Logger
	Tx  *sqlx.Tx
}

// Init -
func (m *Model) Init(tx *sqlx.Tx) (err error) {

	// tx required
	if tx == nil {
		m.Log.Warn("Failed init, tx is required")
		return fmt.Errorf("Failed init, tx is required")
	}

	return nil
}
