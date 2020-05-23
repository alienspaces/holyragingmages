package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
)

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

// RunHTTP - Starts the HTTP server process. Override to implement a custom HTTP server run function.
// The server process exposes a REST API and is intended for clients to manage resources and
// perform actions.
func (rnr *Runner) RunHTTP(args map[string]interface{}) error {

	rnr.Log.Debug("** RunHTTP **")

	// default handler
	router, err := rnr.DefaultRouter()
	if err != nil {
		rnr.Log.Warn("Failed default router >%v<", err)
		return err
	}

	port := rnr.Config.Get("APP_PORT")
	if port == "" {
		rnr.Log.Warn("Missing APP_PORT, cannot start server")
		return fmt.Errorf("Missing APP_PORT, cannot start server")
	}

	// cors
	c := cors.New(cors.Options{
		Debug:          true,
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{
			"X-ProgramID", "X-ProgramName", "Content-Type",
			"X-Authorization", "X-Authorization-Token",
			"Origin", "X-Requested-With", "Accept",
			"Access-Control-Allow-Origin",
			"X-CSRF-Token",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// serve
	rnr.Log.Info("Server running at: http://localhost:%s", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}

// Router - default RouterFunc, override this function for custom routes
func (rnr *Runner) Router(router *httprouter.Router) error {

	rnr.Log.Info("** Router **")

	return nil
}

// Middleware - default MiddlewareFunc, override this function for custom middleware
func (rnr *Runner) Middleware(h Handle) (Handle, error) {

	rnr.Log.Info("** Middleware **")

	return h, nil
}

// Payloader - default PayloaderFunc, override this function for custom payload handling
func (rnr *Runner) Payloader(l logger.Logger) (payloader.Payloader, error) {

	l.Info("** Payloader **")

	return nil, nil
}

// Handler - default HandlerFunc, override this function for custom handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

	l.Info("** Handler **")

	fmt.Fprint(w, "Ok!\n")
}

// DefaultRouter - implements default routes based on runner configuration options
func (rnr *Runner) DefaultRouter() (*httprouter.Router, error) {

	rnr.Log.Info("** DefaultRouter **")

	// default routes
	r := httprouter.New()

	// default index/healthz handler
	h, err := rnr.DefaultMiddleware("/", rnr.HandlerFunc)
	if err != nil {
		rnr.Log.Warn("Failed default middleware >%v<", err)
		return nil, err
	}
	r.GET("/", h)
	r.GET("/healthz", h)

	// register configured routes
	for _, hc := range rnr.HandlerConfig {

		rnr.Log.Info("** Router ** method >%s< path >%s<", hc.Method, hc.Path)

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

	// server defined routes
	err = rnr.RouterFunc(r)
	if err != nil {
		rnr.Log.Warn("Failed router >%v<", err)
		return nil, err
	}

	return r, nil
}

// DefaultMiddleware - implements middlewares based on runner configuration
func (rnr *Runner) DefaultMiddleware(path string, h Handle) (httprouter.Handle, error) {

	rnr.Log.Info("** DefaultMiddleware **")

	// tx
	h, err := rnr.Tx(h)
	if err != nil {
		rnr.Log.Warn("Failed adding tx middleware >%v<", err)
		return nil, err
	}

	// validate body data
	h, err = rnr.Validate(path, h)
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

	// correlation
	h, err = rnr.Correlation(h)
	if err != nil {
		rnr.Log.Warn("Failed adding correlation middleware >%v<", err)
		return nil, err
	}

	// server defined routes
	h, err = rnr.MiddlewareFunc(h)
	if err != nil {
		rnr.Log.Warn("Failed middleware >%v<", err)
		return nil, err
	}

	// wrap everything in a httprouter Handler
	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// delegate
		h(w, r, p, rnr.Log, nil)
	}

	return handle, nil
}

// ReadRequest -
func (rnr *Runner) ReadRequest(l logger.Logger, r *http.Request, s interface{}) error {

	p, err := rnr.PayloaderFunc(l)
	if err != nil {
		l.Warn("Failed PayloaderFunc >%v<", err)
		return err
	}
	if p == nil {
		return fmt.Errorf("Payloader is nil, cannot read request")
	}

	return p.ReadRequest(r, s)
}

// WriteResponse -
func (rnr *Runner) WriteResponse(l logger.Logger, w http.ResponseWriter, s interface{}) error {

	p, err := rnr.PayloaderFunc(l)
	if err != nil {
		l.Warn("Failed PayloaderFunc >%v<", err)
		return err
	}
	if p == nil {
		return fmt.Errorf("Payloader is nil, cannot write response")
	}

	// determine response status
	status := http.StatusOK

	switch r := s.(type) {
	case *Response:
		l.Info("Payload type is a base server response >%#v<", r)
		if r.Error.Code != "" {
			switch r.Error.Code {
			case ErrorCodeNotFound:
				status = http.StatusNotFound
			case ErrorCodeValidation:
				status = http.StatusBadRequest
			case ErrorCodeSystem:
				status = http.StatusInternalServerError
			}
		}
	case Response:
		l.Info("Payload type is a base server response >%#v<", r)
		if r.Error.Code != "" {
			switch r.Error.Code {
			case ErrorCodeNotFound:
				status = http.StatusNotFound
			case ErrorCodeValidation:
				status = http.StatusBadRequest
			case ErrorCodeSystem:
				status = http.StatusInternalServerError
			}
		}
	default:
		l.Info("Payload type is not a base server response >%#v<", r)
		//
	}

	l.Info("Write response status >%d<", status)

	return p.WriteResponse(w, status, s)
}
