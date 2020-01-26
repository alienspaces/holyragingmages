package service

import (
	"fmt"
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

	r.Log.Printf("** Initialised **")

	return nil
}

// Run -
func (r *APIRunner) Run(args map[string]interface{}) error {

	r.Log.Printf("** Run **")

	handler, err := r.Handler()
	if err != nil {
		r.Log.Printf("Failed handler >%v<", err)
		return err
	}

	return http.ListenAndServe(":8080", handler)
}

// Handler -
func (r *APIRunner) Handler() (http.Handler, error) {

	r.Log.Printf("** Handler **")

	router, err := r.Router()
	if err != nil {
		r.Log.Printf("Failed router >%v<", err)
		return nil, err
	}

	return router, nil
}

// Router -
func (r *APIRunner) Router() (http.Handler, error) {

	r.Log.Printf("** Router **")

	return r.DefaultRouter()
}

// DefaultRouter -
func (r *APIRunner) DefaultRouter() (http.Handler, error) {

	r.Log.Printf("** DefaultRouter **")

	router := httprouter.New()
	router.GET("/", r.DefaultHandler)

	return router, nil
}

// DefaultHandler -
func (r *APIRunner) DefaultHandler(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	r.Log.Printf("** DefaultHandler **")

	fmt.Fprint(resp, "Default\n")
}
