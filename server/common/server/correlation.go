package server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/modeller"
)

// Correlation -
func (rnr *Runner) Correlation(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, _ modeller.Modeller) {

		lc, err := l.NewInstance()
		if err != nil {
			rnr.Log.Warn("Failed new log instance >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
			lc.Info("Generated correlation ID >%s<", correlationID)
		}
		lc.Context("correlation-id", correlationID)

		h(w, r, p, lc, nil)
	}

	return handle, nil
}
