package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

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
	RunServerFunc  func(args map[string]interface{}) error
	RunDaemonFunc  func(args map[string]interface{}) error
	RouterFunc     func(router *httprouter.Router) error
	MiddlewareFunc func(h Handle) (Handle, error)
	HandlerFunc    func(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller)
	PreparerFunc   func(l logger.Logger) (preparer.Preparer, error)
	ModellerFunc   func(l logger.Logger) (modeller.Modeller, error)
	PayloaderFunc  func(l logger.Logger) (payloader.Payloader, error)
}

var _ runnable.Runnable = &Runner{}

// Handle - custom service handle
type Handle func(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller)

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
	// Method - The HTTP method
	Method string
	// Path - The HTTP request URI including :parameter placeholders
	Path string
	// QueryParams - A whitelist of allowed query parameters
	QueryParams []string
	// HandlerFunc - Function to handle requests for this method and path
	HandlerFunc func(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller)
	// MiddlewareConfig -
	MiddlewareConfig MiddlewareConfig
	// DocumentationConfig -
	DocumentationConfig DocumentationConfig
}

// DocumentationConfig - Configuration describing how to document a route
type DocumentationConfig struct {
	Document    bool
	Description string
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
func (rnr *Runner) Init(c configurer.Configurer, l logger.Logger, s storer.Storer) error {

	rnr.Log = l
	if rnr.Log == nil {
		msg := "Logger undefined, cannot init runner"
		return fmt.Errorf(msg)
	}

	rnr.Log.Info("** Initialise **")

	rnr.Config = c
	if rnr.Config == nil {
		msg := "Configurer undefined, cannot init runner"
		rnr.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	rnr.Store = s
	if rnr.Store == nil {
		msg := "Storer undefined, cannot init runner"
		rnr.Log.Warn(msg)
		return fmt.Errorf(msg)
	}

	// run server
	if rnr.RunServerFunc == nil {
		rnr.RunServerFunc = rnr.RunServer
	}

	// run daemon
	if rnr.RunDaemonFunc == nil {
		rnr.RunDaemonFunc = rnr.RunDaemon
	}

	// prepare
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

	// signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// run HTTP server
	go func() {
		rnr.Log.Debug("** Running HTTP server process **")
		err = rnr.RunServerFunc(args)
		if err != nil {
			rnr.Log.Warn("Failed run server >%v<", err)
			sigChan <- syscall.SIGTERM
		}
		rnr.Log.Debug("** HTTP server process ended **")
	}()

	// run daemon server
	go func() {
		rnr.Log.Debug("** Running daemon process **")
		err = rnr.RunDaemonFunc(args)
		if err != nil {
			rnr.Log.Warn("Failed run daemon >%v<", err)
			sigChan <- syscall.SIGTERM
		}
		rnr.Log.Debug("** Daemon process ended **")
	}()

	// wait
	sig := <-sigChan

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	bToMb := func(b uint64) uint64 {
		return b / 1024 / 1024
	}

	err = fmt.Errorf("Received SIG >%v< Mem Alloc >%d MiB< TotalAlloc >%d MiB< Sys >%d MiB< NumGC >%d<",
		sig,
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys),
		m.NumGC,
	)

	rnr.Log.Warn(">%v<", err)

	return err
}

// Preparer - default PreparerFunc, override this function for custom prepare
func (rnr *Runner) Preparer(l logger.Logger) (preparer.Preparer, error) {

	l.Info("** Preparer **")

	return nil, nil
}

// Modeller - default ModellerFunc, override this function for custom model
func (rnr *Runner) Modeller(l logger.Logger) (modeller.Modeller, error) {

	l.Info("** Modeller **")

	return nil, nil
}
