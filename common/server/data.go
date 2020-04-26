package server

import (
	"bytes"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
)

// ContextData - data type for data context  key
type ContextData string

const (
	// ContextKeyData - context key for request body data
	ContextKeyData ContextData = "data"
)

// Data -
func (rnr *Runner) Data(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

		// read body into a string
		buf := new(bytes.Buffer)
		if r.Body != nil {
			bytes, err := buf.ReadFrom(r.Body)
			if err != nil {
				rnr.Log.Warn("Failed reading data buffer >%v<", err)
				http.Error(w, "Server Error", http.StatusInternalServerError)
				return
			}
			rnr.Log.Info("Read >%d< bytes", bytes)
		}
		data := buf.String()

		// Add data to context
		ctx := context.WithValue(r.Context(), ContextKeyData, data)

		// delegate request
		h(w, r.WithContext(ctx), p, m)
	}

	return handle, nil
}
