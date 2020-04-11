package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
)

// RunServer - Starts the HTTP server process. Override to implement a custom HTTP server run function.
// The server process exposes a REST API and is intended for clients to manage resources and
// perform actions.
func (rnr *Runner) RunServer(args map[string]interface{}) error {

	rnr.Log.Debug("** RunServer **")

	// default handler
	router, err := rnr.DefaultRouter()
	if err != nil {
		rnr.Log.Warn("Failed default router >%v<", err)
		return err
	}

	port := rnr.Config.Get("APP_PORT")
	if port == "" {
		rnr.Log.Warn("Missing APP_PORT, cannot start service")
		return fmt.Errorf("Missing APP_PORT, cannot start service")
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", port), router)
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
func (rnr *Runner) Payloader() (payloader.Payloader, error) {

	rnr.Log.Info("** Payloader **")

	return nil, nil
}

// Handler - default HandlerFunc, override this function for custom handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

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

	// service defined routes
	h, err = rnr.MiddlewareFunc(h)
	if err != nil {
		rnr.Log.Warn("Failed middleware >%v<", err)
		return nil, err
	}

	// wrap everything in a httprouter Handler
	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// delegate
		h(w, r, p, nil)
	}

	return handle, nil
}

// ReadRequest -
func (rnr *Runner) ReadRequest(r *http.Request, s interface{}) error {
	p, err := rnr.PayloaderFunc()
	if err != nil {
		rnr.Log.Warn("Failed PayloaderFunc >%v<", err)
		return err
	}
	if p == nil {
		return fmt.Errorf("Payloader is nil, cannot read request")
	}
	return p.ReadRequest(r, s)
}

// WriteResponse -
func (rnr *Runner) WriteResponse(w http.ResponseWriter, s interface{}) error {
	p, err := rnr.PayloaderFunc()
	if err != nil {
		rnr.Log.Warn("Failed PayloaderFunc >%v<", err)
		return err
	}
	if p == nil {
		return fmt.Errorf("Payloader is nil, cannot write response")
	}
	return p.WriteResponse(w, s)
}
