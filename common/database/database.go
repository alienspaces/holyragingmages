package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Config -
type Config interface {
	Get(key string) string
}

// Logger -
type Logger interface {
	Printf(format string, v ...interface{})
}

// NewDatabase - Establishes a new database connection
func NewDatabase(l Logger, c Config) (*sqlx.DB, error) {

	dt := c.Get("APP_DATABASE")
	if dt == "" {
		l.Printf("Defaulting to postgres")
		dt = "postgres"
	}

	if dt == "postgres" {
		l.Printf("Using postgres")
		return newPostgresDB(l, c)
	}

	return nil, fmt.Errorf("Unsuported database type %s", dt)
}
