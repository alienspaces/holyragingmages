package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
)

// Authz -
func (rnr *Runner) Authz(h Handle) (Handle, error) {

	// TODO: implement authorization via configuration
	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

		rnr.Log.Info("** Authz ** TODO: Authz unimplemented")

		// ALIEN:
		rnr.Log.Info("** Authz ** params >%#v<", p)

		h(w, r, p, m)
	}

	return handle, nil
}
