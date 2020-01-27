package service

import (
	"github.com/jmoiron/sqlx"
)

// Storer -
type Storer interface {
	GetDb() (*sqlx.DB, error)
	GetTx() (*sqlx.Tx, error)
}

// Configurer -
type Configurer interface {
	Get(key string) string
}

// Logger -
type Logger interface {
	Printf(format string, v ...interface{})
}

// Runnerer -
type Runnerer interface {
	Init(c Configurer, l Logger, s Storer) error
	Run(args map[string]interface{}) error
}

// Service -
type Service struct {
	Store  Storer
	Log    Logger
	Config Configurer
	Runner Runnerer
}

// NewService -
func NewService(c Configurer, l Logger, s Storer, r Runnerer) (*Service, error) {

	svc := Service{
		Config: c,
		Log:    l,
		Store:  s,
		Runner: r,
	}

	err := svc.init()
	if err != nil {
		return nil, err
	}

	return &svc, nil
}

// Init -
func (svc *Service) init() error {

	// TODO: exception handling and alerting
	return svc.Runner.Init(svc.Config, svc.Log, svc.Store)
}

// Run -
func (svc *Service) Run(args map[string]interface{}) error {

	// TODO: exception handling and alerting
	return svc.Runner.Run(args)
}
