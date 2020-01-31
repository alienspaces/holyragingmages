package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/alienspaces/holyragingmages/common/config"
)

// Configurer -
type Configurer interface {
	Get(key string) string
	Set(key string, value string)
	Add(item config.Item) (err error)
}

// Logger -
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

const (
	// DBPostgres -
	DBPostgres string = "postgres"
)

// Store -
type Store struct {
	Log        Logger
	Config     Configurer
	Database   string
	Connection *sqlx.DB
}

// NewStore -
func NewStore(c Configurer, l Logger) (*Store, error) {

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

	configVars := []string{
		// database
		"APP_DB_HOST",
		"APP_DB_PORT",
		"APP_DB_NAME",
		"APP_DB_USER",
		"APP_DB_PASSWORD",
	}
	for _, key := range configVars {
		err := s.Config.Add(config.Item{
			Key:      key,
			Required: true,
		})
		if err != nil {
			s.Log.Warn("Failed adding config item >%v<", err)
			return err
		}
	}

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
		return newPostgresDB(s.Config, s.Log)
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
