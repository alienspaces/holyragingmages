package service

import (
	"github.com/jmoiron/sqlx"
)

// Database -
type Database interface {
	GetConnection() *sqlx.DB
	CloseConnection()
	BeginTrx() *sqlx.Tx
}

// Config -
type Config interface {
	Get(key string) string
}

// Logger -
type Logger interface {
	Printf(format string, v ...interface{})
}

// Container -
type Container struct {
	Database Database
	Logger   Logger
	Config   Config
}

// NewContainer -
func NewContainer(c Config, l Logger, d Database) (*Container, error) {

	container := Container{
		Config:   c,
		Logger:   l,
		Database: d,
	}

	err := container.Init()
	if err != nil {
		return nil, err
	}

	return &container, nil
}

// Init -
func (c *Container) Init() error {

	return nil
}
