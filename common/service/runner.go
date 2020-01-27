package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Runner - implements the runnerer interface
type Runner struct {
	Store  Storer
	Log    Logger
	Config Configurer
}

// ensure we continue to comply with the Runnerer interface
var runner Runnerer = &Runner{}

// Init - override to perform custom initialization
func (rnr *Runner) Init(c Configurer, l Logger, s Storer) error {

	rnr.Config = c
	rnr.Log = l
	rnr.Store = s

	rnr.Log.Printf("** Initialised **")

	return nil
}

// Run - override to perform custom running
func (rnr *Runner) Run(args map[string]interface{}) error {

	rnr.Log.Printf("** Run **")

	// default handler
	router, err := rnr.defaultRouter()
	if err != nil {
		rnr.Log.Printf("Failed default router >%v<", err)
		return err
	}

	return http.ListenAndServe(":8080", router)
}

// Router - override to implement custom routing
func (rnr *Runner) Router(router *httprouter.Router) error {

	rnr.Log.Printf("** Router **")

	return nil
}

// defaultRouter - implements default routes based on runner configuration options
func (rnr *Runner) defaultRouter() (*httprouter.Router, error) {

	rnr.Log.Printf("** DefaultRouter **")

	// default routes
	router := httprouter.New()
	router.GET("/", rnr.defaultMiddleware(rnr.defaultHandler))

	// service defined routes
	err := rnr.Router(router)
	if err != nil {
		rnr.Log.Printf("Failed router >%v<", err)
		return nil, err
	}

	return router, nil
}

// defaultMiddleware - implements middlewares based on runner configuration
func (rnr *Runner) defaultMiddleware(h httprouter.Handle) httprouter.Handle {

	rnr.Log.Printf("** DefaultMiddleware **")

	h, _ = rnr.BasicAuth(h)

	return h
}

// defaultHandler -
func (rnr *Runner) defaultHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rnr.Log.Printf("** DefaultHandler **")

	fmt.Fprint(w, "Default\n")
}

// BasicAuth -
func (rnr *Runner) BasicAuth(h httprouter.Handle) (httprouter.Handle, error) {

	// TODO: complete basic authorization implementation
	isAuthorized := func(user, password string) bool {
		return true
	}

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// get basic auth credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && isAuthorized(user, password) {
			// delegate request
			h(w, r, ps)
		} else {
			// unauthorized
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}

	return handle, nil
}
