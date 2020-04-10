package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
)

// Correlation -
func (rnr *Runner) Correlation(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, m modeller.Modeller) {

		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
			rnr.Log.Info("Generated correlation ID >%s<", correlationID)
		}
		rnr.Log.Context("correlation-id", correlationID)

		h(w, r, ps, m)
	}

	return handle, nil
}
