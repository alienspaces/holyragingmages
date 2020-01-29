package middleware

import (
	"bytes"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ContextData - data type for data context  key
type ContextData string

const (
	// ContextDataKey - context key for request body data
	ContextDataKey ContextData = "data"
)

// Data -
func Data(h httprouter.Handle) (httprouter.Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// read body into a string
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		data := buf.String()

		// Add data to context
		ctx := context.WithValue(r.Context(), ContextDataKey, data)

		// delegate request
		h(w, r.WithContext(ctx), ps)
	}

	return handle, nil
}
