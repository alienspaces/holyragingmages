package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
)

// Authz -
func (rnr *Runner) Authz(path string, h HandlerFunc) (HandlerFunc, error) {

	// TODO: implement authorization via configuration
	handle := func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

		l.Debug("** Authz ** TODO: Authz unimplemented")

		// TODO: Source config from handler, parse JWT, set auth context and pass it
		// down through the handler stack somehow. Another argument? Context on the
		// request object? As long as it is common functions for access / checking.

		h(w, r, pp, qp, l, m)
	}

	return handle, nil
}
