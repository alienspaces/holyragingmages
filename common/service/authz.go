package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Authz -
func (rnr *Runner) Authz(h httprouter.Handle) (httprouter.Handle, error) {

	// TODO: implement authorization via configuration
	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rnr.Log.Info("TODO: Authz unimplemented")
		h(w, r, ps)
	}

	return handle, nil
}
