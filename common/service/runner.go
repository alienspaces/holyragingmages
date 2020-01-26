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

	r.Log.Printf("** Running **")

	handler, err := r.Handler()
	if err != nil {
		r.Log.Printf("Failed handler >%v<", err)
		return err
	}

	return http.ListenAndServe(":8080", handler)
}

// Handler -
func (r *APIRunner) Handler() (http.Handler, error) {

	router := httprouter.New()
	router.GET("/", r.IndexGet)

	return router, nil
}

// IndexGet -
func (r *APIRunner) IndexGet(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//
	r.Log.Printf("IndexGet - Dum dum")

	fmt.Fprint(resp, "Welcome!\n")
}
