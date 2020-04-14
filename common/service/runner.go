package service

import (
	"fmt"
	"net/http"

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
	Config  configurer.Configurer
	Log     logger.Logger
	Store   storer.Storer
	Prepare preparer.Preparer

	// configuration for routes, handlers and middleware
	HandlerConfig []HandlerConfig

	// composable functions
	RouterFunc     func(router *httprouter.Router) error
	MiddlewareFunc func(h Handle) (Handle, error)
	HandlerFunc    func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller)
	ModellerFunc   func() (modeller.Modeller, error)
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

// Request -
type Request struct {
	Pagination RequestPagination `json:"pagination"`
}

// RequestPagination -
type RequestPagination struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

// Response -
type Response struct {
	Error      ResponseError      `json:"error"`
	Pagination ResponsePagination `json:"pagination"`
}

// ResponseError -
type ResponseError struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

// ResponsePagination -
type ResponsePagination struct {
	Number int `json:"page_number"`
	Size   int `json:"page_size"`
	Count  int `json:"page_count"`
}

// ensure we comply with the Runnerer interface
var _ runnable.Runnable = &Runner{}

// Init - override to perform custom initialization
func (rnr *Runner) Init(c configurer.Configurer, l logger.Logger, s storer.Storer, p preparer.Preparer) error {

	rnr.Config = c
	if rnr.Config == nil {
		msg := "Configurer undefined, cannot init runner"
		rnr.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	rnr.Log = l
	if rnr.Log == nil {
		msg := "Logger undefined, cannot init runner"
		rnr.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	rnr.Store = s
	if rnr.Store == nil {
		msg := "Storer undefined, cannot init runner"
		rnr.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	rnr.Prepare = p
	if rnr.Prepare == nil {
		msg := "Preparer undefined, cannot init runner"
		rnr.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	rnr.Log.Info("** Initialise **")

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
func (rnr *Runner) Preparer() (preparer.Preparer, error) {

	rnr.Log.Info("** Preparer **")

	return nil, nil
}

// Modeller - default ModellerFunc, override this function for custom model
func (rnr *Runner) Modeller() (modeller.Modeller, error) {

	rnr.Log.Info("** Modeller **")

	return nil, nil
}
