package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// APIRunner -
type APIRunner struct {
	Store  Storer
	Log    Logger
	Config Configurer
}

// Init -
func (r *APIRunner) Init(c Configurer, l Logger, s Storer) error {

	r.Config = c
	r.Log = l
	r.Store = s

	return nil
}

// Run -
func (r *APIRunner) Run(args map[string]interface{}) error {

	return nil
}

// InitRoutes -
func (r *APIRunner) InitRoutes(h http.Handler) (http.Handler, error) {

	router := httprouter.New()
	router.GET("/", r.IndexGet)

	return router, nil
}

// IndexGet -
func (r *APIRunner) IndexGet(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//
}
