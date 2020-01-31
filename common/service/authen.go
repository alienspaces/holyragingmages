package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Authen -
func (rnr *Runner) Authen(h Handle) (Handle, error) {

	// TODO: implement authentication via configuration
	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, m Modeller) {
		rnr.Log.Info("** Authen ** TODO: Authen unimplemented")
		h(w, r, ps, m)
	}

	return handle, nil
}
