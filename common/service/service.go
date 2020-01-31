package service

import (
	"github.com/jmoiron/sqlx"
)

// Configurer -
type Configurer interface {
	Get(key string) string
}

// Logger -
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// Storer -
type Storer interface {
	GetDb() (*sqlx.DB, error)
	GetTx() (*sqlx.Tx, error)
}

// Runnable -
type Runnable interface {
	Init(c Configurer, l Logger, s Storer) error
	Run(args map[string]interface{}) error
}

// Modeller -
type Modeller interface {
	Init(tx *sqlx.Tx) (err error)
}

// Service -
type Service struct {
	Store  Storer
	Log    Logger
	Config Configurer
	Runner Runnable
}

// NewService -
func NewService(c Configurer, l Logger, s Storer, r Runnable) (*Service, error) {

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
