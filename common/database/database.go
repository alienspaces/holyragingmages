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

const (
	// DBPostgres -
	DBPostgres string = "postgres"
)

// Store -
type Store struct {
	Logger   Logger
	Config   Config
	Database string
}

// NewDatabase - Establishes a new database connection
func NewDatabase(c Config, l Logger) (*Store, error) {

	dt := c.Get("APP_DATABASE")
	if dt == "" {
		l.Printf("Defaulting to postgres")
		dt = DBPostgres
	}

	s := Store{
		Logger:   l,
		Config:   c,
		Database: dt,
	}

	return &s, nil
}

// GetDb -
func (s *Store) GetDb() (*sqlx.DB, error) {
	if s.Database == DBPostgres {
		s.Logger.Printf("Connecting to postgres")
		return newPostgresDB(s.Logger, s.Config)
	}
	return nil, fmt.Errorf("Unsupported database")
}

// GetTx -
func (s *Store) GetTx() (*sqlx.Tx, error) {
	db, err := s.GetDb()
	if err != nil {
		return nil, err
	}
	return db.Beginx()
}
