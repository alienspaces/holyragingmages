package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// BasicAuth -
func BasicAuth(h httprouter.Handle) (httprouter.Handle, error) {

	// TODO: complete basic authorization implementation
	isAuthorized := func(user, password string) bool {
		return true
	}

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// get basic auth credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && isAuthorized(user, password) {
			// delegate request
			h(w, r, ps)
		} else {
			// unauthorized
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}

	return handle, nil
}
