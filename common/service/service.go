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

// Runner -
type Runner interface {
	Init(c Configurer, l Logger, s Storer) error
	Run(args map[string]interface{}) error
}

// Service -
type Service struct {
	Storer     Storer
	Logger     Logger
	Configurer Configurer
	Runner     Runner
}

// NewService -
func NewService(c Configurer, l Logger, s Storer, r Runner) (*Service, error) {

	svc := Service{
		Configurer: c,
		Logger:     l,
		Storer:     s,
		Runner:     r,
	}

	err := svc.init()
	if err != nil {
		return nil, err
	}

	return &svc, nil
}

// Init -
func (svc *Service) init() error {
	return svc.Runner.Init(svc.Configurer, svc.Logger, svc.Storer)
}

// Run -
func (svc *Service) Run(args map[string]interface{}) error {
	return svc.Runner.Run(args)
}
