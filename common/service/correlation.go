package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
)

// Correlation -
func (rnr *Runner) Correlation(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
			rnr.Log.Info("Generated correlation ID >%s<", correlationID)
		}
		rnr.Log.Context("correlation-id", correlationID)

		h(w, r, p, m)
	}

	return handle, nil
}
