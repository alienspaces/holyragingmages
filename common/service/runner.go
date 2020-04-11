package service

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
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
	Store  storer.Storer
	Log    logger.Logger
	Config configurer.Configurer

	// configuration for routes, handlers and middleware
	HandlerConfig []HandlerConfig

	// composable functions
	RouterFunc     func(router *httprouter.Router) error
	MiddlewareFunc func(h Handle) (Handle, error)
	HandlerFunc    func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller)
	PreparerFunc   func(l logger.Logger, tx *sqlx.Tx) (preparer.Preparer, error)
	ModellerFunc   func(c configurer.Configurer, l logger.Logger, s storer.Storer) (modeller.Modeller, error)
	PayloaderFunc  func() (payloader.Payloader, error)
}

var _ runnable.Runnable = &Runner{}

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
	HandlerFunc      func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller)
	MiddlewareConfig MiddlewareConfig
}

// ensure we comply with the Runnerer interface
var _ runnable.Runnable = &Runner{}

// Init - override to perform custom initialization
func (rnr *Runner) Init(c configurer.Configurer, l logger.Logger, s storer.Storer) error {

	rnr.Config = c
	rnr.Log = l
	rnr.Store = s

	rnr.Log.Info("** Initialise **")

	// preparer
	if rnr.PreparerFunc == nil {
		rnr.PreparerFunc = rnr.Preparer
	}

	// model
	if rnr.ModellerFunc == nil {
		rnr.ModellerFunc = rnr.Modeller
	}

	// http server - router
	if rnr.RouterFunc == nil {
		rnr.RouterFunc = rnr.Router
	}

	// http server - middleware
	if rnr.MiddlewareFunc == nil {
		rnr.MiddlewareFunc = rnr.Middleware
	}

	// http server - handler
	if rnr.HandlerFunc == nil {
		rnr.HandlerFunc = rnr.Handler
	}

	// http server - payloader
	if rnr.PayloaderFunc == nil {
		rnr.PayloaderFunc = rnr.Payloader
	}

	return nil
}

// Run - Runs the HTTP server and daemon processes. Override to implement a custom run function.
func (rnr *Runner) Run(args map[string]interface{}) (err error) {

	rnr.Log.Debug("** Run **")

	// run HTTP server
	go func() {
		rnr.Log.Debug("** Running HTTP server process **")
		err = rnr.RunServer(args)
		if err != nil {
			rnr.Log.Warn("Failed run server >%v<", err)
		}
	}()

	// run daemon server
	go func() {
		rnr.Log.Debug("** Running daemon process **")
		err = rnr.RunDaemon(args)
		if err != nil {
			rnr.Log.Warn("Failed run daemon >%v<", err)
		}
	}()

	return err
}

// Preparer - default preparer.PreparerFunc, override this function for custom model
func (rnr *Runner) Preparer(l logger.Logger, tx *sqlx.Tx) (preparer.Preparer, error) {

	rnr.Log.Info("** Preparer **")

	return nil, nil
}

// Modeller - default ModellerFunc, override this function for custom model
func (rnr *Runner) Modeller(c configurer.Configurer, l logger.Logger, s storer.Storer) (modeller.Modeller, error) {

	rnr.Log.Info("** Modeller **")

	return nil, nil
}
