package service

import (
	"bytes"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ContextData - data type for data context  key
type ContextData string

const (
	// ContextKeyData - context key for request body data
	ContextKeyData ContextData = "data"
)

// Data -
func (rnr *Runner) Data(h httprouter.Handle) (httprouter.Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		rnr.Log.Warn("** Data ** request URI >%s< method >%s<", r.RequestURI, r.Method)

		// read body into a string
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		data := buf.String()

		// Add data to context
		ctx := context.WithValue(r.Context(), ContextKeyData, data)

		// delegate request
		h(w, r.WithContext(ctx), ps)
	}

	return handle, nil
}
