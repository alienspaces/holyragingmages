package service

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"

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

// Storer -
type Storer interface {
	Init() error
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

// Handle - custom service handle
type Handle func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m Modeller)

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

	// TODO: exception handling and alerting
	err := svc.Init()
	if err != nil {
		return nil, err
	}

	return &svc, nil
}

// Init -
func (svc *Service) Init() error {

	err := svc.Store.Init()
	if err != nil {
		return err
	}

	return svc.Runner.Init(svc.Config, svc.Log, svc.Store)
}

// Run -
func (svc *Service) Run(args map[string]interface{}) error {

	return svc.Runner.Run(args)
}
