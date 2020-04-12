package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
)

// Handle - custom service handle
type Handle func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller)

// Service -
type Service struct {
	Config  configurer.Configurer
	Log     logger.Logger
	Store   storer.Storer
	Prepare preparer.Preparer
	Runner  runnable.Runnable
}

// NewService -
func NewService(c configurer.Configurer, l logger.Logger, s storer.Storer, p preparer.Preparer, r runnable.Runnable) (*Service, error) {

	svc := Service{
		Config:  c,
		Log:     l,
		Store:   s,
		Prepare: p,
		Runner:  r,
	}

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

	// TODO: alerting, retries
	return svc.Runner.Init(svc.Config, svc.Log, svc.Store, svc.Prepare)
}

// Run -
func (svc *Service) Run(args map[string]interface{}) error {

	// TODO:
	// - alerting on errors
	// - retries on start up
	// - reload  on config changes
	return svc.Runner.Run(args)
}
