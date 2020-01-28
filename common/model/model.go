package model

import (
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
