package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Configurer -
type Configurer interface {
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
	Config   Configurer
	Database string
}

// NewStore -
func NewStore(c Configurer, l Logger) (*Store, error) {

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

// Init - initialize store
func (s *Store) Init() error {

	return nil
}

// GetDb -
func (s *Store) GetDb() (*sqlx.DB, error) {
	if s.Database == DBPostgres {
		s.Logger.Printf("Connecting to postgres")
		return newPostgresDB(s.Config, s.Logger)
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
