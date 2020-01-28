package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Logger -
type Logger interface {
	Printf(format string, v ...interface{})
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
		m.Log.Printf("Failed init, tx is required")
		return fmt.Errorf("Failed init, tx is required")
	}

	return nil
}
