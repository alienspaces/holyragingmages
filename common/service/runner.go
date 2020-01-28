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

	// composable functions
	RouterFunc     func(router *httprouter.Router) error
	MiddlewareFunc func(h httprouter.Handle) httprouter.Handle
	HandlerFunc    func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

// ensure we continue to comply with the Runnerer interface
var runner Runnable = &Runner{}

// Init - override to perform custom initialization
func (rnr *Runner) Init(c Configurer, l Logger, s Storer) error {

	rnr.Config = c
	rnr.Log = l
	rnr.Store = s

	rnr.Log.Printf("** Initialised **")

	// router
	if rnr.RouterFunc == nil {
		rnr.RouterFunc = rnr.Router
	}

	// handler
	if rnr.HandlerFunc == nil {
		rnr.HandlerFunc = rnr.Handler
	}

	return nil
}

// Run - override to perform custom running
func (rnr *Runner) Run(args map[string]interface{}) error {

	rnr.Log.Printf("** Run **")

	// default handler
	router, err := rnr.DefaultRouter()
	if err != nil {
		rnr.Log.Printf("Failed default router >%v<", err)
		return err
	}

	return http.ListenAndServe(":8080", router)
}

// Router - set RouterFunc to set up custom routes
func (rnr *Runner) Router(router *httprouter.Router) error {

	rnr.Log.Printf("** Router **")

	return nil
}

// Middleware - set MiddlewareFunc to set up custom middleware
func (rnr *Runner) Middleware(h httprouter.Handle) httprouter.Handle {

	rnr.Log.Printf("** Middleware **")

	return h
}

// Handler - set HandlerFunc to set up custom handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rnr.Log.Printf("** Handler **")

	fmt.Fprint(w, "Ok!\n")
}

// DefaultRouter - implements default routes based on runner configuration options
func (rnr *Runner) DefaultRouter() (*httprouter.Router, error) {

	rnr.Log.Printf("** DefaultRouter **")

	// default routes
	r := httprouter.New()
	r.GET("/", rnr.DefaultMiddleware(rnr.HandlerFunc))

	// service defined routes
	err := rnr.RouterFunc(r)
	if err != nil {
		rnr.Log.Printf("Failed router >%v<", err)
		return nil, err
	}

	return r, nil
}

// DefaultMiddleware - implements middlewares based on runner configuration
func (rnr *Runner) DefaultMiddleware(h httprouter.Handle) httprouter.Handle {

	rnr.Log.Printf("** DefaultMiddleware **")

	h, _ = rnr.BasicAuth(h)

	// service defined routes
	return rnr.MiddlewareFunc(h)
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
