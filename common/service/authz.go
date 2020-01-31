package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Authz -
func (rnr *Runner) Authz(h Handle) (Handle, error) {

	// TODO: implement authorization via configuration
	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m Modeller) {
		rnr.Log.Info("** Authz ** TODO: Authz unimplemented")
		h(w, r, p, m)
	}

	return handle, nil
}
