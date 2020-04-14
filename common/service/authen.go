package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
)

// Authen -
func (rnr *Runner) Authen(h Handle) (Handle, error) {

	// TODO: implement authentication via configuration
	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

		rnr.Log.Info("** Authen ** TODO: Authen unimplemented")

		// ALIEN:
		rnr.Log.Info("** Authen ** params >%#v<", p)

		h(w, r, p, m)
	}

	return handle, nil
}
