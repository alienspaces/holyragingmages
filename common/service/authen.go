package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Authen -
func (rnr *Runner) Authen(h httprouter.Handle) (httprouter.Handle, error) {

	// TODO: implement authentication via configuration
	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rnr.Log.Info("** Authen ** TODO: Authen unimplemented")
		h(w, r, ps)
	}

	return handle, nil
}
