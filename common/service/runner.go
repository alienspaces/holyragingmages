package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/service/middleware"
)

const (
	// ConfigKeyValidateSchemaLocation - Directory location of JSON schema's
	ConfigKeyValidateSchemaLocation string = "validateSchemaLocation"
	// ConfigKeyValidateMainSchema - Main schema that can include reference schema's
	ConfigKeyValidateMainSchema string = "validateMainSchema"
	// ConfigKeyValidateReferenceSchemas - Schema referenced from the main schema
	ConfigKeyValidateReferenceSchemas string = "validateReferenceSchemas"
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

	rnr.Log.Info("** Initialised **")

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

	rnr.Log.Debug("** Run **")

	// default handler
	router, err := rnr.DefaultRouter()
	if err != nil {
		rnr.Log.Warn("Failed default router >%v<", err)
		return err
	}

	return http.ListenAndServe(":8080", router)
}

// Router - set RouterFunc to set up custom routes
func (rnr *Runner) Router(router *httprouter.Router) error {

	rnr.Log.Info("** Router **")

	return nil
}

// Middleware - set MiddlewareFunc to set up custom middleware
func (rnr *Runner) Middleware(h httprouter.Handle) httprouter.Handle {

	rnr.Log.Info("** Middleware **")

	return h
}

// Handler - set HandlerFunc to set up custom handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rnr.Log.Info("** Handler **")

	fmt.Fprint(w, "Ok!\n")
}

// DefaultRouter - implements default routes based on runner configuration options
func (rnr *Runner) DefaultRouter() (*httprouter.Router, error) {

	rnr.Log.Info("** DefaultRouter **")

	// default routes
	r := httprouter.New()
	r.GET("/", rnr.DefaultMiddleware(rnr.HandlerFunc))

	// service defined routes
	err := rnr.RouterFunc(r)
	if err != nil {
		rnr.Log.Warn("Failed router >%v<", err)
		return nil, err
	}

	return r, nil
}

// DefaultMiddleware - implements middlewares based on runner configuration
func (rnr *Runner) DefaultMiddleware(h httprouter.Handle) httprouter.Handle {

	rnr.Log.Info("** DefaultMiddleware **")

	// h, _ = middleware.BasicAuth(h)

	// request body data
	h, _ = middleware.Data(h)

	// TODO: decide whether or not to bake configuration into
	// type Runner or use config as below..

	// JSON schema validatation
	schemaLoc := rnr.Config.Get(ConfigKeyValidateSchemaLocation)
	mainSchema := rnr.Config.Get(ConfigKeyValidateMainSchema)
	referenceSchemas := rnr.Config.Get(ConfigKeyValidateReferenceSchemas)

	if schemaLoc != "" && mainSchema != "" {
		h, _ = middleware.SchemaValidate(
			h,
			schemaLoc,
			mainSchema,
			referenceSchemas,
		)
	}

	// service defined routes
	return rnr.MiddlewareFunc(h)
}
