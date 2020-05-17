package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
)

// Authz -
func (rnr *Runner) Authz(h Handle) (Handle, error) {

	// TODO: implement authorization via configuration
	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

		l.Info("** Authz ** TODO: Authz unimplemented")

		h(w, r, p, l, m)
	}

	return handle, nil
}
