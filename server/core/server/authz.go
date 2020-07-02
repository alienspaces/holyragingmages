package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
)

// Authz -
func (rnr *Runner) Authz(h Handle) (Handle, error) {

	// TODO: implement authorization via configuration
	handle := func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

		l.Debug("** Authz ** TODO: Authz unimplemented")

		h(w, r, pp, qp, l, m)
	}

	return handle, nil
}
