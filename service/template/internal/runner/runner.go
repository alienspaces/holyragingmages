package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/service"
)

// Runner -
type Runner struct {
	service.Runner
}

// NewRunner -
func NewRunner() *Runner {

	r := Runner{}

	r.RouterFunc = r.Router
	r.MiddlewareFunc = r.Middleware
	r.HandlerFunc = r.Handler

	return &r
}

// Router -
func (rnr *Runner) Router(r *httprouter.Router) error {

	rnr.Log.Info("** Template Router **")

	r.GET("/templates", rnr.DefaultMiddleware(rnr.Handler))

	return nil
}

// Middleware -
func (rnr *Runner) Middleware(h httprouter.Handle) httprouter.Handle {

	rnr.Log.Info("** Template Middleware **")

	return h
}

// Handler -
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rnr.Log.Info("** Template Handler **")

	fmt.Fprint(w, "Okie Dokie!\n")
}
