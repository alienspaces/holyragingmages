package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
)

// Authen -
func (rnr *Runner) Authen(h Handle) (Handle, error) {

	// TODO: implement authentication via configuration
	handle := func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

		l.Debug("** Authen ** TODO: Authen unimplemented")

		h(w, r, pp, qp, l, m)
	}

	return handle, nil
}
