package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	Model  Modeller

	// configuration for routes, handlers and middleware
	HandlerConfig []HandlerConfig

	// composable functions
	RouterFunc     func(router *httprouter.Router) error
	MiddlewareFunc func(h httprouter.Handle) (httprouter.Handle, error)
	HandlerFunc    func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

// MiddlewareConfig - configuration for global default middleware
type MiddlewareConfig struct {
	AuthenType               string
	AuthzType                string
	ValidateSchemaLocation   string
	ValidateSchemaMain       string
	ValidateSchemaReferences []string
}

// HandlerConfig - configuration for routes, handlers and middleware
type HandlerConfig struct {
	Method           string
	Path             string
	HandlerFunc      func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	MiddlewareConfig MiddlewareConfig
}

// ensure we continue to comply with the Runnerer interface
var runner Runnable = &Runner{}

// Init - override to perform custom initialization
func (rnr *Runner) Init(c Configurer, l Logger, s Storer) error {

	rnr.Config = c
	rnr.Log = l
	rnr.Store = s

	rnr.Log.Info("** Initialise **")

	// router
	if rnr.RouterFunc == nil {
		rnr.RouterFunc = rnr.Router
	}

	// middleware
	if rnr.MiddlewareFunc == nil {
		rnr.MiddlewareFunc = rnr.Middleware
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

// Router - default RouterFunc, override this function for custom routes
func (rnr *Runner) Router(router *httprouter.Router) error {

	rnr.Log.Info("** Router **")

	return nil
}

// Middleware - default MiddlewareFunc, override this function for custom middleware
func (rnr *Runner) Middleware(h httprouter.Handle) (httprouter.Handle, error) {

	rnr.Log.Info("** Middleware **")

	return h, nil
}

// Handler - default HandlerFunc, override this function for custom handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rnr.Log.Info("** Handler **")

	fmt.Fprint(w, "Ok!\n")
}

// DefaultRouter - implements default routes based on runner configuration options
func (rnr *Runner) DefaultRouter() (*httprouter.Router, error) {

	rnr.Log.Info("** DefaultRouter **")

	// default routes
	r := httprouter.New()

	// default index handler
	h, err := rnr.DefaultMiddleware("/", rnr.HandlerFunc)
	if err != nil {
		rnr.Log.Warn("Failed default middleware >%v<", err)
		return nil, err
	}
	r.GET("/", h)

	// register configured routes
	for _, hc := range rnr.HandlerConfig {
		h, err := rnr.DefaultMiddleware(hc.Path, hc.HandlerFunc)
		if err != nil {
			rnr.Log.Warn("Failed registering handler >%v<", err)
			return nil, err
		}
		switch hc.Method {
		case http.MethodGet:
			r.GET(hc.Path, h)
		case http.MethodPost:
			r.POST(hc.Path, h)
		case http.MethodPut:
			r.PUT(hc.Path, h)
		case http.MethodPatch:
			r.PATCH(hc.Path, h)
		case http.MethodDelete:
			r.DELETE(hc.Path, h)
		case http.MethodOptions:
			r.OPTIONS(hc.Path, h)
		case http.MethodHead:
			r.HEAD(hc.Path, h)
		default:
			rnr.Log.Warn("Router HTTP method >%s< not supported", hc.Method)
			return nil, fmt.Errorf("Router HTTP method >%s< not supported", hc.Method)
		}
	}

	// service defined routes
	err = rnr.RouterFunc(r)
	if err != nil {
		rnr.Log.Warn("Failed router >%v<", err)
		return nil, err
	}

	return r, nil
}

// DefaultMiddleware - implements middlewares based on runner configuration
func (rnr *Runner) DefaultMiddleware(path string, h httprouter.Handle) (httprouter.Handle, error) {

	rnr.Log.Info("** DefaultMiddleware **")

	// validate body data
	h, err := rnr.Validate(path, h)
	if err != nil {
		rnr.Log.Warn("Failed adding validate middleware >%v<", err)
		return nil, err
	}

	// request body data
	h, err = rnr.Data(h)
	if err != nil {
		rnr.Log.Warn("Failed adding data middleware >%v<", err)
		return nil, err
	}

	// authz
	h, err = rnr.Authz(h)
	if err != nil {
		rnr.Log.Warn("Failed adding authz middleware >%v<", err)
		return nil, err
	}

	// authen
	h, err = rnr.Authen(h)
	if err != nil {
		rnr.Log.Warn("Failed adding authen middleware >%v<", err)
		return nil, err
	}

	// service defined routes
	return rnr.MiddlewareFunc(h)
}
