package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/logger"
)

const (
	// DBPostgres -
	DBPostgres string = "postgres"
)

// Store -
type Store struct {
	Log        logger.Logger
	Config     configurer.Configurer
	Database   string
	Connection *sqlx.DB
}

// NewStore -
func NewStore(c configurer.Configurer, l logger.Logger) (*Store, error) {

	dt := c.Get("APP_DATABASE")
	if dt == "" {
		l.Info("Defaulting to postgres")
		dt = DBPostgres
	}

	s := Store{
		Log:      l,
		Config:   c,
		Database: dt,
	}

	return &s, nil
}

// Init - initialize store
func (s *Store) Init() error {

	c, err := s.GetDb()
	if err != nil {
		s.Log.Warn("Failed getting DB connection >%v<", err)
		return err
	}

	s.Connection = c

	return nil
}

// GetDb -
func (s *Store) GetDb() (*sqlx.DB, error) {
	if s.Database == DBPostgres {
		s.Log.Info("Connecting to postgres")
		return getPostgresDB(s.Config, s.Log)
	}
	return nil, fmt.Errorf("Unsupported database")
}

// GetTx -
func (s *Store) GetTx() (*sqlx.Tx, error) {

	if s.Connection == nil {
		s.Log.Warn("Not connected")
		return nil, fmt.Errorf("Not connection")
	}

	return s.Connection.Beginx()
}
