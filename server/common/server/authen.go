package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/modeller"
)

// Authen -
func (rnr *Runner) Authen(h Handle) (Handle, error) {

	// TODO: implement authentication via configuration
	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

		l.Info("** Authen ** TODO: Authen unimplemented")

		h(w, r, p, l, m)
	}

	return handle, nil
}
