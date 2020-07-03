package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/payloader"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/storer"
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
	RunHTTPFunc    func(args map[string]interface{}) error
	RunDaemonFunc  func(args map[string]interface{}) error
	RouterFunc     func(router *httprouter.Router) error
	MiddlewareFunc func(h Handle) (Handle, error)
	HandlerFunc    func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller)
	PreparerFunc   func(l logger.Logger) (preparer.Preparer, error)
	ModellerFunc   func(l logger.Logger) (modeller.Modeller, error)
	PayloaderFunc  func(l logger.Logger) (payloader.Payloader, error)
}

var _ runnable.Runnable = &Runner{}

// Handle - custom service handle
type Handle func(w http.ResponseWriter, r *http.Request, pathParams httprouter.Params, queryParams map[string]interface{}, l logger.Logger, m modeller.Modeller)

// MiddlewareConfig - configuration for global default middleware
type MiddlewareConfig struct {
	AuthenType               string
	AuthzType                string
	ValidateSchemaLocation   string
	ValidateSchemaMain       string
	ValidateSchemaReferences []string
	// ValidateQueryParams - A whitelist of allowed query parameters
	ValidateQueryParams []string
}

// HandlerConfig - configuration for routes, handlers and middleware
type HandlerConfig struct {
	// Method - The HTTP method
	Method string
	// Path - The HTTP request URI including :parameter placeholders
	Path string
	// HandlerFunc - Function to handle requests for this method and path
	HandlerFunc func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller)
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
// type Request struct {
// 	Pagination RequestPagination `json:"pagination"`
// }

// // RequestPagination -
// type RequestPagination struct {
// 	PageNumber int `json:"page_number"`
// 	PageSize   int `json:"page_size"`
// }

// Response -
type Response struct {
	Error      *ResponseError      `json:"error,omitempty"`
	Pagination *ResponsePagination `json:"pagination,omitempty"`
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

	rnr.Log.Debug("** Initialise **")

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
	if rnr.RunHTTPFunc == nil {
		rnr.RunHTTPFunc = rnr.RunHTTP
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

// TODO: Use this function from HTTP middleware Tx

// InitTx initialises a new database transaction returning a prepared modeller
func (rnr *Runner) InitTx(l logger.Logger) (modeller.Modeller, error) {

	// NOTE: The modeller is created an initialised with every request instead of
	// creating and assigning to a runner struct "Model" property at start up.
	// This prevents directly accessing a shared property from with the handler
	// function which is running in a goroutine. Otherwise accessing the "Model"
	// property would require locking and block simultaneous requests.

	// modeller
	if rnr.ModellerFunc == nil {
		l.Warn("Runner ModellerFunc is nil")
		return nil, fmt.Errorf("ModellerFunc is nil")
	}

	m, err := rnr.ModellerFunc(l)
	if err != nil {
		l.Warn("Failed ModellerFunc >%v<", err)
		return nil, err
	}

	if m == nil {
		l.Warn("Modeller is nil, cannot continue")
		return nil, err
	}

	tx, err := rnr.Store.GetTx()
	if err != nil {
		l.Warn("Failed getting DB connection >%v<", err)
		return m, err
	}

	// NOTE: The PREPARER is created an initialised with every request instead of
	// creating and assigning to a runner struct "Prepare" property at start up.
	// This ensures statements are valid for the current database transaction.

	// preparer
	if rnr.PreparerFunc == nil {
		l.Warn("Runner PreparerFunc is nil")
		return m, err
	}

	p, err := rnr.PreparerFunc(l)
	if err != nil {
		l.Warn("Failed PreparerFunc >%v<", err)
		return m, err
	}

	if p == nil {
		l.Warn("Preparer is nil, cannot continue")
		return m, err
	}

	err = p.Init(tx)
	if err != nil {
		l.Warn("Failed init preparer >%v<", err)
		return m, err
	}

	err = m.Init(p, tx)
	if err != nil {
		l.Warn("Failed init modeller >%v<", err)
		return m, err
	}

	return m, nil
}

// Run starts the HTTP server and daemon processes. Override to implement a custom run function.
func (rnr *Runner) Run(args map[string]interface{}) (err error) {

	rnr.Log.Debug("** Run **")

	// signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// run HTTP server
	go func() {
		rnr.Log.Debug("** Running HTTP server process **")
		err = rnr.RunHTTPFunc(args)
		if err != nil {
			rnr.Log.Error("Failed run server >%v<", err)
			sigChan <- syscall.SIGTERM
		}
		rnr.Log.Debug("** HTTP server process ended **")
	}()

	// run daemon server
	go func() {
		rnr.Log.Debug("** Running daemon process **")
		err = rnr.RunDaemonFunc(args)
		if err != nil {
			rnr.Log.Error("Failed run daemon >%v<", err)
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