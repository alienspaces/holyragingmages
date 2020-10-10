package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
)

// Handler - default handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Item handler **")

	fmt.Fprint(w, "Hello from item!\n")
}

// Router -
func (rnr *Runner) Router(r *httprouter.Router) error {

	rnr.Log.Info("** Item Router **")

	return nil
}

// Middleware -
func (rnr *Runner) Middleware(h server.HandlerFunc) (server.HandlerFunc, error) {

	rnr.Log.Info("** Item Middleware **")

	return h, nil
}
